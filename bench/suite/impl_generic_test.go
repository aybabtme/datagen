package suite

import (
	"cmp"

	dgheap "github.com/aybabtme/datagen/generic/heap"
	dgmap "github.com/aybabtme/datagen/generic/map/redblackbst"
	dgqueue "github.com/aybabtme/datagen/generic/queue"
	dgset "github.com/aybabtme/datagen/generic/set/redblackbst"
)

// generic: the generic ports of datagen's datastructures, in generic/.

func newGenericHeap[K key]() *dgheap.Heap[K]    { return dgheap.NewHeap(cmp.Compare[K]) }
func newGenericQueue[K key]() *dgqueue.Queue[K] { return dgqueue.NewQueue[K](0) }

type genericMap[K key] struct{ m *dgmap.RedBlack[K, string] }

func newGenericMap[K key]() genericMap[K] {
	return genericMap[K]{dgmap.NewRedBlack[K, string](cmp.Compare[K])}
}

func (a genericMap[K]) Put(k K, v string)      { a.m.Put(k, v) }
func (a genericMap[K]) Get(k K) (string, bool) { return a.m.Get(k) }
func (a genericMap[K]) Delete(k K)             { a.m.Delete(k) }
func (a genericMap[K]) DeleteMin() K           { k, _, _ := a.m.DeleteMin(); return k }

type genericSet[K key] struct{ s *dgset.RedBlack[K] }

func newGenericSet[K key]() genericSet[K] { return genericSet[K]{dgset.NewRedBlack(cmp.Compare[K])} }

func (a genericSet[K]) Put(k K)      { a.s.Put(k) }
func (a genericSet[K]) Has(k K) bool { return a.s.Contains(k) }
func (a genericSet[K]) Delete(k K)   { a.s.Delete(k) }
func (a genericSet[K]) DeleteMin() K { k, _ := a.s.DeleteMin(); return k }
