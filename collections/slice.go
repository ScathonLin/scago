package collections

type Slice []interface{}

func (slice Slice) Foreach(fn func(interface{})) {
	if slice == nil || len(slice) == 0 {
		return
	}
	for _, e := range slice {
		fn(e)
	}
}

func (slice Slice) Map(fn func(interface{}) interface{}) Slice {
	var resultSlice Slice = make([]interface{}, len(slice))
	for index, ele := range slice {
		resultSlice[index] = fn(ele)
	}
	return resultSlice
}

func (slice Slice) Filter(fn func(interface{}) bool) Slice {
	var finalResult Slice = make([]interface{}, 0)
	for _, ele := range slice {
		if fn(ele) {
			finalResult = append(finalResult, ele)
		}
	}
	return finalResult
}
