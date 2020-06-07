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
	alc *common.ArrayLikeCollection
}

func (stk *stack) IsEmpty() bool {
	return stk.Size() <= 0
}

func (stk *stack) Size() int {
	return stk.alc.Size()
}
func (stk *stack) Cap() int {
	return stk.alc.Cap()
}

func (stk *stack) Pop() interface{} {
	alc := stk.alc
	eles := alc.Elements()
	ele := (*eles)[alc.Size()-1]
	*eles = (*eles)[:alc.Size()-1]
	alc.SetSize(alc.Size() - 1)
	return ele
}

func (stk *stack) Peek() interface{} {
	alc := stk.alc
	eles := alc.Elements()
	return (*eles)[alc.Size()-1]
}

func (stk *stack) Push(ele interface{}) {
	alc := stk.alc
	if alc.Size() >= maxCap {
		panic("no more space can allocate for the new ele..the size of stack have reached the max capacity.")
	}
	eles := alc.Elements()
	(*eles)[alc.Size()] = ele
	alc.SetSize(alc.Size() + 1)
	if float32(alc.Size())/float32(alc.Cap()) >= resizeThreashold {
		stk.resize()
	}
}

func (stk *stack) resize() {
	// allocate new space.
	alc := stk.alc
	newCap := alc.Cap() << 1
	if newCap > maxCap {
		newCap = maxCap
	}
	newspace := make([]interface{}, newCap)
	eles := alc.Elements()
	originElements := alc.Elements()
	// copy origin data to new space.
	for i := 0; i < alc.Size(); i++ {
		newspace[i] = (*originElements)[i]
	}
	*eles = newspace
	alc.SetCap(newCap)
}

func (stk *stack) Foreach(foreach func(interface{})) {
	if stk.IsEmpty() {
		return
	}
	alc := stk.alc
	eles := alc.Elements()
	for i := 0; i < alc.Size(); i++ {
		foreach((*eles)[i])
	}
}

func (stk *stack) Map(mapFunc func(interface{}) interface{}) []interface{} {
	if stk.IsEmpty() {
		return make([]interface{}, 0)
	}
	eles := stk.alc.Elements()
	results := make([]interface{}, stk.alc.Size())
	for i := 0; i < stk.alc.Size(); i++ {
		results[i] = mapFunc((*eles)[i])
	}
	return results
}

func (stk *stack) ToString() string {
	if stk.alc.Size() <= 0 {
		return emptyString
	}
	// use high performance way to build string.
	var builder strings.Builder
	for i := 0; i < stk.Size(); i++ {
		builder.WriteString(fmt.Sprintf("%v", (*stk.alc.Elements())[i]))
		if i != stk.Size()-1 {
			builder.WriteString(",")
		}
	}
	return builder.String()
}

// NewStack to make a new stack with defaultcapacity.
func NewStack() *stack {
	alc := &common.ArrayLikeCollection{
		BaseCollection: common.NewBaseCollection(
			make([]interface{}, initCapacity),
			0,
			initCapacity),
	}
	stk := &stack{alc: alc}
	return stk
}

func (stk *stack) Iterator() common.Iterator {
	return stk.alc.Iterator()
}
