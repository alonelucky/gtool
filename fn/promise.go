package fn

import (
	"context"
	"fmt"
	"reflect"

	"github.com/alonelucky/gtool/numbers"
	"github.com/alonelucky/gtool/reflects"
)

type Promise func() (i interface{}, e error)

type result struct {
	data interface{}
	err  error
	idx  int
}

// 并发执行,并收集全部结果
func All(fns []Promise, maxs ...int) (lst []interface{}, err error) {
	fmt.Println(fns, maxs)
	var ctx = context.Background()
	var max = len(fns)
	if len(maxs) > 0 {
		max = maxs[0]
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var (
		cfn = make(chan Promise, max)
		ch  = make(chan *result, max)
	)

	for i := 0; i < max; i++ {
		go func(k int) {
			for {
				select {
				case <-ctx.Done():
					// 任务池终止
					return
				case fn := <-cfn:
					if fn == nil {
						continue
					}
					i, e := fn()
					ch <- &result{
						data: i,
						err:  e,
						idx:  k,
					}
				}
			}
		}(i)
	}

	l := len(fns)
	for i := 0; i < l; i++ {
		cfn <- fns[i]
	}

	close(cfn)

	for r := range ch {
		lst = append(lst, r.data)
		if err == nil && r.err != nil {
			cancel()
			close(ch)
			err = r.err
		}

		if len(lst) == l {
			cancel()
			close(ch)
		}
	}
	return
}

func WorkFlow(lst interface{}, max int, fn func(i interface{}, idx int) (v interface{}, e error)) (ret []interface{}, e error) {
	var (
		inv = reflects.Indirect(reflect.ValueOf(lst))

		ctx = context.Background()
		err []error
	)

	l := inv.Len()

	// 如果没有限制, 则自行限制10个
	if max <= 0 {
		max = numbers.MinInt(10, l)
	}

	var (
		work    = make(chan *result, max)
		results = make(chan *result, max)
	)

	ctx, cancel := context.WithCancel(ctx)
	defer func() {
		cancel()
		close(work)
	}()

	// 执行池
	for i := 0; i < max; i++ {
		go func() {
			for {
				select {
				case data := <-work:
					r, e := fn(data.data, data.idx)
					results <- &result{data: r, err: e, idx: data.idx}
				case <-ctx.Done():
					// 协程终止
					return
				}

			}
		}()
	}

	// 任务派发
	for i := 0; i < l; i++ {
		work <- &result{
			data: inv.Index(i).Interface(),
			idx:  i,
		}
	}

	for v := range results {
		l--
		// 接收数据
		if v.err != nil {
			err = append(err, v.err)
			cancel()
			close(results)
			continue
		}
		ret = append(ret, v.data)
		if l == 1 {
			cancel()
			close(results)
		}
	}

	if len(err) > 0 {
		e = err[0]
		return
	}

	return
}
