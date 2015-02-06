// +build own

package bench

import (
	"testing"

	. "github.com/aybabtme/datagen/codegen"
)

func Benchmark_Heap_StringHeap(b *testing.B) {
	const n = 10000
	buf := make([]string, 0, n)
	h := NewStringHeap(buf...)
	vals := makeStrings(n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			h.Push(vals[j])
		}
		for h.Len() > 0 {
			h.Pop()
		}
	}
}

func Benchmark_Heap_IntHeap(b *testing.B) {
	const n = 10000
	buf := make([]int, 0, n)
	h := NewIntHeap(buf...)
	vals := makeInts(n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			h.Push(vals[j])
		}
		for h.Len() > 0 {
			h.Pop()
		}
	}
}

func Benchmark_Heap_Float64Heap(b *testing.B) {
	const n = 10000
	buf := make([]float64, 0, n)
	h := NewFloat64Heap(buf...)
	vals := makeFloat64s(n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			h.Push(vals[j])
		}
		for h.Len() > 0 {
			h.Pop()
		}
	}
}

func Benchmark_Heap_BytesHeap(b *testing.B) {
	const n = 10000
	buf := make([][]byte, 0, n)
	h := NewBytesHeap(buf...)
	vals := makeBytes(n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			h.Push(vals[j])
		}
		for h.Len() > 0 {
			h.Pop()
		}
	}
}

func Benchmark_Heap_StringHeap_Push(b *testing.B) {
	h := NewStringHeap()
	vals := makeStrings(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Push(vals[i])
	}
}

func Benchmark_Heap_IntHeap_Push(b *testing.B) {
	h := NewIntHeap()
	vals := makeInts(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Push(vals[i])
	}
}

func Benchmark_Heap_Float64Heap_Push(b *testing.B) {
	h := NewFloat64Heap()
	vals := makeFloat64s(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Push(vals[i])
	}
}

func Benchmark_Heap_BytesHeap_Push(b *testing.B) {
	h := NewBytesHeap()
	vals := makeBytes(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Push(vals[i])
	}
}

func Benchmark_Heap_StringHeap_Pop(b *testing.B) {
	h := NewStringHeap()
	vals := makeStrings(b.N)
	for i := 0; i < b.N; i++ {
		h.Push(vals[i])
	}
	b.ResetTimer()
	for h.Len() != 0 {
		_ = h.Pop()
	}
}

func Benchmark_Heap_IntHeap_Pop(b *testing.B) {
	h := NewIntHeap()
	vals := makeInts(b.N)
	for i := 0; i < b.N; i++ {
		h.Push(vals[i])
	}
	b.ResetTimer()
	for h.Len() != 0 {
		_ = h.Pop()
	}
}

func Benchmark_Heap_Float64Heap_Pop(b *testing.B) {
	h := NewFloat64Heap()
	vals := makeFloat64s(b.N)
	for i := 0; i < b.N; i++ {
		h.Push(vals[i])
	}
	b.ResetTimer()
	for h.Len() != 0 {
		_ = h.Pop()
	}
}

func Benchmark_Heap_BytesHeap_Pop(b *testing.B) {
	h := NewBytesHeap()
	vals := makeBytes(b.N)
	for i := 0; i < b.N; i++ {
		h.Push(vals[i])
	}
	b.ResetTimer()
	for h.Len() != 0 {
		_ = h.Pop()
	}
}
