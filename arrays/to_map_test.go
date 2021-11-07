package arrays

import (
	"testing"
)

func TestToMap(t *testing.T) {
	var (
		lst = []string{"1", "2", "3"}
		out = make(map[int]string)
	)

	ToMap(lst, out, func(item interface{}, i int) (k, v interface{}) {
		return i, item
	})
	t.Log(out)
}

func TestColumn(t *testing.T) {
	type (
		Abc struct {
			Name  string
			Age   int
			Class string
		}
	)

	var (
		lst1 = []*Abc{
			{Name: "123"},
			{Name: "234"},
			{Name: "345"},
			{Name: "456"},
			{Name: "567"},
		}
		mm10 = make(map[string]*Abc)
		mm11 = make(map[int]*Abc)
	)

	Column(lst1, &mm10, "Name")
	Column(lst1, &mm11, "Age")
}
