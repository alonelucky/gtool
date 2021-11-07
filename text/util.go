package text

import (
	"unicode/utf16"
	"unicode/utf8"
	"unsafe"
)

type charType string

const (
	CharTypeASCII charType = "ascii"
	CharTypeUTF8           = "utf8"
	CharTypeUTF16          = "utf16"
	CharTypeGBK            = "gbk"
)

func NewCharType(tp string) charType {
	return charType(tp)
}

func Length(src string, tps ...charType) int {
	var tp charType = CharTypeUTF8
	if len(tps) > 0 {
		tp = tps[0]
	}

	switch tp {
	case CharTypeASCII:
		return len(src)
	case CharTypeUTF8:
		return utf8.RuneCountInString(src)
	case CharTypeUTF16:
		return len(utf16.Encode([]rune(src)))
	default:
		return len(src)
	}
}

func Bytes(src string) []byte {
	return *(*[]byte)(unsafe.Pointer(&src))
}

func String(bt []byte) string {
	return *(*string)(unsafe.Pointer(&bt))
}
