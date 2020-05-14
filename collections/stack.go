package collections

import (
	"math"
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
	elements []Ele
	size     int
	cap      int
}

func (stk *stack) IsEmpty() bool {
	return stk.size == 0 && len(stk.elements) == 0
}

func (stk *stack) Size() int {
	return stk.size
}
func (stk *stack) Cap() int {
	return stk.cap
}

func (stk *stack) Pop() Ele {
	ele := stk.elements[stk.size-1]
	stk.elements = stk.elements[:stk.size-1]
	stk.size--
	return ele
}

func (stk *stack) Peek() Ele {
	return stk.elements[stk.size-1]
}

func (stk *stack) Push(ele Ele) {
	if stk.size >= maxCap {
		panic("no more space can allocate for the new ele..the size of stack have reached the max capacity.")
	}
	stk.elements[stk.size] = ele
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
	newspace := make([]Ele, newCap)
	originElements := stk.elements
	// copy origin data to new space.
	for i := 0; i < stk.size; i++ {
		newspace[i] = originElements[i]
	}
	stk.elements = newspace
	stk.cap = newCap
}

func (stk *stack) Foreach(foreach func(Ele)) {
	if stk.IsEmpty() {
		return
	}
	eles := stk.elements
	for i := 0; i < stk.size; i++ {
		foreach(eles[i])
	}
}

func (stk *stack) Map(mapFunc func(Ele) interface{}) []interface{} {
	if stk.IsEmpty() {
		return make([]interface{}, 0)
	}
	eles := stk.elements
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
		builder.WriteString(stk.elements[i].toString())
		if i != stk.Size()-1 {
			builder.WriteString(",")
		}
	}
	return builder.String()
}

// New to make a new stack with defaultcapacity.
func New() *stack {
	stk := &stack{size: 0, elements: make([]Ele, initCapacity), cap: initCapacity}
	return stk
}
