package covert

import "testing"

func TestInt64SliceToMap(t *testing.T) {
	var int64Slice = []int64{1, 2, 3, 4, 5}
	var int64Map = Int64SliceToMap(int64Slice)
	for _, v := range int64Slice {
		if _, ok := int64Map[v]; !ok {
			t.Errorf("int64Map[%d] is not exist", v)
		}
	}
}
