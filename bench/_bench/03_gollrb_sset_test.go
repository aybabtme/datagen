// +build other

package bench

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/petar/GoLLRB/llrb"
)

// GoLLRB

func Benchmark_SortedSet_Bytes_Insert(b *testing.B) {
	goLLRB_Insert(b, makeLlrbBytesSet(b.N))
}
func Benchmark_SortedSet_Bytes_Delete(b *testing.B) {
	goLLRB_Delete(b, makeLlrbBytesSet(b.N))
}
func Benchmark_SortedSet_Bytes_DeleteMin(b *testing.B) {
	goLLRB_DeleteMin(b, makeLlrbBytesSet(b.N))
}

func Benchmark_SortedSet_Float64_Insert(b *testing.B) {
	goLLRB_Insert(b, makeLlrbFloat64sSet(b.N))
}
func Benchmark_SortedSet_Float64_Delete(b *testing.B) {
	goLLRB_Delete(b, makeLlrbFloat64sSet(b.N))
}
func Benchmark_SortedSet_Float64_DeleteMin(b *testing.B) {
	goLLRB_DeleteMin(b, makeLlrbFloat64sSet(b.N))
}

func Benchmark_SortedSet_Int_Insert(b *testing.B) {
	goLLRB_Insert(b, makeLlrbIntsSet(b.N))
}
func Benchmark_SortedSet_Int_Delete(b *testing.B) {
	goLLRB_Delete(b, makeLlrbIntsSet(b.N))
}
func Benchmark_SortedSet_Int_DeleteMin(b *testing.B) {
	goLLRB_DeleteMin(b, makeLlrbIntsSet(b.N))
}

func Benchmark_SortedSet_String_Insert(b *testing.B) {
	goLLRB_Insert(b, makeLlrbStringsSet(b.N))
}
func Benchmark_SortedSet_String_Delete(b *testing.B) {
	goLLRB_Delete(b, makeLlrbStringsSet(b.N))
}
func Benchmark_SortedSet_String_DeleteMin(b *testing.B) {
	goLLRB_DeleteMin(b, makeLlrbStringsSet(b.N))
}

// GoLLRB scaffolding

type StringSet string

func (s StringSet) Less(than llrb.Item) bool { return string(s) < string(than.(StringSet)) }

type IntSet int

func (s IntSet) Less(than llrb.Item) bool { return int(s) < int(than.(IntSet)) }

type Float64Set float64

func (s Float64Set) Less(than llrb.Item) bool { return float64(s) < float64(than.(Float64Set)) }

type BytesSet []byte

func (s BytesSet) Less(than llrb.Item) bool {
	return bytes.Compare([]byte(s), []byte(than.(BytesSet))) < 0
}

func makeLlrbStringsSet(n int) []llrb.Item {
	items := make([]llrb.Item, 0, n)
	for i := 0; i < n; i++ {
		items = append(items, StringSet(strconv.Itoa(i)))
	}
	return items
}

func makeLlrbIntsSet(n int) []llrb.Item {
	items := make([]llrb.Item, 0, n)
	for i := 0; i < n; i++ {
		items = append(items, IntSet(i))
	}
	return items
}

func makeLlrbFloat64sSet(n int) []llrb.Item {
	items := make([]llrb.Item, 0, n)
	for i := 0; i < n; i++ {
		items = append(items, Float64Set(float64(i)))
	}
	return items
}

func makeLlrbBytesSet(n int) []llrb.Item {
	items := make([]llrb.Item, 0, n)
	for i := 0; i < n; i++ {
		items = append(items, BytesSet([]byte(strconv.Itoa(i))))
	}
	return items
}
