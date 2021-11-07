package arrays

import (
	"testing"
)

func TestJoins(t *testing.T) {
	var (
		lst  = []string{"1", "2", "3"}
		lst2 = []int{1, 2, 3}
	)

	var s = Joins(lst, "")
	if s != "123" {
		t.Error(s)
	}

	var s2 = Joins(lst, ",")
	if s2 != "1,2,3" {
		t.Error(s2)
	}

	var s3 = Joins(lst2, ",")
	if s3 != "1,2,3" {
		t.Error(s3)
	}
}
