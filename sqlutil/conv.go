package sqlutil

// Int32SliceToInt64Slice converts int32 array to int64 array
func Int32SliceToInt64Slice(a []int32) []int64 {
	res := make([]int64, 0, len(a))
	for _, v := range a {
		res = append(res, int64(v))
	}
	return res
}

// Int64SliceToAbstractSlice converts int64 array to []interface{}
func Int64SliceToAbstractSlice(a []int64) []interface{} {
	res := []interface{}{}
	for _, v := range a {
		res = append(res, v)
	}
	return res
}
