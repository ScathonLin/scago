package queue

import (
	"math"
	"scago/collections/common"
)

const (
	initCapacity int = 16
	maxCap       int = math.MaxInt32
	// Expansion factor
	resizeThreashold float32 = 0.8
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
	queue := &arrQueue{alc: alc}
	return queue
}

// remove and return the first element of queue.
func (queue *arrQueue) PollFirst() interface{} {
	queue.preCheck()
	alc := queue.alc
	alc.SetSize(alc.Size() - 1)
	eles := alc.Elements()
	ele := (*eles)[0]
	*eles = (*eles)[1:]
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
	if float32(alc.Size())/float32(alc.Cap()) >= resizeThreashold {
		queue.resize()
	}
	eles := alc.Elements()
	*eles = append([]interface{}{ele}, *eles...)
	alc.SetSize(alc.Size() + 1)
	return ele
}

// insert element to the tail of queue.
func (queue *arrQueue) OfferLast(ele interface{}) interface{} {
	alc := queue.alc
	if float32(alc.Size())/float32(alc.Cap()) >= resizeThreashold {
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
	return queue.alc.Cap()
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
	newCap := queue.alc.Cap() << 1
	if newCap > maxCap {
		panic("exceed the max capacity of the queue,cannot finish the capacity expanding...")
	}
	eles := *queue.alc.Elements()
	eles = append(eles, make([]interface{}, newCap-queue.alc.Cap(), newCap-queue.alc.Cap())...)
	queue.alc.SetCap(newCap)
}

func (queue *arrQueue) ToString() string {
	return queue.alc.ToString()
}
func (queue *arrQueue) Iterator() common.Iterator {
	return queue.alc.Iterator()
}
