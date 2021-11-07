package maps

import (
	"reflect"
	"sort"

	"github.com/alonelucky/gtool/reflects"
)

// Range 遍历map, 返回数组
func Range(in, out interface{}, fn func(k, v interface{}) interface{}) {
	if in == nil {
		return
	}
	var (
		inv  = reflect.ValueOf(in)
		outv = reflect.ValueOf(out)
	)

	if inv.Type().Kind() != reflect.Map {
		panic("in param must be a map interface{}")
	}

	if v := reflects.Indirect(outv); outv.Kind() != reflect.Ptr || v.Kind() != reflect.Slice {
		panic("in param must be a slice ptr")
	}

	var (
		iter = inv.MapRange()
	)

	outv = reflects.Indirect(outv)
	var arr = reflect.MakeSlice(outv.Type(), 0, 4)
	for iter.Next() {
		v := fn(iter.Key().Interface(), iter.Value().Interface())
		vv := reflect.ValueOf(v)
		arr = reflect.Append(arr, vv)
	}
	outv.Set(arr)
}

func Keys(in, out interface{}) {
	Range(in, out, func(k, v interface{}) interface{} {
		return k
	})
}

func Values(in, out interface{}) {
	Range(in, out, func(k, v interface{}) interface{} {
		return v
	})
}

// KeysString is get map keys.
// when input map[string]interface{} out []string
func KeysString(in map[string]interface{}) (out []string) {
	for k, _ := range in {
		out = append(out, k)
	}
	sort.Strings(out)
	return
}

func KeysOrder(in interface{}) (out []string) {
	Range(in, out, func(k, v interface{}) interface{} {
		return k
	})
	return
}
