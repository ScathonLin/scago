package queue

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

type arrQueue struct {
	alc *common.ArrayLikeCollection
}

// NewArrayQueue to make a new stack with defaultcapacity.
func NewArrayQueue() *arrQueue {
	alc := &common.ArrayLikeCollection{
		BaseCollection: common.NewBaseCollection(
			make([]interface{}, initCapacity),
			0,
			initCapacity),
	}
	stk := &arrQueue{alc: alc}
	return stk
}

// remove and return the first element of queue.
func (queue *arrQueue) PollFirst() interface{} {
	queue.preCheck()
	alc := queue.alc
	alc.SetSize(alc.Size() - 1)
	eles := *alc.Elements()
	ele := eles[0]
	alc.SetElements(eles[1:])
	return ele
}

// remove and return the last element of queue.
func (queue *arrQueue) PollLast() interface{} {
	queue.preCheck()
	alc := queue.alc
	alc.SetSize(alc.Size() - 1)
	ele := (*alc.Elements())[alc.Size()]
	(*alc.Elements())[alc.Size()] = nil
	return ele
}

// return the first element of queue.
func (queue *arrQueue) PeekFirst() interface{} {
	queue.preCheck()
	return (*queue.alc.Elements())[0]
}

// return the last element of queue;
func (queue *arrQueue) PeekLast() interface{} {
	queue.preCheck()
	return (*queue.alc.Elements())[queue.Size()-1]
}

// insert element to the head of queue.
func (queue *arrQueue) OfferFirst(ele interface{}) interface{} {
	alc := queue.alc
	if float32(alc.Size())/float32(alc.GetCap()) >= resizeThreashold {
		queue.resize()
	}
	eles := *alc.Elements()
	eles = append([]interface{}{ele}, eles...)
	alc.SetElements(eles)
	alc.SetSize(alc.Size() + 1)
	return ele
}

// insert element to the tail of queue.
func (queue *arrQueue) OfferLast(ele interface{}) interface{} {
	alc := queue.alc
	if float32(alc.Size())/float32(alc.GetCap()) >= resizeThreashold {
		queue.resize()
	}
	(*alc.Elements())[alc.Size()] = ele
	alc.SetSize(alc.Size() + 1)
	return ele
}

// the same as PollFirst.
func (queue *arrQueue) RemoveFirst() interface{} {
	return queue.PollFirst()
}

// the same as PollLast.
func (queue *arrQueue) RemoveLast() interface{} {
	return queue.PollLast()
}
func (queue *arrQueue) Size() int {
	return queue.alc.Size()
}

func (queue *arrQueue) Cap() int {
	return queue.alc.GetCap()
}

func (queue *arrQueue) IsEmpty() bool {
	return queue.Size() <= 0
}

func (queue *arrQueue) preCheck() {
	if queue.IsEmpty() {
		panic("queue is empty.")
	}
}

func (queue *arrQueue) resize() {
	newCap := queue.alc.GetCap() << 1
	if newCap > maxCap {
		panic("exceed the max capacity of the queue,cannot finish the capacity expanding...")
	}
	eles := *queue.alc.Elements()
	eles = append(eles, make([]interface{}, newCap-queue.alc.GetCap(), newCap-queue.alc.GetCap())...)
	queue.alc.SetCap(newCap)
}

func (queue *arrQueue) ToString() string {
	if queue.alc.Size() <= 0 {
		return emptyString
	}
	// use high performance way to build string.
	var builder strings.Builder
	for i := 0; i < queue.Size(); i++ {
		builder.WriteString(fmt.Sprintf("%v", (*queue.alc.Elements())[i]))
		if i != queue.Size()-1 {
			builder.WriteString(",")
		}
	}
	return builder.String()
}
