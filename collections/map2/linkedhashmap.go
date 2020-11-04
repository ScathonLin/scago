package map2

const (
	// default cap of map.
	defaultCapacity = 16
	// upper limit cap of map.
	maxCapacity = (1 << 31) - 1
)

// node is actual item storaged in the map.
// it includes key, value and 2 pointer used to keep the order of insertion operation.
type node struct {
	key   interface{}
	value interface{}
	prev  *node
	next  *node
}

// linkedhashmap defines the global structure of linkedhashmap.
// It has a map container implemented by golang standard library
// and head & tail pointer used to keep the order of insertion
type linkedhashmap struct {
	container map[interface{}]*node
	head      *node
	tail      *node
}

//NewLinkedHashMap create a default linkedhashmap with default capacity.
func NewLinkedHashMap() *linkedhashmap {
	head, tail := &node{}, &node{}
	head.next = tail
	tail.prev = head
	return &linkedhashmap{container: make(map[interface{}]*node, defaultCapacity), head: head, tail: tail}
}

//NewLinekedHashMapWithCap create a linkedhashmap with given capatity which will be
//caculated a new value that is power of 2 which is closest to it among all nums greater than it.
func NewLinekedHashMapWithCap(capacity int32) *linkedhashmap {
	head, tail := &node{}, &node{}
	head.next = tail
	tail.prev = head
	return &linkedhashmap{container: make(map[interface{}]*node, tableSizeFor(capacity)), head: head, tail: tail}
}

//Remove node by given key in the map.
func (lhm *linkedhashmap) Remove(key interface{}) interface{} {
	if vnode, ok := lhm.container[key]; ok {
		vnode.prev.next = vnode.next
		vnode.next.prev = vnode.prev
		delete(lhm.container, key)
		return vnode.value
	}
	return nil
}

//Contains is used to check whether existing the value matched the given key.
func (lhm *linkedhashmap) Contains(key interface{}) bool {
	_, ok := lhm.container[key]
	return ok
}

//Values returned all the values in the map in insertion order.
func (lhm *linkedhashmap) Values() []interface{} {
	values := make([]interface{}, len(lhm.container))
	i := 0
	h := *lhm.head
	t := *lhm.tail
	for ; *(h.next) != t; i++ {
		values[i] = h.next.value
		h.next = h.next.next
	}
	return values
}

//Keys returned all the keys in the map in insertion order.
func (lhm *linkedhashmap) Keys() []interface{} {
	keys := make([]interface{}, len(lhm.container))
	i := 0
	// we must use value type here for head and tail pointer, because head and tail are pointer type originally,
	// if we used them directly, we will indeed change the prev and next pointer between two nodes while
	// we iterate the doubly linked list which is not correct.
	h := *lhm.head
	t := *lhm.tail
	for ; *(h.next) != t; i++ {
		keys[i] = h.next.key
		h.next = h.next.next
	}
	return keys
}

//Put a pair of key and value into the map.
func (lhm *linkedhashmap) Put(key, value interface{}) {
	if vnode, ok := lhm.container[key]; ok {
		vnode.value = value
		return
	}
	newNode := &node{key: key, value: value}
	lhm.container[key] = newNode
	// link the node to the tail of double direct linkedlist.
	lhm.tail.prev.next = newNode
	newNode.prev = lhm.tail.prev
	newNode.next = lhm.tail
	lhm.tail.prev = newNode
}

//Get returns a value which matches the given key.
func (lhm *linkedhashmap) Get(key interface{}) interface{} {
	if node, ok := lhm.container[key]; ok {
		return node.value
	}
	return nil
}

//Size returns num of kvs int map.
func (lhm *linkedhashmap) Size() int32 {
	return int32(len(lhm.container))
}

//tableSizeFor recalculate a new cap which is closest to the given cap
//among all the nums geater than given cap.
func tableSizeFor(capacity int32) int32 {
	capacity--
	capacity |= capacity >> 1
	capacity |= capacity >> 2
	capacity |= capacity >> 4
	capacity |= capacity >> 8
	capacity |= capacity >> 16
	if capacity < 0 {
		return 1
	} else if capacity+1 > maxCapacity {
		return maxCapacity
	} else {
		return capacity + 1
	}

}
