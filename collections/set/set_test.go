package set

import (
	"fmt"
	"testing"
)

func TestNewSet(t *testing.T) {
	set := NewOrdinarySet(16)
	set.Add("1")
	set.Add("2")
	set.Add("2")
	fmt.Println(set.Contains(2))
	fmt.Println(set.Size())
	fmt.Println(set.IsEmpty())
	fmt.Println(set.ToString())
	set.Add("3")
	set.Add("4")
	iter := set.Iterator()
	for iter.HasNext() {
		fmt.Println(iter.Next())
	}
}
