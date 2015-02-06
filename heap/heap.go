package heap

// Most of the implementation is adapted from Algorithms 4ed by Sedgewick
// and Wayne.

// Comments are adapted from `container/heap`.
// 	 Copyright 2009 The Go Authors. All rights reserved.
// 	 Use of this source code is governed by a BSD-style
// 	 license that can be found in the LICENSE file.

// Heap is a container of KType, where the elements can be efficiently
// retrieved in their decreasing order (according to their comparison
// rules).
type Heap struct {
	n  int
	pq []KType
}

// NewHeap creates a heap, optionaly with keys already populating
// it. The complexity is O(n) where n = len(keys).
func NewHeap(keys ...KType) *Heap {
	h := &Heap{
		n:  len(keys),
		pq: append(make([]KType, 1), keys...),
	}
	h.Fix()
	return h
}

func (h Heap) compare(a, b KType) int { return a.Compare(b) }

// Len is the number of elements stored in the heap.
func (h *Heap) Len() int { return h.n }

// Peek at the largest element (according to their comparison rules), without
// removing it from the heap.
func (h *Heap) Peek() KType { return h.pq[1] }

// Fix re-establishes the heap ordering. This is useful if elements
// of the heap have had their comparison value changed. It is equivalent to,
// but less expenasive than, Pop'ing all the elements and Push'ing them
// again.
// The complexity is O(n).
func (h *Heap) Fix() {
	for i := (len(h.pq) - 1) / 2; i > 0; i-- {
		h.sink(i, len(h.pq)-1)
	}
}

// Push pushes the element k onto the heap. The complexity is
// O(log(n)) where n == h.Len().
func (h *Heap) Push(k KType) {
	h.n++
	h.pq = append(h.pq, k)
	h.swim(h.n)
}

// Pop removes the largest element (according to their comparison rules) from
// the heap and returns it. The complexity is O(log(n)) where n == h.Len().
func (h *Heap) Pop() KType {
	h.n--
	val := h.pq[1]
	h.swap(1, len(h.pq)-1)
	h.pq = h.pq[:len(h.pq)-1]
	h.sink(1, len(h.pq)-1)

	return val
}

// Remove removes k from the heap, if it exists. Equality is defined by
// Compare == 0.
// The complexity is O(n+log(n)) where n == h.Len().
func (h *Heap) Remove(k KType) bool {

	cmp := h.compare(h.pq[1], k)
	if cmp == 0 {
		_ = h.Pop()
		return true
	}
	if cmp < 0 {
		// larger than largest, don't try to find it
		return false
	}

	i := 0
	for _, j := range h.pq[1:] {
		i++
		if h.compare(j, k) != 0 {
			continue
		}
		h.swap(i, 1)
		h.sink(i, len(h.pq)-1)
		_ = h.Pop()
		return true
	}
	// not in the heap
	return false
}

func (h *Heap) swap(i, j int)      { h.pq[i], h.pq[j] = h.pq[j], h.pq[i] }
func (h *Heap) less(i, j int) bool { return h.compare(h.pq[i], h.pq[j]) < 0 }

func (h *Heap) swim(k int) {
	for k > 1 && h.less(k/2, k) {
		h.swap(k/2, k)
		k = k / 2
	}
}

func (h *Heap) sink(k, n int) {

	for k*2 <= n {
		j := 2 * k
		if j < n && h.less(j, j+1) {
			j++
		}
		if !h.less(k, j) {
			break
		}
		h.swap(k, j)
		k = j
	}
}
