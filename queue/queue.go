package queue

// GENERATED CODE!!!

// Implementation adapted from github.com/eapache/queue:
//    The MIT License (MIT)
//    Copyright (c) 2014 Evan Huus

var nilKType KType

// Queue represents a single instance of the queue data structure.
type Queue struct {
	buf               []KType
	head, tail, count int
	minlen            int
}

// NewQueue constructs and returns a new Queue with an initial capacity.
func NewQueue(capacity int) *Queue {
	// min capacity of 16
	if capacity < 16 {
		capacity = 16
	}
	return &Queue{buf: make([]KType, capacity), minlen: capacity}
}

// Len returns the number of elements currently stored in the queue.
func (q *Queue) Len() int {
	return q.count
}

// Push puts an element on the end of the queue.
func (q *Queue) Push(elem KType) {
	if q.count == len(q.buf) {
		q.resize()
	}

	q.buf[q.tail] = elem
	q.tail = (q.tail + 1) % len(q.buf)
	q.count++
}

// Peek returns the element at the head of the queue. This call panics
// if the queue is empty.
func (q *Queue) Peek() KType {
	if q.Len() <= 0 {
		panic("queue: empty queue")
	}
	return q.buf[q.head]
}

// Get returns the element at index i in the queue. If the index is
// invalid, the call will panic.
func (q *Queue) Get(i int) KType {
	if i >= q.Len() || i < 0 {
		panic("queue: index out of range")
	}
	modi := (q.head + i) % len(q.buf)
	return q.buf[modi]
}

// Pop removes the element from the front of the queue.
// This call panics if the queue is empty.
func (q *Queue) Pop() KType {
	if q.Len() <= 0 {
		panic("queue: empty queue")
	}
	v := q.buf[q.head]
	// set to nil to avoid keeping reference to objects
	// that would otherwise be garbage collected
	q.buf[q.head] = nilKType
	q.head = (q.head + 1) % len(q.buf)
	q.count--
	if len(q.buf) > q.minlen && q.count*4 <= len(q.buf) {
		q.resize()
	}
	return v
}

func (q *Queue) resize() {
	newBuf := make([]KType, q.count*2)

	if q.tail > q.head {
		copy(newBuf, q.buf[q.head:q.tail])
	} else {
		copy(newBuf, q.buf[q.head:len(q.buf)])
		copy(newBuf[len(q.buf)-q.head:], q.buf[:q.tail])
	}

	q.head = 0
	q.tail = q.count
	q.buf = newBuf
}
