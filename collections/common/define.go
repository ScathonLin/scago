package common

import (
	"fmt"
	"strings"
)

type Collection interface {
	IsEmpty() bool
	Size() int
	Iterator() Iterator
	ToString() string
}

const (
	emptyString = ""
)

//===============ArrayLikeBaseCollection Start======================
type ArrayLikeBaseCollection struct {
	// using pointer to advoid some operation are invalid to the elements array.
	elementData *[]interface{}
	size        int
	cap         int
}

func (bc *ArrayLikeBaseCollection) IsEmpty() bool {
	return bc.size == 0 && len(*bc.elementData) == 0
}

func (bc *ArrayLikeBaseCollection) SetSize(newSize int) {
	bc.size = newSize
}

func (bc *ArrayLikeBaseCollection) SetCap(newCap int) {
	bc.cap = newCap
}

func (bc *ArrayLikeBaseCollection) Size() int {
	return bc.size
}

func (bc *ArrayLikeBaseCollection) Cap() int {
	return bc.cap
}

func (bc *ArrayLikeBaseCollection) Elements() *[]interface{} {
	return bc.elementData
}

func (bc *ArrayLikeBaseCollection) ToString() string {
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

func (bc *ArrayLikeBaseCollection) Iterator() Iterator {
	return NewBaseIterator(bc.Size(), *bc.Elements())
}

func NewBaseCollection(elements []interface{}, size, cap int) *ArrayLikeBaseCollection {
	return &ArrayLikeBaseCollection{
		elementData: &elements,
		size:        size,
		cap:         cap,
	}
}

//===============ArrayLikeBaseCollection End======================

type ArrayLikeCollection struct {
	*ArrayLikeBaseCollection
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
