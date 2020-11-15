package queue

type Order int
type Comparator func(interface{}, interface{}) int

const (
	ASC        Order = 1             // asc sort.
	DESC       Order = -1            // desc sort.
	MaxCap     int   = (1 << 31) - 2 // position with index 0 is not used.
	DefaultCap int   = (1 << 6) + 1  // default cap is 64
)

// priorityQueue struct.
type priorityQueue struct {
	elems      []interface{} // data int priority queue.
	comparator Comparator    // comparator to judge priority.
	order      Order         // big root heap or small root heap.
	size       int           // queue size.
	capacity   int           // the capacity of the queue, max is (1 << 31) - 1
}

func NewPriorityQueue(comparator Comparator, order Order) *priorityQueue {
	if comparator == nil {
		panic("comparator is not specified.")
	}
	elems := make([]interface{}, DefaultCap)
	pq := &priorityQueue{
		elems:      elems,
		comparator: comparator,
		order:      order,
		size:       0,
		capacity:   DefaultCap,
	}
	return pq
}

func NewPriorityQueueWithCap(capacity int, comparator Comparator, order Order) *priorityQueue {
	// must specify a compartor which used to adjust the heap.
	if comparator == nil {
		panic("comparator is not specified.")
	}
	// check the capacity whether exceed the limit.
	if capacity > MaxCap {
		panic("capacity exceed limit, the max cap is (1<<31) - 2")
	}
	elems := make([]interface{}, capacity)
	pq := &priorityQueue{
		elems:      elems,
		comparator: comparator,
		order:      order,
		size:       0,
		capacity:   capacity,
	}
	return pq
}

func NewPriorityQueueWithElems(elems []interface{}, comparator Comparator, order Order) *priorityQueue {
	if len(elems) > MaxCap {
		panic("size of elems exceed the limit, max is (1 <<31 ) - 2")
	}
	// expand more space to storage elements by push operation in future.
	newcap := len(elems) + (len(elems) >> 1)
	if newcap > MaxCap {
		newcap = MaxCap
	}
	newArr := make([]interface{}, newcap)
	for i, v := range elems {
		newArr[i+1] = v
	}
	pq := &priorityQueue{
		elems:      newArr,
		comparator: comparator,
		order:      order,
		size:       len(elems),
		capacity:   newcap,
	}
	pq.adjust()
	return pq
}

//Push a element into priority queue.
func (pq *priorityQueue) Push(ele interface{}) {
	// push operation
	if pq.size >= MaxCap || pq.capacity == MaxCap {
		panic("queue has no space to storage new element")
	}
	// storage expand operation will be triggered when used space exceeds 80% percent of origin allocated space.
	if float32(pq.size*1.0/pq.capacity) > 0.8 {
		pq.expand()
	}
	pq.size++
	pq.elems[pq.size] = ele
	// shiftUp to adjust the heap.
	pq.shiftUp(pq.size)
}

//Pop a element from priority queue.
func (pq *priorityQueue) Pop() interface{} {
	result := pq.elems[1]
	pq.elems[1] = pq.elems[pq.size]
	pq.size--
	pq.shiftDown(1)
	return result
}

//Top returns the root element in the priotity queue, in other words, get the first element in the heap.
func (pq *priorityQueue) Top() interface{} {
	return pq.elems[1]
}

//Size gets the count of elements in the priority queue.
func (pq *priorityQueue) Size() int {
	return pq.size
}

//Cap gets the capacity of the priority queue.
func (pq *priorityQueue) Cap() int {
	return pq.capacity
}

//adjust the heap to maintain the correntness of the heap.
func (pq *priorityQueue) adjust() {
	for i := pq.size >> 1; i > 0; i-- {
		pq.shiftDown(i)
	}
}

//shiftUp is used to float the element up to maintain the correntness of the heap.
func (pq *priorityQueue) shiftUp(k int) {
	prnt := k >> 1
	for prnt > 0 {
		if pq.comparator(pq.elems[k], pq.elems[prnt])*int(pq.order) < 0 {
			pq.elems[k], pq.elems[prnt] = pq.elems[prnt], pq.elems[k]
			k, prnt = prnt, prnt>>1
		} else {
			break
		}
	}
}

//shiftDown is used to float the element down to maintain the correntness of the heap.
func (pq *priorityQueue) shiftDown(k int) {
	u := k
	for u <= pq.size {
		lc, rc := u<<1, (u<<1)+1
		if lc <= pq.size && pq.comparator(pq.elems[lc], pq.elems[u])*int(pq.order) < 0 {
			u = lc
		}
		if rc <= pq.size && pq.comparator(pq.elems[rc], pq.elems[u])*int(pq.order) < 0 {
			u = rc
		}
		if u == k {
			break
		}
		// swap.
		pq.elems[u], pq.elems[k] = pq.elems[k], pq.elems[u]
		k = u
	}
}

// expand is used to allocate more space to storage more elements
func (pq *priorityQueue) expand() {
	newcap := pq.capacity + (pq.capacity >> 1)
	if newcap > MaxCap {
		newcap = MaxCap
	}
	pq.capacity = newcap
	// copy elements to new space.
	newarr := make([]interface{}, pq.capacity)
	for i := 1; i <= pq.size; i++ {
		newarr[i] = pq.elems[i]
	}
	pq.elems = newarr
}
