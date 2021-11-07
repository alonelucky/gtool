package arrays

import "testing"

func TestFilter(t *testing.T) {
	var (
		example1 = []int32{2996, 9251, 1596, 9214, 4207, 352, 1912, 9240, 3955, 9754, 3798, 8519, 3794, 2621, 6726, 3747, 128, 6286, 8441, 7888, 9840, 22, 7341}
		example2 = []int64{350, 939, 9768, 4652, 1364, 1466, 9498, 4476, 6907, 3160, 2577, 2377, 9165, 1743, 1138, 701, 3401, 7761, 8027}
		example3 = []int{220, 88, 196, 198, 8, 223, 9, 70, 79, 182, 240, 95, 111, 225, 102, 91, 19, 61, 42, 169, 3, 23, 57, 113, 139, 149, 248, 202, 80, 94, 212, 40, 77, 224, 111, 121, 120, 210, 70, 154, 57, 13, 239, 34, 4, 133, 156, 167, 66, 138}

		out1 []int32
	)

	Filter(example1, &out1, func(v interface{}, i int) bool {
		return v.(int32)%100 > 50
	})

	if Joins(out1, ",") != "2996,9251,1596,352,3955,9754,3798,3794,6286,7888" {
		t.Error(Joins(out1, ","))
	}

	Filter(&example2, nil, func(v interface{}, i int) bool {
		return i > 3
	})
	if Joins(example2, ",") != "1364,1466,9498,4476,6907,3160,2577,2377,9165,1743,1138,701,3401,7761,8027" {
		t.Error(Joins(example2, ","))
	}

	Filter(&example3, nil, func(v interface{}, i int) bool {
		return v.(int) > 200
	})

	if Joins(example3, ",") != "220,223,240,225,248,202,212,224,210,239" {
		t.Error(Joins(example3, ","))
	}

}
