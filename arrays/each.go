package arrays

import (
	"reflect"

	"github.com/alonelucky/gtool/reflects"
)

// Each 遍历, 不做返回操作
func Each(in interface{}, fn func(v interface{}, i int)) {
	if in == nil {
		return
	}

	var inv = reflect.ValueOf(in)

	if !reflects.IsArray(inv) {
		return
	}

	l := inv.Len()

	for i := 0; i < l; i++ {
		fn(inv.Index(i), i)
	}
}
