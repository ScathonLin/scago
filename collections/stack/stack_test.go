package collections

import (
	"fmt"
	"strconv"
	"testing"
)

func TestNew(t *testing.T) {
	stack := NewStack()
	fmt.Println("Test New Stack...")
	fmt.Printf("stack size is : %d\n", stack.Size())
	fmt.Printf("stack cap is: %d\n", stack.Cap())
}

func TestBasicOperation(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	fmt.Println("stack size is : " + strconv.Itoa(stack.Size()))
	fmt.Println("stack cap is : " + strconv.Itoa(stack.Cap()))
	fmt.Println(stack.ToString())

	fmt.Println("===========================")

	fmt.Printf("peek is : %v\n", stack.Peek())
	fmt.Printf("pop one ele is : %v\n", stack.Pop())
	fmt.Printf("pop two ele is : %v\n", stack.Pop())
	fmt.Printf("stack size is : %s\n", strconv.Itoa(stack.Size()))
	fmt.Printf("stack cap is : %v\n", strconv.Itoa(stack.Cap()))
	fmt.Printf("peek is : %v\n", stack.Peek())
	fmt.Println(stack.ToString())

}

func TestResize(t *testing.T) {
	stack := NewStack()
	for i := 0; i < 13; i++ {
		stack.Push(i + 1)
	}
	fmt.Println("stack size is : " + strconv.Itoa(stack.Size()))
	fmt.Println("stack cap is : " + strconv.Itoa(stack.Cap()))
	fmt.Println(stack.ToString())
	fmt.Println("===========================")

	fmt.Printf("peek is : %v\n", stack.Peek())
	fmt.Printf("pop one ele is : %v\n", stack.Pop())
	fmt.Printf("pop two ele is : %v\n", stack.Pop())
	fmt.Printf("stack size is : %s\n", strconv.Itoa(stack.Size()))
	fmt.Printf("stack cap is : %v\n", strconv.Itoa(stack.Cap()))
	fmt.Printf("peek is : %v\n", stack.Peek())
	fmt.Println(stack.ToString())
}

func TestForeach(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Foreach(func(ele interface{}) {
		fmt.Println(ele)
	})
}

func TestMap(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	results := stack.Map(func(ele interface{}) interface{} {
		return ele
	})
	for _, res := range results {
		fmt.Printf("%v\n", res)
	}
}

func TestStack_Iterator(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	iterator := stack.Iterator()
	for iterator.HasNext() {
		fmt.Println(iterator.Next())
	}
}
func TestSlice(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	arr = append(arr, make([]int, 3, 3)...)
	fmt.Printf("%v\n", arr)
}
