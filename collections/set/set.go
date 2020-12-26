package set

import (
	"fmt"
	"scago/collections/common"
	"strings"
)

type Set interface {
	IsEmpty() bool
	Size() int
	Add(interface{})
	Remove(interface{}) interface{}
	ToString() string
	Contains(interface{}) bool
	ContainsAll([]interface{}) bool
}

type AbstractSet struct {
	elems map[interface{}]interface{}
}

func (s *AbstractSet) ToString() string {
	if s.elems == nil {
		panic("set is not initialized.")
	}
	if len(s.elems) <= 0 {
		return "set:[]"
	}
	// use high performance way to build string.
	var builder strings.Builder
	keys := make([]string, len(s.elems))
	i := 0
	for k, _ := range s.elems {
		keys[i] = fmt.Sprintf("%v", k)
		i++
	}
	builder.WriteString(strings.Join(keys, ","))
	return builder.String()
}

type ordinarySet struct {
	AbstractSet
}

func NewOrdinarySet(capacity int) *ordinarySet {
	s := new(ordinarySet)
	s.elems = make(map[interface{}]interface{}, capacity)
	return s
}

func (s *ordinarySet) Add(elem interface{}) {
	s.elems[elem] = true
}

func (s *ordinarySet) Contains(elem interface{}) bool {
	_, existed := s.elems[elem]
	return existed
}

func (s *ordinarySet) Remove(elem interface{}) interface{} {
	delete(s.elems, elem)
	return elem
}
func (s *ordinarySet) IsEmpty() bool {
	return s.elems == nil || len(s.elems) == 0
}

func (s *ordinarySet) Size() int {
	return len(s.elems)
}

func (s *ordinarySet) ContainsAll(elemsToCheck []interface{}) bool {
	for _, ele := range elemsToCheck {
		if _, ok := s.elems[ele]; !ok {
			return false
		}
	}
	return true
}
