package suite

import "github.com/petar/GoLLRB/llrb"

// gollrb: github.com/petar/GoLLRB, the interface based LLRB tree the
// original suite compared against.

type llrbItem[K key] struct {
	k K
	v string
}

func (i llrbItem[K]) Less(than llrb.Item) bool { return i.k < than.(llrbItem[K]).k }

type gollrbMap[K key] struct{ t *llrb.LLRB }

func newGollrbMap[K key]() gollrbMap[K] { return gollrbMap[K]{llrb.New()} }

func (a gollrbMap[K]) Put(k K, v string) { a.t.ReplaceOrInsert(llrbItem[K]{k, v}) }
func (a gollrbMap[K]) Get(k K) (string, bool) {
	it := a.t.Get(llrbItem[K]{k: k})
	if it == nil {
		return "", false
	}
	return it.(llrbItem[K]).v, true
}
func (a gollrbMap[K]) Delete(k K)   { a.t.Delete(llrbItem[K]{k: k}) }
func (a gollrbMap[K]) DeleteMin() K { return a.t.DeleteMin().(llrbItem[K]).k }

type gollrbSet[K key] struct{ t *llrb.LLRB }

func newGollrbSet[K key]() gollrbSet[K] { return gollrbSet[K]{llrb.New()} }

func (a gollrbSet[K]) Put(k K)      { a.t.ReplaceOrInsert(llrbItem[K]{k: k}) }
func (a gollrbSet[K]) Has(k K) bool { return a.t.Has(llrbItem[K]{k: k}) }
func (a gollrbSet[K]) Delete(k K)   { a.t.Delete(llrbItem[K]{k: k}) }
func (a gollrbSet[K]) DeleteMin() K { return a.t.DeleteMin().(llrbItem[K]).k }
