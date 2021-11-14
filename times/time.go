package times

import (
	"reflect"
	"time"

	"github.com/alonelucky/gtool/reflects"
)

const (
	DefaultFormat       = "2006-01-02 15:04:05" // 格式化时间: 年-月-日 时:分:秒(24小时制)
	DefaultFormat12     = "2006-01-02 03:04:05" // 格式化时间: 年-月-日 时:分:秒(12小时制)
	DefaultFormatDate   = "2006-01-02"
	DefaultFormatTime   = "15:04:05" // 格式化时间: 时:分:秒(24小时制)
	DefaultFormatTime12 = "03:04:05" // 格式化时间: 时:分:秒(12小时制)
)

const (
	DaySecond int64 = 3600 * 24
)

type Times struct {
	t time.Time
}

// Now is replace time.Now()
//
// without params returns current time
// with int type will parse time.Unix
// with string type will use defaultformat
// with time.Time will use time
func Now(times ...interface{}) (t Times) {
	t = Times{}
	if l := len(times); l > 0 {
		f := times[0]

		if v, ok := f.(string); ok {
			// 获取指定模板解析
			var fs string = DefaultFormat
			if l > 1 {
				fs, _ = times[1].(string)
			}

			now, e := time.Parse(fs, v)
			if e != nil {
				now = time.Now()
			}
			t.t = now
			return t
		}

		if v, ok := f.(time.Time); ok {
			t.t = v
			return t
		}

		if v := reflect.ValueOf(f); reflects.IsInt(v) {
			t.t = time.Unix(v.Int(), 0)
			return
		}

	}

	t.t = time.Now()
	return
}

func (t Times) Format(formats ...string) (s string) {
	var fs string = DefaultFormat
	if len(formats) > 0 {
		fs = formats[0]
	}

	return t.t.Format(string(fs))
}

func (t Times) Time() time.Time {
	return t.t
}

func (t Times) Between(start, end time.Time) bool {
	return t.t.Before(start) && t.t.After(end)
}

func (t Times) Before(n time.Time) bool {
	return t.t.Before(n)
}

func (t Times) After(n time.Time) bool {
	return t.t.After(n)
}

// 获取当天的开始时间
func (t Times) DayStart(n time.Time) Times {
	s := t.SecondOfDay()
	return t.Second(-1 * s)
}

func (t Times) DayEnd(n time.Time) Times {
	s := t.SecondOfDay()
	return t.Second(DaySecond - s)
}

// UTC 时间
func (t Times) UTC() time.Time {
	return t.t.UTC()
}

// Unix 时间戳,秒级
func (t Times) Unix() int64 {
	return t.t.Unix()
}

func (t Times) Second(i int64) Times {
	return t.Add(time.Duration(i) * time.Second)
}

func (t Times) Minute(i int64) Times {
	return t.Add(time.Duration(i) * time.Minute)
}

func (t Times) Hour(i int64) Times {
	return t.Add(time.Duration(i) * time.Hour)
}

func (t Times) Date(i int64) Times {
	return t.Add(time.Duration(i*24) * time.Hour)
}

func (t Times) Week(i int64) Times {
	return t.Add(time.Duration(i*24) * time.Hour)
}

// DayOfYear 今天在一年中的位置
func (t Times) DayOfYear() int64 {
	return int64(t.t.YearDay())
}

// IsToday 是否是今天
func (t Times) IsToday(n time.Time) bool {
	y, m, d := t.t.Date()
	Y, M, D := n.Date()
	return y == Y && m == M && d == D
}

// Add time.Time{}.Add
func (t Times) Add(d time.Duration) Times {
	return Times{t: t.t.Add(d)}
}

// Difference 两个时间之间的微秒差距
func (t Times) Difference(d time.Time) int64 {
	return t.t.Sub(d).Microseconds()
}

// Diffdays 两个时间相差的天数,毫秒级, 浮点型
func (t Times) Diffdays(d time.Time) float64 {
	return float64(t.t.Sub(d).Milliseconds()) / 1000 * 3600
}

// DaysOfMonth is specifies the total number of days included in the month.
//
// 指定月份包含的天数
func (t Times) DaysOfMonth(m time.Month) int {
	switch m {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 4, 6, 9, 11:
		return 30
	default:
	}

	y := t.t.Year()

	if y%100 > 0 && y%4 == 0 {
		return 29
	}

	if y%400 == 0 {
		return 29
	}

	return 28
}

// 获取当前秒是一天中的多少秒
func (t Times) SecondOfDay() int64 {
	h, m, s := t.t.Clock()
	return int64(h*3600 + m*60 + s)
}

// MonthDays is total number of days in the current month
//
// 当前月份包含的总天数
func (t Times) MonthDays() int {
	m := t.t.Month()
	return t.DaysOfMonth(m)
}
