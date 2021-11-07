package numbers

import (
	"fmt"
	"math"
	"reflect"
	"strconv"

	"github.com/alonelucky/gtool/reflects"
)

// RoundAll 支持任意数字类型保留两位小数, 包括
func RoundAll(n interface{}, pos ...int) float64 {
	var (
		nv = reflect.ValueOf(n)
	)

	if nv.Kind() == reflect.String {
		f, _ := strconv.ParseFloat(n.(string), 64)
		return Round(f, pos...)
	}

	if reflects.IsNumber(nv) == reflect.Invalid {
		return 0
	}

	if reflects.IsInt(nv) {
		return Round(float64(nv.Int()), pos...)
	}

	if reflects.IsUnit(nv) {
		return Round(float64(nv.Uint()), pos...)
	}

	if reflects.IsFloat(nv) {
		return Round(float64(nv.Float()), pos...)
	}

	return 0
}

// ToFixed 保留两位小数返回字符串类型
func ToFixed(n float64, pos ...int) string {
	var p = 2
	if len(pos) > 0 {
		p = pos[0]
	}
	return fmt.Sprintf(`%.`+strconv.FormatInt(int64(p), 10)+`f`, Round(n, pos...))
}

// Round 通用Round方法, 实现四舍五入
func Round(n float64, pos ...int) float64 {
	var p float64
	if len(pos) > 0 {
		p = math.Pow10(pos[0])
	}

	return math.Round(n*p) / p
}

// BankRound 银行家算法, 四舍六入五成双, (奇数进位,偶数舍去)
//
// Banker's algorithm, rounded to double
func BankRound(n float64, pos ...int) float64 {
	var p float64
	if len(pos) > 0 {
		// 增加一位用于判断5成双
		p = math.Pow10(pos[0])
	}

	x := math.Floor(n * p * 10)
	if z := uint64(x) % 2; z == 1 {
		// 进一位
		x += 10
	}

	return math.Floor(x/10) / p
}

// Round32 32bits 的四舍五入算法
func Round32(n float32, pos ...int) float32 {
	return float32(Round(float64(n), pos...))
}

func MinInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}
