package arrays

import (
	"reflect"

	"github.com/alonelucky/gtool/reflects"
)

// Sort is order by slice.
//
// when out is nil, sort by input, input must be slice ptr.
//
// when out is no-nil, sort by input , output out, out must be slice ptr.
func Sort(in, out interface{}, fns ...func(a, b interface{}) int8) {
	if in == nil {
		return
	}

	var (
		inv  = reflect.ValueOf(in)
		outv = reflect.ValueOf(out)
	)

	inv = reflects.Indirect(inv)
	l := inv.Len()

	outv = reflects.Indirect(outv)
	if outv.Kind() == reflect.Slice {
		copyslice(outv, inv)
	} else {
		outv = inv
	}

	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l-1; j++ {
			jv := outv.Index(j)
			iv := outv.Index(i)
			var gte bool
			if len(fns) > 0 {
				gte = fns[0](iv.Interface(), jv.Interface()) > 0
			} else if reflects.SameKind(iv, jv) && reflects.Comparable(iv) {
				if reflects.IsInt(iv) {
					gte = iv.Int() > jv.Int()
				} else if reflects.IsUnit(iv) {
					gte = iv.Uint() > jv.Uint()
				} else if reflects.IsFloat(iv) {
					gte = iv.Float() > jv.Float()
				} else if iv.Kind() == reflect.String {
					gte = iv.String() > jv.String()
				} else {
					gte = iv == jv
				}
			}

			if gte {
				tmp := iv.Interface()
				iv.Set(jv)
				jv.Set(reflect.ValueOf(tmp))
			}
		}
	}
}

func SortInt32(in []int32) {
	l := len(in)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if in[i] > in[j] {
				in[i], in[j] = in[j], in[i]
			}
		}
	}
}

type SortFn struct{}

var SortFunc = SortFn{}

func (SortFn) String(a, b interface{}) int8 {
	var (
		as = a.(string)
		bs = b.(string)
	)

	if as > bs {
		return 1
	}
	return -1
}

func (SortFn) Int(a, b interface{}) int8 {
	var (
		ai = a.(int)
		bi = b.(int)
	)
	if ai > bi {
		return 1
	}
	return -1
}

func (SortFn) Int8(a, b interface{}) int8 {
	var (
		ai = a.(int8)
		bi = b.(int8)
	)
	return ai - bi
}

func (SortFn) Int16(a, b interface{}) int8 {
	var (
		ai = a.(int16)
		bi = b.(int16)
	)
	if ai > bi {
		return 1
	}
	return -1
}

func (SortFn) Int32(a, b interface{}) int8 {
	var (
		ai = a.(int32)
		bi = b.(int32)
	)
	if ai > bi {
		return 1
	}
	return -1
}

func (SortFn) Int64(a, b interface{}) int8 {
	var (
		ai = a.(int64)
		bi = b.(int64)
	)
	if ai > bi {
		return 1
	}
	return -1
}

func (SortFn) Uint(a, b interface{}) int8 {
	var (
		ai = a.(uint)
		bi = b.(uint)
	)
	if ai > bi {
		return 1
	}
	return -1
}

func (SortFn) Uint8(a, b interface{}) int8 {
	var (
		ai = a.(uint8)
		bi = b.(uint8)
	)
	if ai > bi {
		return 1
	}
	return -1
}

func (SortFn) Uint16(a, b interface{}) int8 {
	var (
		ai = a.(uint16)
		bi = b.(uint16)
	)
	if ai > bi {
		return 1
	}
	return -1
}

func (SortFn) Uint32(a, b interface{}) int8 {
	var (
		ai = a.(uint32)
		bi = b.(uint32)
	)
	if ai > bi {
		return 1
	}
	return -1
}

func (SortFn) Uint64(a, b interface{}) int8 {
	var (
		ai = a.(uint64)
		bi = b.(uint64)
	)
	if ai > bi {
		return 1
	}
	return -1
}
