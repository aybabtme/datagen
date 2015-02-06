// +build other

package bench

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/petar/GoLLRB/llrb"
)

// GoLLRB

func Benchmark_SortedMap_BytesToString_Insert(b *testing.B) {
	goLLRB_Insert(b, makeLlrbBytes(b.N))
}
func Benchmark_SortedMap_BytesToString_Delete(b *testing.B) {
	goLLRB_Delete(b, makeLlrbBytes(b.N))
}
func Benchmark_SortedMap_BytesToString_DeleteMin(b *testing.B) {
	goLLRB_DeleteMin(b, makeLlrbBytes(b.N))
}

func Benchmark_SortedMap_Float64ToString_Insert(b *testing.B) {
	goLLRB_Insert(b, makeLlrbFloat64s(b.N))
}
func Benchmark_SortedMap_Float64ToString_Delete(b *testing.B) {
	goLLRB_Delete(b, makeLlrbFloat64s(b.N))
}
func Benchmark_SortedMap_Float64ToString_DeleteMin(b *testing.B) {
	goLLRB_DeleteMin(b, makeLlrbFloat64s(b.N))
}

func Benchmark_SortedMap_IntToString_Insert(b *testing.B) {
	goLLRB_Insert(b, makeLlrbInts(b.N))
}
func Benchmark_SortedMap_IntToString_Delete(b *testing.B) {
	goLLRB_Delete(b, makeLlrbInts(b.N))
}
func Benchmark_SortedMap_IntToString_DeleteMin(b *testing.B) {
	goLLRB_DeleteMin(b, makeLlrbInts(b.N))
}

func Benchmark_SortedMap_StringToString_Insert(b *testing.B) {
	goLLRB_Insert(b, makeLlrbStrings(b.N))
}
func Benchmark_SortedMap_StringToString_Delete(b *testing.B) {
	goLLRB_Delete(b, makeLlrbStrings(b.N))
}
func Benchmark_SortedMap_StringToString_DeleteMin(b *testing.B) {
	goLLRB_DeleteMin(b, makeLlrbStrings(b.N))
}

func goLLRB_Insert(b *testing.B, items []llrb.Item) {
	tree := llrb.New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.ReplaceOrInsert(items[i])
	}
}

func goLLRB_Delete(b *testing.B, items []llrb.Item) {
	tree := llrb.New()
	for i := 0; i < b.N; i++ {
		tree.ReplaceOrInsert(items[i])
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Delete(items[i])
	}
}

func goLLRB_DeleteMin(b *testing.B, items []llrb.Item) {
	tree := llrb.New()
	for i := 0; i < b.N; i++ {
		tree.ReplaceOrInsert(items[i])
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.DeleteMin()
	}
}

// GoLLRB scaffolding

type String struct {
	key string
	val string
}

func (s String) Less(than llrb.Item) bool { return string(s.key) < string(than.(String).key) }

type Int struct {
	key int
	val string
}

func (s Int) Less(than llrb.Item) bool { return int(s.key) < int(than.(Int).key) }

type Float64 struct {
	key float64
	val string
}

func (s Float64) Less(than llrb.Item) bool { return float64(s.key) < float64(than.(Float64).key) }

type Bytes struct {
	key []byte
	val string
}

func (s Bytes) Less(than llrb.Item) bool {
	return bytes.Compare([]byte(s.key), []byte(than.(Bytes).key)) < 0
}

func makeLlrbStrings(n int) []llrb.Item {
	items := make([]llrb.Item, 0, n)
	for i := 0; i < n; i++ {
		items = append(items, String{
			key: strconv.Itoa(i),
			val: strconv.Itoa(i),
		})
	}
	return items
}

func makeLlrbInts(n int) []llrb.Item {
	items := make([]llrb.Item, 0, n)
	for i := 0; i < n; i++ {
		items = append(items, Int{
			key: i,
			val: strconv.Itoa(i),
		})
	}
	return items
}

func makeLlrbFloat64s(n int) []llrb.Item {
	items := make([]llrb.Item, 0, n)
	for i := 0; i < n; i++ {
		items = append(items, Float64{
			key: float64(i),
			val: strconv.Itoa(i),
		})
	}
	return items
}

func makeLlrbBytes(n int) []llrb.Item {
	items := make([]llrb.Item, 0, n)
	for i := 0; i < n; i++ {
		items = append(items, Bytes{
			key: []byte(strconv.Itoa(i)),
			val: strconv.Itoa(i),
		})
	}
	return items
}
