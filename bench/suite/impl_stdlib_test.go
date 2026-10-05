package suite

import (
	"container/heap"
	"container/list"
)

// stdlib: container/heap and container/list, which box every element in an
// interface.

type stdlibHeapSlice[K key] []K

func (h stdlibHeapSlice[K]) Len() int           { return len(h) }
func (h stdlibHeapSlice[K]) Less(i, j int) bool { return h[i] < h[j] }
func (h stdlibHeapSlice[K]) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *stdlibHeapSlice[K]) Push(x any)        { *h = append(*h, x.(K)) }
func (h *stdlibHeapSlice[K]) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

type stdlibHeap[K key] struct{ h *stdlibHeapSlice[K] }

func newStdlibHeap[K key]() stdlibHeap[K] { return stdlibHeap[K]{new(stdlibHeapSlice[K])} }

func (a stdlibHeap[K]) Push(k K) { heap.Push(a.h, k) }
func (a stdlibHeap[K]) Pop() K   { return heap.Pop(a.h).(K) }
func (a stdlibHeap[K]) Len() int { return a.h.Len() }

type stdlibQueue[K key] struct{ l *list.List }

func newStdlibQueue[K key]() stdlibQueue[K] { return stdlibQueue[K]{list.New()} }

func (a stdlibQueue[K]) Push(k K) { a.l.PushBack(k) }
func (a stdlibQueue[K]) Pop() K   { return a.l.Remove(a.l.Front()).(K) }
