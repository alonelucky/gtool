package reflects

import (
	"reflect"
)

// IsArray 是否是数组/切片类型
//
// Is it array / slice type
func IsArray(v reflect.Value) bool {
	k := v.Kind()
	return k == reflect.Array || k == reflect.Slice
}

// Indirect 循环获取最终类型
//
// Loop to get final type
func Indirect(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v
}

// IsInt 是否是整型
func IsInt(v reflect.Value) bool {
	k := v.Kind()
	return k > 1 && k < 7
}

// IsUnit 是否是Unit类型
func IsUnit(v reflect.Value) bool {
	k := v.Kind()
	return k > 6 && k < 12
}

// IsFloat 是否是浮点型
func IsFloat(v reflect.Value) bool {
	k := v.Kind()
	return k == reflect.Float32 || k == reflect.Float64
}

// IsNumber 是否是数字类型, 返回当前的数字类型
func IsNumber(v reflect.Value) reflect.Kind {
	if IsInt(v) || IsFloat(v) || IsUnit(v) {
		return v.Kind()
	}
	return reflect.Invalid
}

// SameKind 判断是否是同类型不同字节的
func SameKind(a, b reflect.Value) bool {
	var (
		ak = Indirect(a).Kind()
		bk = Indirect(b).Kind()
	)

	if (ak > 1 && ak < 7) && (bk > 1 && bk < 7) {
		return true
	}

	if (ak > 6 && ak < 12) && (bk > 6 && bk < 12) {
		return true
	}

	if (ak == 13 || ak == 14) && (bk == 13 || bk == 14) {
		return true
	}

	return ak == bk
}

// Comparable 可比较的类型
//
// 默认原始类型支持比较 int/unit/float/string
func Comparable(v reflect.Value) bool {
	k := v.Kind()
	return (k > 1 && k < 14) || k == 24
}
