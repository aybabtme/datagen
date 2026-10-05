package suite

import "github.com/aybabtme/datagen/codegen"

// own: datagen's code generated implementations. Heaps and queues already
// have the benchmark's method set.

func newOwnIntHeap() *codegen.IntHeap       { return codegen.NewIntHeap() }
func newOwnStringHeap() *codegen.StringHeap { return codegen.NewStringHeap() }

func newOwnIntQueue() *codegen.IntQueue       { return codegen.NewIntQueue(0) }
func newOwnStringQueue() *codegen.StringQueue { return codegen.NewStringQueue(0) }

type ownIntMap struct{ m *codegen.SortedIntToStringMap }

func newOwnIntMap() ownIntMap { return ownIntMap{codegen.NewSortedIntToStringMap()} }

func (a ownIntMap) Put(k int, v string)      { a.m.Put(k, v) }
func (a ownIntMap) Get(k int) (string, bool) { return a.m.Get(k) }
func (a ownIntMap) Delete(k int)             { a.m.Delete(k) }
func (a ownIntMap) DeleteMin() int           { k, _, _ := a.m.DeleteMin(); return k }

type ownStringMap struct {
	m *codegen.SortedStringToStringMap
}

func newOwnStringMap() ownStringMap { return ownStringMap{codegen.NewSortedStringToStringMap()} }

func (a ownStringMap) Put(k string, v string)      { a.m.Put(k, v) }
func (a ownStringMap) Get(k string) (string, bool) { return a.m.Get(k) }
func (a ownStringMap) Delete(k string)             { a.m.Delete(k) }
func (a ownStringMap) DeleteMin() string           { k, _, _ := a.m.DeleteMin(); return k }

type ownIntSet struct{ s *codegen.SortedIntSet }

func newOwnIntSet() ownIntSet { return ownIntSet{codegen.NewSortedIntSet()} }

func (a ownIntSet) Put(k int)      { a.s.Put(k) }
func (a ownIntSet) Has(k int) bool { return a.s.Contains(k) }
func (a ownIntSet) Delete(k int)   { a.s.Delete(k) }
func (a ownIntSet) DeleteMin() int { k, _ := a.s.DeleteMin(); return k }

type ownStringSet struct{ s *codegen.SortedStringSet }

func newOwnStringSet() ownStringSet { return ownStringSet{codegen.NewSortedStringSet()} }

func (a ownStringSet) Put(k string)      { a.s.Put(k) }
func (a ownStringSet) Has(k string) bool { return a.s.Contains(k) }
func (a ownStringSet) Delete(k string)   { a.s.Delete(k) }
func (a ownStringSet) DeleteMin() string { k, _ := a.s.DeleteMin(); return k }
