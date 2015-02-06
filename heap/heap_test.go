package heap

import (
	"math/rand"
	"testing"
)

type Int int

func (i Int) Compare(other KType) int { return int(i - other.(Int)) }

// Adapted from `container/heap`.

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

func verify(t *testing.T, h *Heap, i Int) {
	n := Int(h.Len())
	j1 := 2*i + 1
	j2 := 2*i + 2
	if j1 < n {
		if j1.Compare(i) < 0 {
			t.Errorf("heap invariant invalidated [%v] = %v > [%v] = %v", i, h.pq[1+i], j1, h.pq[1+j1])
			return
		}
		verify(t, h, j1)
	}
	if j2 < n {
		if j2.Compare(i) < 1 {
			t.Errorf("heap invariant invalidated [%v] = %v > [%v] = %v", i, h.pq[1+i], j1, h.pq[1+j2])
			return
		}
		verify(t, h, j2)
	}
}

func TestHeap0(t *testing.T) {
	h := NewHeap()
	if h.Len() != 0 {
		t.Errorf("want Len=%d, was %d", 0, h.Len())
	}

	for i := 0; i < 20; i++ {
		if h.Len() != i {
			t.Errorf("want Len=%d, was %d", i, h.Len())
		}
		h.Push(Int(i)) // all elements are the same
	}
	verify(t, h, 0)

	for i := h.Len(); i > 0; i-- {
		if h.Len() != i {
			t.Errorf("want Len=%d, was %d", i, h.Len())
		}
		x := h.Pop()
		verify(t, h, 0)
		if x != Int(i-1) {
			t.Errorf("%v.th pop got %v; want %v", i, x, Int(i-1))
		}
	}
}

func TestHeap1(t *testing.T) {
	h := NewHeap()
	for i := 20; i > 0; i-- {
		h.Push(Int(i)) // all elements are different
	}
	verify(t, h, 0)

	for i := 20; h.Len() > 0; i-- {
		x := h.Pop()
		verify(t, h, 0)
		if x != Int(i) {
			t.Errorf("%v.th pop got %v; want %v", i, x, i)
		}
	}
}

func TestHeap(t *testing.T) {
	h := NewHeap()
	verify(t, h, 0)

	for i := 20; i > 10; i-- {
		h.Push(Int(i))
		if want, got := Int(20), h.Peek().(Int); got != want {
			t.Errorf("peek: want %v, got %v", want, got)
		}
	}
	verify(t, h, 0)

	for i := 10; i > 0; i-- {
		h.Push(Int(i))
		verify(t, h, 0)
	}

	for i := h.Len(); h.Len() > 0; i-- {
		x := h.Pop()
		if i > 20 {
			h.Push(Int(20 + i))
		}
		verify(t, h, 0)
		if x != Int(i) {
			t.Errorf("%v.th pop got %v; want %v", i, x, i)
		}
	}
}

func TestHeapRemoveTop(t *testing.T) {
	h := NewHeap()
	for i := 0; i < 10; i++ {
		h.Push(Int(i))
	}
	verify(t, h, 0)

	for h.Len() > 0 {
		i := h.Len() - 1
		if !h.Remove(Int(i)) {
			t.Errorf("should have removed %d", i)
		}
		verify(t, h, 0)
	}
}

func TestHeapRemove(t *testing.T) {
	h := NewHeap()
	for i := 0; i < 40; i++ {
		h.Push(Int(i))
	}
	verify(t, h, 0)

	for i := 10; i < 20; i++ {
		if !h.Remove(Int(i)) {
			t.Errorf("should have removed %d", i)
		}
		verify(t, h, 0)
	}
}

func TestHeapRemoveNotThere(t *testing.T) {
	h := NewHeap()
	for i := 10; i < 20; i++ {
		h.Push(Int(i))
	}
	verify(t, h, 0)

	for i := 0; i < 10; i++ {
		removed := h.Remove(Int(i))
		if removed {
			t.Errorf("should not have removed %d", i)
		}
		verify(t, h, 0)
	}
	for i := 20; i < 40; i++ {
		removed := h.Remove(Int(i))
		if removed {
			t.Errorf("should not have removed %d", i)
		}
		verify(t, h, 0)
	}
	rand.Perm(10)
	for _, n := range rand.Perm(10) {
		i := n + 10
		removed := h.Remove(Int(i))
		if !removed {
			t.Errorf("should have removed %d", i)
		}
		verify(t, h, 0)
	}
}

func BenchmarkDup(b *testing.B) {
	const n = 10000
	buf := make([]KType, 0, n)
	h := NewHeap(buf...)
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			h.Push(Int(0)) // all elements are the same
		}
		for h.Len() > 0 {
			h.Pop()
		}
	}
}

func TestHeapFix(t *testing.T) {
	h := NewHeap()
	verify(t, h, 0)

	for i := 200; i > 0; i -= 10 {
		h.Push(Int(i))
	}
	verify(t, h, 0)

	if h.pq[1] != Int(200) {
		t.Fatalf("Expected head to be 10, was %v", h.pq[1])
	}
	h.pq[1] = Int(0)
	h.Fix()
	verify(t, h, 0)

	for i := 100; i > 0; i-- {
		elem := rand.Intn(h.Len())
		if i&1 == 0 {
			v := h.pq[1+elem].(Int)
			h.pq[1+elem] = v * Int(2)
		} else {
			v := h.pq[1+elem].(Int)
			h.pq[1+elem] = v / Int(2)
		}
		h.Fix()
		verify(t, h, 0)
	}
}
