package array

import "testing"

func TestDiff(t *testing.T) {
	var (
		lst1 = []string{"123", "234", "456"}
		lst2 = []string{"234", "567"}
		out  []string
	)

	Diff(&out, lst1, lst2)

	t.Error(out)

}
