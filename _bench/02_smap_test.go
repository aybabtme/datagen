// +build own

package bench

import (
	"testing"

	. "github.com/aybabtme/datagen/codegen"
)

// SortedBytesToStringMap

func Benchmark_SortedMap_BytesToString_Insert(b *testing.B) {
	tree := NewSortedBytesToStringMap()
	bytes := makeBytes(b.N)
	strs := makeStrings(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Put(bytes[b.N-i-1], strs[i])
	}
}

func Benchmark_SortedMap_BytesToString_Delete(b *testing.B) {
	tree := NewSortedBytesToStringMap()
	bytes := makeBytes(b.N)
	strs := makeStrings(b.N)

	for i := 0; i < b.N; i++ {
		tree.Put(bytes[b.N-i-1], strs[i])
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Delete(bytes[i])
	}
}

func Benchmark_SortedMap_BytesToString_DeleteMin(b *testing.B) {
	tree := NewSortedBytesToStringMap()
	bytes := makeBytes(b.N)
	strs := makeStrings(b.N)

	for i := 0; i < b.N; i++ {
		tree.Put(bytes[b.N-i-1], strs[i])
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.DeleteMin()
	}
}

// SortedFloat64ToStringMap

func Benchmark_SortedMap_Float64ToString_Insert(b *testing.B) {
	tree := NewSortedFloat64ToStringMap()
	strs := makeStrings(b.N)

	b.ResetTimer()
	for i := float64(b.N - 1); i > 0; i-- {
		tree.Put(i, strs[int(i)])
	}
}

func Benchmark_SortedMap_Float64ToString_Delete(b *testing.B) {
	tree := NewSortedFloat64ToStringMap()
	strs := makeStrings(b.N)

	for i := float64(b.N - 1); i > 0; i-- {
		tree.Put(i, strs[int(i)])
	}

	b.ResetTimer()
	for i := float64(b.N); i > 0; i-- {
		tree.Delete(i)
	}
}

func Benchmark_SortedMap_Float64ToString_DeleteMin(b *testing.B) {
	tree := NewSortedFloat64ToStringMap()
	strs := makeStrings(b.N)

	for i := float64(b.N - 1); i > 0; i-- {
		tree.Put(i, strs[int(i)])
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.DeleteMin()
	}
}

// SortedIntToStringMap

func Benchmark_SortedMap_IntToString_Insert(b *testing.B) {
	tree := NewSortedIntToStringMap()
	strs := makeStrings(b.N)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Put(b.N-i-1, strs[i])
	}
}

func Benchmark_SortedMap_IntToString_Delete(b *testing.B) {
	tree := NewSortedIntToStringMap()
	strs := makeStrings(b.N)

	for i := 0; i < b.N; i++ {
		tree.Put(b.N-i-1, strs[i])
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Delete(i)
	}
}

func Benchmark_SortedMap_IntToString_DeleteMin(b *testing.B) {
	tree := NewSortedIntToStringMap()
	strs := makeStrings(b.N)

	for i := 0; i < b.N; i++ {
		tree.Put(b.N-i-1, strs[i])
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.DeleteMin()
	}
}

// SortedStringToStringMap

func Benchmark_SortedMap_StringToString_Insert(b *testing.B) {
	tree := NewSortedStringToStringMap()
	strs := makeStrings(b.N)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Put(strs[b.N-i-1], strs[i])
	}
}

func Benchmark_SortedMap_StringToString_Delete(b *testing.B) {
	tree := NewSortedStringToStringMap()
	strs := makeStrings(b.N)

	for i := 0; i < b.N; i++ {
		tree.Put(strs[b.N-i-1], strs[i])
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Delete(strs[i])
	}
}

func Benchmark_SortedMap_StringToString_DeleteMin(b *testing.B) {
	tree := NewSortedStringToStringMap()
	strs := makeStrings(b.N)

	for i := 0; i < b.N; i++ {
		tree.Put(strs[b.N-i-1], strs[i])
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.DeleteMin()
	}
}
