package arrays

import "testing"

func TestDiff(t *testing.T) {
	var (
		lst1 = []string{"123", "234", "456"}
		lst2 = []string{"234", "567"}
		out  []string
	)

	Diff(&out, lst1, lst2)
	if Joins(out, ",") != "123,456" {
		t.Error(out)
	}
}
