package queue

import (
	"fmt"
	"scago/collections/common"
	"testing"
)

func TestNewArrayQueue(t *testing.T) {
	queue := NewArrayQueue()
	fmt.Println(queue.IsEmpty())
}

func TestArrQueue_OfferFirst(t *testing.T) {
	q := NewArrayQueue()
	fmt.Printf("Size:%d,cap:%d,content:%v\n", q.Size(), q.Cap(), q.ToString())
	for i := 0; i < 5; i++ {
		q.OfferFirst(i)
	}
	fmt.Printf("Size:%d,cap:%d,content:%v\n", q.Size(), q.Cap(), q.ToString())
	fmt.Printf("Peek First: %v ---> Size:%d,cap:%d,content:%v\n", q.PeekFirst(), q.Size(), q.Cap(), q.ToString())
	fmt.Printf("Poll/Remove First: %v ---> Size:%d,cap:%d,content:%v\n", q.PollFirst(), q.Size(), q.Cap(), q.ToString())
	fmt.Printf("Poll/Remove Last: %v ---> Size:%d,cap:%d,content:%v\n", q.PollLast(), q.Size(), q.Cap(), q.ToString())
	q.OfferFirst(10)
	fmt.Printf("OfferFirst: %v ---> Size:%d,cap:%d,content:%v\n", 10, q.Size(), q.Cap(), q.ToString())
	q.OfferLast(20)
	fmt.Printf("OfferLast: %v ---> Size:%d,cap:%d,content:%v\n", 20, q.Size(), q.Cap(), q.ToString())
	for i := 6; i <= 13; i++ {
		q.OfferFirst(i)
	}
	fmt.Printf("Size:%d,cap:%d,content:%v\n", q.Size(), q.Cap(), q.ToString())
	for i := 14; i < 20; i++ {
		q.OfferLast(i)
	}
	fmt.Printf("Size:%d,cap:%d,content:%v\n", q.Size(), q.Cap(), q.ToString())
	iter := q.Iterator()
	for iter.HasNext() {
		fmt.Printf("%v\n", iter.Next())
	}
	var c common.Collection = q
	println(c.IsEmpty())
	iter = c.Iterator()
	for iter.HasNext() {
		fmt.Printf("%v\n", iter.Next())
	}
}

func Test(t *testing.T) {
	a := getArr()
	change(&a)
	fmt.Printf("%v\n", a)

}

type arr struct {
	a []int
}

func getArr() []int {
	return []int{0}
}
func change(a *[]int) {
	*a = append(*a, 1, 2, 3)
}
