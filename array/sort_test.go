package array

import "testing"

func TestSort(t *testing.T) {
	var (
		string0 = []string{"qonyrl67tz", "7n8jua8m9f3", "f2nhl3hz10i", "oac3136dj1", "pnrgz9u0rc", "7p7xnld3416"}
		string1 = []string{"1vvymufbbth", "dedqza1d0wa"}

		example0 = []int16{2872, 1960, 9718, 9440, 3663, 2773, 2672, 5733, 1494, 4435, 6072, 1817, 4115, 4042, 4069, 2349, 5383, 3318, 2286, 7137}
		example1 = []int32{2996, 9251, 1596, 9214, 4207, 352, 1912, 9240, 3955, 9754, 3798, 8519, 3794, 2621, 6726, 3747, 128, 6286, 8441, 7888, 9840, 22, 7341}
		example2 = []int64{350, 939, 9768, 4652, 1364, 1466, 9498, 4476, 6907, 3160, 2577, 2377, 9165, 1743, 1138, 701, 3401, 7761, 8027}
		example3 = []int{220, 88, 196, 198, 8, 223, 9, 70, 79, 182, 240, 95, 111, 225, 102, 91, 19, 61, 42, 169, 3, 23, 57, 113, 139, 149, 248, 202, 80, 94, 212, 40, 77, 224, 111, 121, 120, 210, 70, 154, 57, 13, 239, 34, 4, 133, 156, 167, 66, 138}
		example4 = []int8{28, 60, 1}

		out1 []string
		out2 []int32
	)

	Sort(example0, nil, SortFunc.Int16)
	if Joins(example0, ",") != "1494,1817,1960,2286,2349,2672,2773,2872,3318,3663,4042,4069,4115,4435,5383,5733,6072,9440,9718,7137" {
		t.Error("example0", example0)
	}

	Sort(example1, &out2, SortFunc.Int32)
	if Joins(out2, ",") != "22,128,352,1596,1912,2621,2996,3747,3794,3798,3955,4207,6286,6726,7888,8441,8519,9214,9240,9251,9754,9840,7341" {
		t.Error("example1", example1)
	}

	Sort(example2, nil, SortFunc.Int64)
	if Joins(example2, ",") != "350,701,939,1138,1364,1466,1743,2377,2577,3160,3401,4476,4652,6907,7761,9165,9498,9768,8027" {
		t.Error("example2", example2)
	}

	Sort(example3, nil, SortFunc.Int)
	if Joins(example3, ",") != "3,4,8,9,13,19,23,34,40,42,57,57,61,66,70,70,77,79,80,88,91,94,95,102,111,111,113,120,121,133,139,149,154,156,167,169,182,196,198,202,210,212,220,223,224,225,239,240,248,138" {
		t.Error("example3", example3)
	}

	Sort(example4, nil, SortFunc.Int8)
	if Joins(example4, ",") != "28,60,1" {
		t.Error("example4", example4)
	}

	Sort(string0, nil, SortFunc.String)
	if Joins(string0, ",") != "7n8jua8m9f3,f2nhl3hz10i,oac3136dj1,pnrgz9u0rc,qonyrl67tz,7p7xnld3416" {
		t.Error("string0", Joins(string0, ","))
	}

	Sort(string1, nil, SortFunc.String)
	if Joins(string1, ",") != "1vvymufbbth,dedqza1d0wa" {
		t.Error("string1", Joins(string1, ","))
	}

	Sort(string1, &out1, SortFunc.String)
	if Joins(out1, ",") != "1vvymufbbth,dedqza1d0wa" {
		t.Error("string1-out1", Joins(out1, ","))
	}
}
