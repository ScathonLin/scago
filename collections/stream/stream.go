package stream

// description: streaming process api with funcitonal programming like.
// author: linhuadong.
// create: 2020-12-19.

import (
	"fmt"
	"sort"
)

type Function func(e interface{}) interface{}
type Consumer func(e interface{})
type Supply func() interface{}
type Predicate func(e interface{}) bool
type BiFunction func(e1, e2 interface{}) interface{}
type Compare func(e1, e2 interface{}) int

type stream struct {
	head       *stream
	next       *stream
	dataset    []interface{}
	res        interface{}
	actionFunc func()
}

//Foreach visits each element in the slice.
func (strm *stream) Foreach(consumerFunc Consumer) {
	strm.action()
	for _, elem := range strm.head.dataset {
		consumerFunc(elem)
	}
}

//Map element to a new elment which be transfered by mapFunc.
func (strm *stream) Map(mapFunc Function) *stream {
	nextStream := new(stream)
	strm.next = nextStream
	nextStream.head = strm.head
	nextStream.actionFunc = func() {
		ds := strm.head.dataset
		newDataSet := make([]interface{}, len(ds), cap(ds))
		for i, elem := range ds {
			newDataSet[i] = mapFunc(elem)
		}
		strm.head.dataset = newDataSet
	}
	return nextStream
}

func (strm *stream) CollectToMap(keyFunc Function, valFunc Function) map[interface{}]interface{} {
	strm.action()
	ds := strm.head.dataset
	res := make(map[interface{}]interface{}, len(ds))
	for _, e := range ds {
		key := keyFunc(e)
		if _, ok := res[key]; ok {
			panic("duplicate key of: " + fmt.Sprintf("%v", key))
		}
		res[key] = valFunc(e)
	}
	return res
}

//Filter elements those match the filter condition
func (strm *stream) Filter(fn Predicate) *stream {
	nextStream := new(stream)
	strm.next = nextStream
	nextStream.head = strm.head
	nextStream.actionFunc = func() {
		ds := strm.head.dataset
		back := make([]interface{}, len(ds), cap(ds))
		copy(back, ds)
		cursor := 0
		for _, elem := range back {
			if fn(elem) {
				back[cursor] = elem
				cursor++
			}
		}
		back = back[:cursor]
		for i := cursor; i < len(back); i++ {
			back[i] = nil
		}
		strm.head.dataset = back
	}
	return nextStream
}

//Collect all elements after streaming processing.
func (strm *stream) Collect() []interface{} {
	strm.action()
	return strm.head.dataset
}

//Count the elements after streaming processing.
func (strm *stream) Count() int {
	strm.action()
	return len(strm.head.dataset)
}

//Distinct the elements after streaming processing.
func (strm *stream) Distinct(compare Compare) []interface{} {
	strm.action()
	ds := strm.head.dataset
	res := make([]interface{}, len(ds), cap(ds))
	copy(res, ds)
	sort.Slice(res, func(i, j int) bool {
		return compare(res[i], res[j]) < 0
	})
	cursor := 0
	for i, e := range res {
		if i == 0 || compare(e, res[i-1]) != 0 {
			res[cursor] = e
			cursor++
		}
	}
	return res[:cursor]
}

// action trigger the actual data processing.
func (strm *stream) action() {
	processNode := strm.head.next
	for processNode != nil {
		processNode.actionFunc()
		processNode = processNode.next
	}
}

//NewStream create a stream handle with given dataset.
func NewStream(dataset []interface{}) *stream {
	strm := new(stream)
	strm.dataset = dataset
	strm.head = strm
	return strm
}
