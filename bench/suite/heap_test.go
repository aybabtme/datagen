package suite

import "testing"

type heapLike[K key] interface {
	Push(k K)
	Pop() K
	Len() int
}

// heapOps, each against a heap of n keys:
//   - sort: push n random keys into an empty heap, then pop them all
//     (sec/op is per heap, ns/key per push+pop).
//   - pushpop: push a random key, then pop the top.
//
// datagen's heaps pop the largest key while the others pop the smallest;
// with random keys the work is the same.
var heapOps = []string{"sort", "pushpop"}

func heapImpl[H heapLike[K], K key](name string, newHeap func() H) impl[K] {
	return impl[K]{
		name:  name,
		run:   func(b *testing.B, op string, keys []K) { benchHeap(b, op, newHeap, keys) },
		check: func(t *testing.T, keys []K) { checkHeap(t, newHeap, keys) },
	}
}

func benchHeap[H heapLike[K], K key](b *testing.B, op string, newHeap func() H, keys []K) {
	switch op {
	case "sort":
		for b.Loop() {
			h := newHeap()
			for _, k := range keys {
				h.Push(k)
			}
			for h.Len() > 0 {
				h.Pop()
			}
		}
		perKey(b, len(keys))
	case "pushpop":
		h := newHeap()
		for _, k := range keys {
			h.Push(k)
		}
		i := 0
		for b.Loop() {
			h.Push(keys[i])
			h.Pop()
			i = next(i, len(keys))
		}
	default:
		b.Fatalf("unknown op %q", op)
	}
}

var (
	intHeaps = []impl[int]{
		heapImpl("own", newOwnIntHeap),
		heapImpl("generic", newGenericHeap[int]),
		heapImpl("stdlib", newStdlibHeap[int]),
		heapImpl("jba", newJbaHeap[int]),
	}
	stringHeaps = []impl[string]{
		heapImpl("own", newOwnStringHeap),
		heapImpl("generic", newGenericHeap[string]),
		heapImpl("stdlib", newStdlibHeap[string]),
		heapImpl("jba", newJbaHeap[string]),
	}
)

func BenchmarkHeap(b *testing.B) { suite(b, heapOps, intHeaps, stringHeaps) }
