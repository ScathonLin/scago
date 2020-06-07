package queue

import "scago/collections/common"

// Queue defines the actions of queuq.
type Queue interface {
	common.Collection
	// remove and return the first element of queue.
	PollFirst() interface{}
	// remove and return the last element of queue.
	PollLast() interface{}
	// return the first element of queue.
	PeekFirst() interface{}
	// return the last element of queue;
	PeekLast() interface{}
	// insert element to the head of queue.
	OfferFirst(interface{}) bool
	// insert element to the tail of queue.
	OfferLast(interface{}) bool
	// the same as PollFirst.
	RemoveFirst() interface{}
	// the same as PollLast.
	RemoveLast() interface{}
}
