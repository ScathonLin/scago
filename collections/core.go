package collections

import "fmt"

// Ele is the element in collection
type Ele struct {
	Value interface{}
	index int
}

func (ele Ele) toString() string {
	return fmt.Sprintf("[value:%v]", ele.Value)
}
