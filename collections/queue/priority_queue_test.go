package queue

import (
	"fmt"
	"testing"
)

var intComparator Comparator = func(a interface{}, b interface{}) int {
	ai := a.(int)
	bi := b.(int)
	return ai - bi
}

func TestNewPriorityQueueWithElems(t *testing.T) {
	arr := [...]int{4, 5, 1, 3, 2}
	input := make([]interface{}, len(arr))
	for i, v := range arr {
		input[i] = v
	}
	q := NewPriorityQueueWithElems(input, intComparator, ASC)
	fmt.Printf("Top is: %v, size = %d\n", q.Top(), q.Size())
	fmt.Printf("Pop is: %v, size = %d\n", q.Pop(), q.Size())
	fmt.Printf("Top is: %v, size = %d\n", q.Top(), q.Size())
	fmt.Println("==============Start pop all elements============")
	size := q.Size()
	for i := 0; i < size; i++ {
		fmt.Printf("Pop is: %v, size = %d\n", q.Pop(), q.Size())
	}
	fmt.Println("==============Over pop all elements============")

	fmt.Println("------------Test Desc Sort-----------")
	q = NewPriorityQueueWithElems(input, intComparator, DESC)
	fmt.Printf("Top is: %v, size = %d\n", q.Top(), q.Size())
	fmt.Printf("Pop is: %v, size = %d\n", q.Pop(), q.Size())
	fmt.Printf("Top is: %v, size = %d\n", q.Top(), q.Size())
	fmt.Println("==============Start pop all elements============")
	size = q.Size()
	for i := 0; i < size; i++ {
		fmt.Printf("Pop is: %v, size = %d\n", q.Pop(), q.Size())
	}
	fmt.Println("==============Over pop all elements============")
}

func TestNewPriorityQueue(t *testing.T) {
	q := NewPriorityQueue(intComparator, ASC)
	q.Push(4)
	q.Push(5)
	q.Push(1)
	q.Push(3)
	q.Push(2)
	fmt.Printf("Top is: %v, size = %d\n", q.Top(), q.Size())
	fmt.Printf("Pop is: %v, size = %d\n", q.Pop(), q.Size())
	fmt.Printf("Top is: %v, size = %d\n", q.Top(), q.Size())
	fmt.Println("==============Start pop all elements============")
	size := q.Size()
	for i := 0; i < size; i++ {
		fmt.Printf("Pop is: %v, size = %d\n", q.Pop(), q.Size())
	}
	fmt.Println("==============Over pop all elements============")
}

func TestNewPriorityQueueWithCap(t *testing.T) {
	q := NewPriorityQueueWithCap(28, intComparator, ASC)
	q.Push(4)
	q.Push(5)
	q.Push(1)
	q.Push(3)
	q.Push(2)
	fmt.Printf("Top is: %v, size = %d\n", q.Top(), q.Size())
	fmt.Printf("Pop is: %v, size = %d\n", q.Pop(), q.Size())
	fmt.Printf("Top is: %v, size = %d\n", q.Top(), q.Size())
	fmt.Println("==============Start pop all elements============")
	size := q.Size()
	for i := 0; i < size; i++ {
		fmt.Printf("Pop is: %v, size = %d\n", q.Pop(), q.Size())
	}
	fmt.Println("==============Over pop all elements============")
}

type task struct {
	taskName string
	priority int
}

var taskPriorityCompartor Comparator = func(t1 interface{}, t2 interface{}) int {
	return t1.(task).priority - t2.(task).priority
}

// test complex object.
func TestPriorityQueue_Pop(t *testing.T) {
	q := NewPriorityQueue(taskPriorityCompartor, DESC)
	q.Push(task{"task1", 10})
	q.Push(task{"task2", 8})
	q.Push(task{"task3", 11})
	q.Push(task{"task4", 9})
	q.Push(task{"task5", 7})
	q.Push(task{"task6", 2})
	q.Push(task{"task7", 4})
	q.Push(task{"task8", 5})
	sz := q.Size()
	fmt.Println("------Print All Elements In The PriorityQueue-------")
	for i := 0; i < sz; i++ {
		fmt.Printf("current head element is %v, size = %d\n", q.Pop(), q.Size())
	}
	fmt.Println("------Over Print All Elements In The PriorityQueue--------")
}

func TestPrioriyQueueExpand(t *testing.T) {
	q := NewPriorityQueue(intComparator, DESC)
	for i := 0; i < 50; i++ {
		q.Push(i)
	}
	fmt.Println("------Pop All elements in the queue.--------")
	sz := q.Size()
	for i := 0; i < sz; i++ {
		fmt.Printf("current head element is : %v, size = %d, cap = %d\n", q.Pop(), q.Size(), q.Cap())
	}
	fmt.Println("------Over Pop All elements in the queue.--------")
	fmt.Println("------Start Test Queue with given elements------")
	arr := make([]interface{}, 50)
	for i := 0; i < 50; i++ {
		arr[i] = i
	}
	q = NewPriorityQueueWithElems(arr, intComparator, DESC)
	sz = q.Size()
	for i := 0; i < sz; i++ {
		fmt.Printf("current head element is : %v, size = %d, cap = %d\n", q.Pop(), q.Size(), q.Cap())
	}
	fmt.Println("------Over Test Queue with given elements------")

}
