package collections

import (
	"fmt"
	"testing"
)

func TestSlice_Foreach(t *testing.T) {
	var slice Slice = make([]interface{}, 10)
	for i := range slice {
		slice[i] = i
	}

	slice.Foreach(func(item interface{}) {
		fmt.Println(item)
	})
}

func TestSlice_Map(t *testing.T) {
	var slice Slice = make([]interface{}, 10)
	for i := range slice {
		slice[i] = i
	}

	slice.Map(func(ele interface{}) interface{} {
		ele = ele.(int) << 1
		return ele
	}).Foreach(func(elem interface{}) {
		fmt.Println(elem)
	})
}

func TestSlice_Filter(t *testing.T) {
	var slice Slice = make([]interface{}, 10)
	for i := range slice {
		slice[i] = i
	}

	slice.Filter(func(ele interface{}) bool {
		return ele.(int)%2 == 0
	}).Foreach(func(elem interface{}) {
		fmt.Println(elem)
	})
}
