package collections

import (
	"fmt"
	"math"
	"scago/collections/common"
	"strings"
)

const (
	initCapacity int = 16
	maxCap       int = math.MaxInt32
	// Expansion factor
	resizeThreashold float32 = 0.8
	emptyString      string  = ""
)

type stack struct {
	elementData []interface{}
	size        int
	cap         int
}

func (stk *stack) IsEmpty() bool {
	return stk.size == 0 && len(stk.elementData) == 0
}

func (stk *stack) Size() int {
	return stk.size
}
func (stk *stack) Cap() int {
	return stk.cap
}

func (stk *stack) Pop() interface{} {
	ele := stk.elementData[stk.size-1]
	stk.elementData = stk.elementData[:stk.size-1]
	stk.size--
	return ele
}

func (stk *stack) Peek() interface{} {
	return stk.elementData[stk.size-1]
}

func (stk *stack) Push(ele interface{}) {
	if stk.size >= maxCap {
		panic("no more space can allocate for the new ele..the size of stack have reached the max capacity.")
	}
	stk.elementData[stk.size] = ele
	stk.size++
	if float32(stk.size)/float32(stk.cap) >= resizeThreashold {
		stk.resize()
	}
}

func (stk *stack) resize() {
	// allocate new space.
	newCap := stk.cap << 1
	if newCap > maxCap {
		newCap = maxCap
	}
	newspace := make([]interface{}, newCap)
	originElements := stk.elementData
	// copy origin data to new space.
	for i := 0; i < stk.size; i++ {
		newspace[i] = originElements[i]
	}
	stk.elementData = newspace
	stk.cap = newCap
}

func (stk *stack) Foreach(foreach func(interface{})) {
	if stk.IsEmpty() {
		return
	}
	eles := stk.elementData
	for i := 0; i < stk.size; i++ {
		foreach(eles[i])
	}
}

func (stk *stack) Map(mapFunc func(interface{}) interface{}) []interface{} {
	if stk.IsEmpty() {
		return make([]interface{}, 0)
	}
	eles := stk.elementData
	results := make([]interface{}, stk.size)
	for i := 0; i < stk.size; i++ {
		results[i] = mapFunc(eles[i])
	}
	return results
}

func (stk *stack) ToString() string {
	if stk.size <= 0 {
		return emptyString
	}
	// use high performance way to build string.
	var builder strings.Builder
	for i := 0; i < stk.Size(); i++ {
		builder.WriteString(fmt.Sprintf("%v", stk.elementData[i]))
		if i != stk.Size()-1 {
			builder.WriteString(",")
		}
	}
	return builder.String()
}

// NewStack to make a new stack with defaultcapacity.
func NewStack() *stack {
	stk := &stack{size: 0, elementData: make([]interface{}, initCapacity), cap: initCapacity}
	return stk
}

func (stk *stack) Iterator() common.Iterator {
	return common.NewBaseIterator(stk.Size(), stk.elementData)
}
