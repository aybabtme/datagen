package suite

import (
	gbtree "github.com/google/btree"
	"github.com/tidwall/btree"
	"github.com/tidwall/btype"
)

// B-tree based sorted maps and sets.

// btype: github.com/tidwall/btype.

type btypeMap[K key] struct{ m *btype.Map[K, string] }

func newBtypeMap[K key]() btypeMap[K] { return btypeMap[K]{btype.NewMap[K, string]()} }

func (a btypeMap[K]) Put(k K, v string)      { a.m.Set(k, v) }
func (a btypeMap[K]) Get(k K) (string, bool) { return a.m.Get(k) }
func (a btypeMap[K]) Delete(k K)             { a.m.Delete(k) }
func (a btypeMap[K]) DeleteMin() K           { k, _, _ := a.m.PopFront(); return k }

type btypeSet[K key] struct{ s *btype.Set[K] }

func newBtypeSet[K key]() btypeSet[K] { return btypeSet[K]{btype.NewSet[K]()} }

func (a btypeSet[K]) Put(k K)      { a.s.Insert(k) }
func (a btypeSet[K]) Has(k K) bool { return a.s.Contains(k) }
func (a btypeSet[K]) Delete(k K)   { a.s.Delete(k) }
func (a btypeSet[K]) DeleteMin() K { k, _ := a.s.PopFront(); return k }

// tidwall: github.com/tidwall/btree, default degree.

type tidwallMap[K key] struct{ m *btree.Map[K, string] }

func newTidwallMap[K key]() tidwallMap[K] { return tidwallMap[K]{btree.NewMap[K, string](0)} }

func (a tidwallMap[K]) Put(k K, v string)      { a.m.Set(k, v) }
func (a tidwallMap[K]) Get(k K) (string, bool) { return a.m.Get(k) }
func (a tidwallMap[K]) Delete(k K)             { a.m.Delete(k) }
func (a tidwallMap[K]) DeleteMin() K           { k, _, _ := a.m.PopMin(); return k }

type tidwallSet[K key] struct{ s *btree.Set[K] }

func newTidwallSet[K key]() tidwallSet[K] { return tidwallSet[K]{new(btree.Set[K])} }

func (a tidwallSet[K]) Put(k K)      { a.s.Insert(k) }
func (a tidwallSet[K]) Has(k K) bool { return a.s.Contains(k) }
func (a tidwallSet[K]) Delete(k K)   { a.s.Delete(k) }
func (a tidwallSet[K]) DeleteMin() K { k, _ := a.s.PopMin(); return k }

// google: github.com/google/btree. It has no map type, so maps store
// key/value items. Degree 32, as in tidwall/btree-benchmark.

const googleDegree = 32

type googleItem[K key] struct {
	k K
	v string
}

type googleMap[K key] struct{ t *gbtree.BTreeG[googleItem[K]] }

func newGoogleMap[K key]() googleMap[K] {
	return googleMap[K]{gbtree.NewG(googleDegree, func(a, b googleItem[K]) bool { return a.k < b.k })}
}

func (a googleMap[K]) Put(k K, v string) { a.t.ReplaceOrInsert(googleItem[K]{k, v}) }
func (a googleMap[K]) Get(k K) (string, bool) {
	it, ok := a.t.Get(googleItem[K]{k: k})
	return it.v, ok
}
func (a googleMap[K]) Delete(k K)   { a.t.Delete(googleItem[K]{k: k}) }
func (a googleMap[K]) DeleteMin() K { it, _ := a.t.DeleteMin(); return it.k }

type googleSet[K key] struct{ t *gbtree.BTreeG[K] }

func newGoogleSet[K key]() googleSet[K] { return googleSet[K]{gbtree.NewOrderedG[K](googleDegree)} }

func (a googleSet[K]) Put(k K)      { a.t.ReplaceOrInsert(k) }
func (a googleSet[K]) Has(k K) bool { return a.t.Has(k) }
func (a googleSet[K]) Delete(k K)   { a.t.Delete(k) }
func (a googleSet[K]) DeleteMin() K { k, _ := a.t.DeleteMin(); return k }
