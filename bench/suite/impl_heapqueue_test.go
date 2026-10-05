package suite

import (
	"cmp"

	"github.com/eapache/queue/v2"
	"github.com/gammazero/deque"
	"github.com/jba/heap"
)

// jba: github.com/jba/heap, the proposed container/heap/v2.

type jbaHeap[K key] struct{ h *heap.Heap[K] }

func newJbaHeap[K key]() jbaHeap[K] { return jbaHeap[K]{heap.New(cmp.Compare[K])} }

func (a jbaHeap[K]) Push(k K) { a.h.Insert(k) }
func (a jbaHeap[K]) Pop() K   { return a.h.TakeMin() }
func (a jbaHeap[K]) Len() int { return a.h.Len() }

// gammazero: github.com/gammazero/deque, used as a FIFO.

type gammazeroQueue[K key] struct{ q *deque.Deque[K] }

func newGammazeroQueue[K key]() gammazeroQueue[K] { return gammazeroQueue[K]{new(deque.Deque[K])} }

func (a gammazeroQueue[K]) Push(k K) { a.q.PushBack(k) }
func (a gammazeroQueue[K]) Pop() K   { return a.q.PopFront() }

// eapache: github.com/eapache/queue/v2, the generic version of the queue
// datagen's queue was adapted from.

type eapacheQueue[K key] struct{ q *queue.Queue[K] }

func newEapacheQueue[K key]() eapacheQueue[K] { return eapacheQueue[K]{queue.New[K]()} }

func (a eapacheQueue[K]) Push(k K) { a.q.Add(k) }
func (a eapacheQueue[K]) Pop() K   { return a.q.Remove() }
