package suite

import (
	"slices"
	"testing"
)

// TestImpls checks that every adapter behaves like the datastructure it
// claims to be, so that the benchmarks measure the work they say they do.
func TestImpls(t *testing.T) {
	ints, strs := intKeys(1000), stringKeys(1000)
	checkAll(t, "map", ints, strs, intMaps, stringMaps)
	checkAll(t, "set", ints, strs, intSets, stringSets)
	checkAll(t, "heap", ints, strs, intHeaps, stringHeaps)
	checkAll(t, "queue", ints, strs, intQueues, stringQueues)
}

func checkAll(t *testing.T, kind string, ints []int, strs []string, intImpls []impl[int], strImpls []impl[string]) {
	for _, im := range intImpls {
		t.Run(kind+"/int/"+im.name, func(t *testing.T) { im.check(t, ints) })
	}
	for _, im := range strImpls {
		t.Run(kind+"/string/"+im.name, func(t *testing.T) { im.check(t, strs) })
	}
}

func checkMap[M sortedMap[K], K key](t *testing.T, newMap func() M, keys []K) {
	m := newMap()
	for _, k := range keys {
		m.Put(k, val)
	}
	for _, k := range keys {
		if v, ok := m.Get(k); !ok || v != val {
			t.Fatalf("Get(%v) = %q, %v", k, v, ok)
		}
	}
	gone, kept := keys[:len(keys)/2], keys[len(keys)/2:]
	for _, k := range gone {
		m.Delete(k)
	}
	for _, k := range gone {
		if _, ok := m.Get(k); ok {
			t.Fatalf("Get(%v) found a deleted key", k)
		}
	}
	got := make([]K, 0, len(kept))
	for range kept {
		got = append(got, m.DeleteMin())
	}
	if want := sorted(kept); !slices.Equal(got, want) {
		t.Fatalf("DeleteMin order: got %v..., want %v...", got[:5], want[:5])
	}
}

func checkSet[S sortedSet[K], K key](t *testing.T, newSet func() S, keys []K) {
	s := newSet()
	for _, k := range keys {
		s.Put(k)
	}
	for _, k := range keys {
		if !s.Has(k) {
			t.Fatalf("Has(%v) = false", k)
		}
	}
	gone, kept := keys[:len(keys)/2], keys[len(keys)/2:]
	for _, k := range gone {
		s.Delete(k)
	}
	for _, k := range gone {
		if s.Has(k) {
			t.Fatalf("Has(%v) found a deleted key", k)
		}
	}
	got := make([]K, 0, len(kept))
	for range kept {
		got = append(got, s.DeleteMin())
	}
	if want := sorted(kept); !slices.Equal(got, want) {
		t.Fatalf("DeleteMin order: got %v..., want %v...", got[:5], want[:5])
	}
}

// checkHeap accepts either order: datagen's heaps pop the largest key first.
func checkHeap[H heapLike[K], K key](t *testing.T, newHeap func() H, keys []K) {
	h := newHeap()
	for _, k := range keys {
		h.Push(k)
	}
	if h.Len() != len(keys) {
		t.Fatalf("Len() = %d, want %d", h.Len(), len(keys))
	}
	got := make([]K, 0, len(keys))
	for h.Len() > 0 {
		got = append(got, h.Pop())
	}
	want := sorted(keys)
	if !slices.Equal(got, want) {
		slices.Reverse(want)
		if !slices.Equal(got, want) {
			t.Fatalf("Pop order isn't sorted: got %v...", got[:5])
		}
	}
}

func checkQueue[Q fifo[K], K key](t *testing.T, newQueue func() Q, keys []K) {
	q := newQueue()
	for _, k := range keys {
		q.Push(k)
	}
	for i, want := range keys {
		if got := q.Pop(); got != want {
			t.Fatalf("Pop() #%d = %v, want %v", i, got, want)
		}
	}
}
