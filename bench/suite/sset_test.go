package suite

import "testing"

type sortedSet[K key] interface {
	Put(k K)
	Has(k K) bool
	Delete(k K)
	DeleteMin() K
}

// setOps are the same as mapOps, with has in place of get.
var setOps = []string{"build-seq", "build-rand", "has", "churn", "popmin"}

func setImpl[S sortedSet[K], K key](name string, newSet func() S) impl[K] {
	return impl[K]{
		name:  name,
		run:   func(b *testing.B, op string, keys []K) { benchSet(b, op, newSet, keys) },
		check: func(t *testing.T, keys []K) { checkSet(t, newSet, keys) },
	}
}

func benchSet[S sortedSet[K], K key](b *testing.B, op string, newSet func() S, keys []K) {
	switch op {
	case "build-seq", "build-rand":
		order := keys
		if op == "build-seq" {
			order = sorted(keys)
		}
		for b.Loop() {
			s := newSet()
			for _, k := range order {
				s.Put(k)
			}
		}
		perKey(b, len(keys))
		return
	}

	s := newSet()
	for _, k := range keys {
		s.Put(k)
	}
	i := 0
	switch op {
	case "has":
		for b.Loop() {
			if !s.Has(keys[i]) {
				b.Fatalf("key %v missing", keys[i])
			}
			i = next(i, len(keys))
		}
	case "churn":
		for b.Loop() {
			s.Delete(keys[i])
			s.Put(keys[i])
			i = next(i, len(keys))
		}
	case "popmin":
		for b.Loop() {
			s.Put(s.DeleteMin())
		}
	default:
		b.Fatalf("unknown op %q", op)
	}
}

var (
	intSets = []impl[int]{
		setImpl("own", newOwnIntSet),
		setImpl("generic", newGenericSet[int]),
		setImpl("gollrb", newGollrbSet[int]),
		setImpl("btype", newBtypeSet[int]),
		setImpl("tidwall", newTidwallSet[int]),
		setImpl("google", newGoogleSet[int]),
	}
	stringSets = []impl[string]{
		setImpl("own", newOwnStringSet),
		setImpl("generic", newGenericSet[string]),
		setImpl("gollrb", newGollrbSet[string]),
		setImpl("btype", newBtypeSet[string]),
		setImpl("tidwall", newTidwallSet[string]),
		setImpl("google", newGoogleSet[string]),
	}
)

func BenchmarkSortedSet(b *testing.B) { suite(b, setOps, intSets, stringSets) }
