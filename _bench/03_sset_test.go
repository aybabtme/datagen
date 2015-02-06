// +build own

package bench

import (
	"testing"

	. "github.com/aybabtme/datagen/codegen"
)

// SortedBytesSet

func Benchmark_SortedSet_Bytes_Insert(b *testing.B) {
	tree := NewSortedBytesSet()
	bytes := makeBytes(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Put(bytes[b.N-i-1])
	}
}

func Benchmark_SortedSet_Bytes_Delete(b *testing.B) {
	tree := NewSortedBytesSet()
	bytes := makeBytes(b.N)

	for i := 0; i < b.N; i++ {
		tree.Put(bytes[b.N-i-1])
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Delete(bytes[i])
	}
}

func Benchmark_SortedSet_Bytes_DeleteMin(b *testing.B) {
	tree := NewSortedBytesSet()
	bytes := makeBytes(b.N)

	for i := 0; i < b.N; i++ {
		tree.Put(bytes[b.N-i-1])
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.DeleteMin()
	}
}

// SortedFloat64Set

func Benchmark_SortedSet_Float64_Insert(b *testing.B) {
	tree := NewSortedFloat64Set()

	b.ResetTimer()
	for i := float64(b.N - 1); i > 0; i-- {
		tree.Put(i)
	}
}

func Benchmark_SortedSet_Float64_Delete(b *testing.B) {
	tree := NewSortedFloat64Set()

	for i := float64(b.N - 1); i > 0; i-- {
		tree.Put(i)
	}

	b.ResetTimer()
	for i := float64(b.N); i > 0; i-- {
		tree.Delete(i)
	}
}

func Benchmark_SortedSet_Float64_DeleteMin(b *testing.B) {
	tree := NewSortedFloat64Set()

	for i := float64(b.N - 1); i > 0; i-- {
		tree.Put(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.DeleteMin()
	}
}

// SortedIntSet

func Benchmark_SortedSet_Int_Insert(b *testing.B) {
	tree := NewSortedIntSet()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Put(b.N - i - 1)
	}
}

func Benchmark_SortedSet_Int_Delete(b *testing.B) {
	tree := NewSortedIntSet()

	for i := 0; i < b.N; i++ {
		tree.Put(b.N - i - 1)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Delete(i)
	}
}

func Benchmark_SortedSet_Int_DeleteMin(b *testing.B) {
	tree := NewSortedIntSet()

	for i := 0; i < b.N; i++ {
		tree.Put(b.N - i - 1)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.DeleteMin()
	}
}

// SortedStringSet

func Benchmark_SortedSet_String_Insert(b *testing.B) {
	tree := NewSortedStringSet()
	strs := makeStrings(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Put(strs[b.N-i-1])
	}
}

func Benchmark_SortedSet_String_Delete(b *testing.B) {
	tree := NewSortedStringSet()
	strs := makeStrings(b.N)
	for i := 0; i < b.N; i++ {
		tree.Put(strs[b.N-i-1])
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Delete(strs[i])
	}
}

func Benchmark_SortedSet_String_DeleteMin(b *testing.B) {
	tree := NewSortedStringSet()
	strs := makeStrings(b.N)
	for i := 0; i < b.N; i++ {
		tree.Put(strs[b.N-i-1])
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.DeleteMin()
	}
}
