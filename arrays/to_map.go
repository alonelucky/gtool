package arrays

import (
	"reflect"

	"github.com/alonelucky/gtool/reflects"
)

// ToMap 将切片转换为map类型
//
// Convert slice to map type
func ToMap(in, out interface{}, fn func(item interface{}, i int) (k, v interface{})) {
	if in == nil {
		return
	}

	var (
		inv  = reflect.ValueOf(in)
		outv = reflect.ValueOf(out)
	)

	if !reflects.IsArray(inv) {
		return
	}

	outv = reflects.Indirect(outv)

	if outv.Kind() != reflect.Map {
		return
	}

	l := inv.Len()
	for i := 0; i < l; i++ {
		k, v := fn(inv.Index(i).Interface(), i)
		outv.SetMapIndex(reflect.ValueOf(k), reflect.ValueOf(v))
	}
}

// Column array<M> to map[T]M
//
// 数组类型转换为 map[T]M,
//
// 其中field为需要提取的key，M为map则为keyName,M为struct则为fieldName
//
// The array type is extracted as map type
func Column(lst, out interface{}, field ...string) {
	if lst == nil || out == nil || len(field) == 0 {
		return
	}

	var (
		key    = field[0]
		inv    = reflects.Indirect(reflect.ValueOf(lst))
		outv   = reflect.ValueOf(out)
		length int
	)

	if !reflects.IsArray(inv) || outv.Kind() != reflect.Ptr {
		return
	}

	if length = inv.Len(); length == 0 {
		return
	}

	outv = reflects.Indirect(outv)

	if outv.Kind() != reflect.Map {
		return
	}

	idx := -1

	if key == "" {
		idx = 0
	}

	for i := 0; i < length; i++ {
		v := inv.Index(i)
		switch cv := reflects.Indirect(v); cv.Kind() {
		case reflect.Struct:
			var sv reflect.Value
			if idx == -1 {
				sv = cv.FieldByName(key)
			} else {
				sv = cv.Field(idx)
			}
			outv.SetMapIndex(sv, v)
		case reflect.Map:
			outv.SetMapIndex(cv.MapIndex(reflect.ValueOf(key)), v)
		default:
			return
		}
	}
}
