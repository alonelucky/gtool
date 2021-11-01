package text

import "regexp"

var cache = make(map[string]*regexp.Regexp)

var email = regexp.MustCompile("")

func IsEmail(src string) bool {
	return email.MatchString(src)
}

func IsUUID(src string) bool {

}

func IsURL(src string) bool {

}

func IsHex(src string) bool {

}

func IsBase64(src string) bool {

}

func IsIPV4(src string) bool {

}

func IsIPV6(src string) bool {

}

func IsLowerCase(src string) bool {

}

func IsUpperCase(src string) bool {

}

func IsSnakCase(src string) bool {

}
