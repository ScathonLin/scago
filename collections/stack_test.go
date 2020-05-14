package collections

import (
	"fmt"
	"strconv"
	"testing"
)

func TestNew(t *testing.T) {
	stack := New()
	fmt.Println("Test New Stack...")
	fmt.Printf("stack size is : %d\n", stack.Size())
	fmt.Printf("stack cap is: %d\n", stack.Cap())
}

func TestBasicOperation(t *testing.T) {
	stack := New()
	stack.Push(Ele{value: 1})
	stack.Push(Ele{value: 2})
	stack.Push(Ele{value: 3})
	stack.Push(Ele{value: 4})
	stack.Push(Ele{value: 5})
	fmt.Println("stack size is : " + strconv.Itoa(stack.Size()))
	fmt.Println("stack cap is : " + strconv.Itoa(stack.Cap()))
	fmt.Println(stack.ToString())

	fmt.Println("===========================")

	fmt.Println("peek is : " + stack.Peek().toString())
	fmt.Println("pop one ele is : " + stack.Pop().toString())
	fmt.Println("pop two ele is : " + stack.Pop().toString())
	fmt.Println("stack size is : " + strconv.Itoa(stack.Size()))
	fmt.Println("stack cap is : " + strconv.Itoa(stack.Cap()))
	fmt.Println("peek is : " + stack.Peek().toString())
	fmt.Println(stack.ToString())

}

func TestResize(t *testing.T) {
	stack := New()
	for i := 0; i < 13; i++ {
		stack.Push(Ele{value: i + 1})
	}
	fmt.Println("stack size is : " + strconv.Itoa(stack.Size()))
	fmt.Println("stack cap is : " + strconv.Itoa(stack.Cap()))
	fmt.Println(stack.ToString())

	fmt.Println("===========================")

	fmt.Println("peek is : " + stack.Peek().toString())
	fmt.Println("pop one ele is : " + stack.Pop().toString())
	fmt.Println("pop two ele is : " + stack.Pop().toString())
	fmt.Println("stack size is : " + strconv.Itoa(stack.Size()))
	fmt.Println("stack cap is : " + strconv.Itoa(stack.Cap()))
	fmt.Println("peek is : " + stack.Peek().toString())
	fmt.Println(stack.ToString())
}

func TestForeach(t *testing.T) {
	stack := New()
	stack.Push(Ele{value: 1})
	stack.Push(Ele{value: 2})
	stack.Push(Ele{value: 3})
	stack.Push(Ele{value: 4})
	stack.Push(Ele{value: 5})
	stack.Foreach(func(ele Ele) {
		fmt.Println(ele.toString())
	})
}

func TestMap(t *testing.T) {
	stack := New()
	stack.Push(Ele{value: 1})
	stack.Push(Ele{value: 2})
	stack.Push(Ele{value: 3})
	stack.Push(Ele{value: 4})
	stack.Push(Ele{value: 5})
	results := stack.Map(func(ele Ele) interface{} {
		return ele.toString()
	})
	for _, res := range results {
		fmt.Printf("%v\n", res)
	}
}
