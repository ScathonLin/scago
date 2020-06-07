package common

import (
	"fmt"
	"strings"
)

type Collection interface {
	IsEmpty() bool
	Iterator() Iterator
}

const (
	emptyString = ""
)

//===============BaseCollection Start======================
type BaseCollection struct {
	Collection
	// using pointer to advoid some operation are invalid to the elements array.
	elementData *[]interface{}
	size        int
	cap         int
}

func (bc *BaseCollection) IsEmpty() bool {
	return bc.size == 0 && len(*bc.elementData) == 0
}

func (bc *BaseCollection) SetSize(newSize int) {
	bc.size = newSize
}

func (bc *BaseCollection) SetCap(newCap int) {
	bc.cap = newCap
}

func (bc *BaseCollection) Size() int {
	return bc.size
}

func (bc *BaseCollection) Cap() int {
	return bc.cap
}

func (bc *BaseCollection) Elements() *[]interface{} {
	return bc.elementData
}

func (bc *BaseCollection) ToString() string {
	if bc.size <= 0 {
		return emptyString
	}
	// use high performance way to build string.
	var builder strings.Builder
	for i := 0; i < bc.size; i++ {
		builder.WriteString(fmt.Sprintf("%v", (*bc.elementData)[i]))
		if i != bc.size-1 {
			builder.WriteString(",")
		}
	}
	return builder.String()
}

func (bc *BaseCollection) Iterator() Iterator {
	return NewBaseIterator(bc.Size(), *bc.Elements())
}

func NewBaseCollection(elements []interface{}, size, cap int) *BaseCollection {
	return &BaseCollection{
		elementData: &elements,
		size:        size,
		cap:         cap,
	}
}

//===============BaseCollection End======================

type ArrayLikeCollection struct {
	*BaseCollection
}

type LinkedListLikeCollection struct {
}

type MapLikeCollection struct {
}

type TreeLikeCollection struct {
}

type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type BaseIterator struct {
	cursor      int
	size        int
	elementData []interface{}
}

func (iterator *BaseIterator) HasNext() bool {
	return iterator.cursor >= 0 && iterator.cursor < iterator.size
}

func (iterator *BaseIterator) Next() interface{} {
	iterator.preCheck()
	iterator.cursor++
	return iterator.elementData[iterator.cursor-1]
}

func (iterator *BaseIterator) preCheck() {
	if iterator.cursor < 0 {
		panic("cursor of iterator must not be negative.")
	}
	if iterator.cursor >= iterator.size {
		panic("cursor of iterator has exceed the length of the collection.")
	}
}

func NewBaseIterator(size int, elementData []interface{}) *BaseIterator {
	return &BaseIterator{
		cursor:      0,
		size:        size,
		elementData: elementData,
	}
}
