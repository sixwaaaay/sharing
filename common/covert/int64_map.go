package covert

// Int64SliceToMap covert []int64 to map[int64]struct{}
func Int64SliceToMap(s []int64) map[int64]struct{} {
	m := make(map[int64]struct{})
	for _, v := range s {
		m[v] = struct{}{}
	}
	return m
}
