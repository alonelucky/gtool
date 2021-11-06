package fn

import (
	"context"
	"reflect"
)

type Promise func() (i interface{}, e error)

type result struct {
	data interface{}
	err  error
	idx  int
}

// 并发执行,并收集全部结果
func All(fns []Promise, maxs ...int) (lst []interface{}, err error) {
	var ctx = context.Background()
	var max = len(fns)
	if len(maxs) > 0 {
		max = maxs[0]
	}

	ctx, cancel := context.WithCancel(ctx)
	defer func() {
		cancel()
	}()

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

	for i := 0; i < len(fns); i++ {
		cfn <- fns[i]
	}

	close(cfn)

	for r := range ch {
		lst = append(lst, r.data)
		if err == nil && r.err != nil {
			cancel()
			err = r.err
		}
	}

	return
}

func WorkFlow(lst interface{}, max int, fn func(i interface{}, idx int) (v interface{}, e error)) (ret []interface{}, e error) {
	var (
		inv = indirect(reflect.ValueOf(lst))

		ctx = context.Background()
		err []error
	)

	l := inv.Len()

	// 如果没有限制, 则自行限制10个
	if max <= 0 {
		max = l

		if max > 10 {
			max = 10
		}
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
		go func(i int) {
			work <- &result{
				data: inv.Index(i).Interface(),
				idx:  i,
			}
		}(i)
	}

	for v := range results {
		if l == 0 {
			close(results)
		}
		// 接收数据
		if v.err != nil {
			err = append(err, v.err)
			cancel()
			continue
		}
		ret = append(ret, v.data)
		//
		l--
	}

	if len(err) > 0 {
		e = err[0]
		return
	}

	return
}

func indirect(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Ptr {
		v = reflect.Indirect(v)
	}
	return v
}
