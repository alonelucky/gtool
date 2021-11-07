package arrays

import "testing"

func TestContains(t *testing.T) {
	var (
		lst1 = []string{"123", "234", "456"}
		lst2 = []int32{1, 2, 3}
	)

	if ok := Contains(lst1, "123"); !ok {
		t.Error("lst1-123", ok)
	}

	if ok := Contains(lst1, "223"); ok {
		t.Error("lst1-223", ok)
	}

	if ok := Contains(lst2, 0); ok {
		t.Error("lst2-0", ok)
	}

	if ok := Contains(lst2, 2); !ok {
		t.Error("lst2-2", ok)
	}
}
