package maps

import (
	"reflect"
	"sort"
)

func Range(in, out interface{}, fn func(k, v interface{}) reflect.Value) {
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

	if outv.Type().Kind() != reflect.Slice {
		panic("in param must be a slice interface{}")
	}

	var (
		iter = inv.MapRange()
	)

	for iter.Next() {
		reflect.Append(outv, fn(iter.Key(), iter.Value()))
	}
}

func Keys(in, out interface{}) {
	Range(in, out, func(k, v interface{}) reflect.Value {
		return reflect.ValueOf(k)
	})
}

func Values(in, out interface{}) {
	Range(in, out, func(k, v interface{}) reflect.Value {
		return reflect.ValueOf(v)
	})
}

// KeysString is get map keys.
// when input map[string]interface{} out []string
func KeysString(in map[string]interface{}) (out []string) {
	for k, _ := range in {
		out = append(out, k)
	}
	return
}

func KeysOrder(in interface{}) (out []string) {
	Range(in, out, func(k, v interface{}) reflect.Value {
		return reflect.ValueOf(k)
	})
	sort.Strings(out)
	return
}
