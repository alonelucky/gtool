package arrays

import (
	"fmt"
	"strings"
	"testing"
)

func TestMap(t *testing.T) {
	var (
		lst = []string{"1", "2", "3"}
		out = []string{}
	)

	Map(lst, &out, func(v interface{}, i int) interface{} {
		return fmt.Sprintf("%v-%d", v, i)
	})

	if strings.Join(out, ",") != "1-0,2-1,3-2" {
		t.Error(out)
	}
}

func TestMapSelf(t *testing.T) {
	var (
		lst  = []string{"1", "2", "3"}
		lst2 = []int{1, 2, 3}
	)

	t.Log("testing []string")
	MapSelf(&lst, func(v interface{}, i int) interface{} {
		return fmt.Sprintf("%v-%d", v, i)
	})
	if strings.Join(lst, ",") != "1-0,2-1,3-2" {
		t.Error(lst)
	}

	t.Log("testing []int")
	var a = &lst2
	MapSelf(&a, func(v interface{}, i int) interface{} {
		return v.(int) * i
	})

	if Joins(lst2, ",") != "0,2,6" {
		t.Error(lst2)
	}
}
