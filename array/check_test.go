package array

import "testing"

func TestContains(t *testing.T) {
	var (
		lst1 = []string{"123", "234", "456"}
		lst2 = []int32{1, 2, 3}
	)

	t.Error(Contains(lst1, "123"))
	t.Error(Contains(lst1, "223"))
	t.Error(Contains(lst2, 0))
	t.Error(Contains(lst2, 2))
}
