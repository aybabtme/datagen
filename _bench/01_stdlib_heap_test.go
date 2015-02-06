// +build other

package bench

import (
	"bytes"
	"container/heap"
	"strconv"
	"testing"
)

// container/heap

func Benchmark_Heap_StringHeap(b *testing.B) {
	const n = 10000
	h := makeStdLibStringHeap(n)
	vals := makeStrings(n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			heap.Push(&h, vals[j])
		}
		for h.Len() > 0 {
			heap.Pop(&h)
		}
	}
}

func Benchmark_Heap_IntHeap(b *testing.B) {
	const n = 10000
	h := makeStdLibIntHeap(n)
	vals := makeInts(n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			heap.Push(&h, vals[j])
		}
		for h.Len() > 0 {
			heap.Pop(&h)
		}
	}
}

func Benchmark_Heap_Float64Heap(b *testing.B) {
	const n = 10000
	h := makeStdLibFloat64Heap(n)
	vals := makeFloat64s(n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			heap.Push(&h, vals[j])
		}
		for h.Len() > 0 {
			heap.Pop(&h)
		}
	}
}

func Benchmark_Heap_BytesHeap(b *testing.B) {
	const n = 10000
	h := makeStdLibBytesHeap(n)
	vals := makeBytes(n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			heap.Push(&h, vals[j])
		}
		for h.Len() > 0 {
			heap.Pop(&h)
		}
	}
}

func Benchmark_Heap_StringHeap_Push(b *testing.B) {
	h := makeStdLibStringHeap(0)
	vals := makeStrings(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		heap.Push(&h, vals[i])
	}
}

func Benchmark_Heap_IntHeap_Push(b *testing.B) {
	h := makeStdLibIntHeap(0)
	vals := makeInts(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		heap.Push(&h, vals[i])
	}
}

func Benchmark_Heap_Float64Heap_Push(b *testing.B) {
	h := makeStdLibFloat64Heap(0)
	vals := makeFloat64s(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		heap.Push(&h, vals[i])
	}
}

func Benchmark_Heap_BytesHeap_Push(b *testing.B) {
	h := makeStdLibBytesHeap(0)
	vals := makeBytes(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		heap.Push(&h, vals[i])
	}
}

func Benchmark_Heap_StringHeap_Pop(b *testing.B) {
	h := makeStdLibStringHeap(0)
	vals := makeStrings(b.N)
	for i := 0; i < b.N; i++ {
		heap.Push(&h, vals[i])
	}
	b.ResetTimer()
	for h.Len() != 0 {
		_ = heap.Pop(&h)
	}
}

func Benchmark_Heap_IntHeap_Pop(b *testing.B) {
	h := makeStdLibIntHeap(0)
	vals := makeInts(b.N)
	for i := 0; i < b.N; i++ {
		heap.Push(&h, vals[i])
	}
	b.ResetTimer()
	for h.Len() != 0 {
		_ = heap.Pop(&h)
	}
}

func Benchmark_Heap_Float64Heap_Pop(b *testing.B) {
	h := makeStdLibFloat64Heap(0)
	vals := makeFloat64s(b.N)
	for i := 0; i < b.N; i++ {
		heap.Push(&h, vals[i])
	}
	b.ResetTimer()
	for h.Len() != 0 {
		_ = heap.Pop(&h)
	}
}

func Benchmark_Heap_BytesHeap_Pop(b *testing.B) {
	h := makeStdLibBytesHeap(0)
	vals := makeBytes(b.N)
	for i := 0; i < b.N; i++ {
		heap.Push(&h, vals[i])
	}
	b.ResetTimer()
	for h.Len() != 0 {
		_ = heap.Pop(&h)
	}
}

// container/heap scaffolding

type StdLibStringHeap []string

type StdLibIntHeap []int

type StdLibFloat64Heap []float64

type StdLibBytesHeap [][]byte

func makeStdLibStringHeap(n int) StdLibStringHeap {
	items := make(StdLibStringHeap, 0, n)
	for i := 0; i < n; i++ {
		items = append(items, strconv.Itoa(i))
	}
	return items
}

func makeStdLibIntHeap(n int) StdLibIntHeap {
	items := make(StdLibIntHeap, 0, n)
	for i := 0; i < n; i++ {
		items = append(items, i)
	}
	return items
}

func makeStdLibFloat64Heap(n int) StdLibFloat64Heap {
	items := make(StdLibFloat64Heap, 0, n)
	for i := 0; i < n; i++ {
		items = append(items, float64(i))
	}
	return items
}

func makeStdLibBytesHeap(n int) StdLibBytesHeap {
	items := make(StdLibBytesHeap, 0, n)
	for i := 0; i < n; i++ {
		items = append(items, []byte(strconv.Itoa(i)))
	}
	return items
}

func (h StdLibStringHeap) Len() int            { return len(h) }
func (h StdLibStringHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h StdLibStringHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *StdLibStringHeap) Push(x interface{}) { *h = append(*h, x.(string)) }
func (h *StdLibStringHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h StdLibIntHeap) Len() int            { return len(h) }
func (h StdLibIntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h StdLibIntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *StdLibIntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *StdLibIntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h StdLibFloat64Heap) Len() int            { return len(h) }
func (h StdLibFloat64Heap) Less(i, j int) bool  { return h[i] < h[j] }
func (h StdLibFloat64Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *StdLibFloat64Heap) Push(x interface{}) { *h = append(*h, x.(float64)) }
func (h *StdLibFloat64Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h StdLibBytesHeap) Len() int { return len(h) }
func (h StdLibBytesHeap) Less(i, j int) bool {
	left := []byte(h[i])
	right := []byte(h[j])
	return bytes.Compare(left, right) < 0
}
func (h StdLibBytesHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *StdLibBytesHeap) Push(x interface{}) { *h = append(*h, x.([]byte)) }
func (h *StdLibBytesHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
