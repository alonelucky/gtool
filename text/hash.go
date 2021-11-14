package text

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
)

const (
	_md5 int8 = iota + 1
	_sha1
	_sha256
	_sha512
)

func _hash(tp int8, src []byte, salt ...[]byte) string {
	var h hash.Hash
	switch tp {
	case _md5:
		h = md5.New()
	case _sha1:
		h = sha1.New()
	case _sha256:
		h = sha256.New()
	case _sha512:
		h = sha512.New()
	}

	h.Write(src)

	if len(salt) > 0 {
		h.Write(salt[0])
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}

func MD5(src []byte, salt ...[]byte) string {
	return _hash(_md5, src, salt...)
}

func Sha1(src []byte, salt ...[]byte) string {
	return _hash(_sha1, src, salt...)
}

func Sha256(src []byte, salt ...[]byte) string {
	return _hash(_sha256, src, salt...)
}

func Sha512(src []byte, salt ...[]byte) string {
	return _hash(_sha512, src, salt...)
}
