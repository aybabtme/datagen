package suite

import "testing"

type sortedMap[K key] interface {
	Put(k K, v string)
	Get(k K) (string, bool)
	Delete(k K)
	DeleteMin() K
}

// mapOps, each against a map of n keys:
//   - build-seq, build-rand: insert n keys in sorted/random order into an
//     empty map (sec/op is per map, ns/key per insert).
//   - get: look up a random key.
//   - churn: delete a random key and insert it back.
//   - popmin: delete the smallest key and insert it back.
var mapOps = []string{"build-seq", "build-rand", "get", "churn", "popmin"}

func mapImpl[M sortedMap[K], K key](name string, newMap func() M) impl[K] {
	return impl[K]{
		name:  name,
		run:   func(b *testing.B, op string, keys []K) { benchMap(b, op, newMap, keys) },
		check: func(t *testing.T, keys []K) { checkMap(t, newMap, keys) },
	}
}

func benchMap[M sortedMap[K], K key](b *testing.B, op string, newMap func() M, keys []K) {
	switch op {
	case "build-seq", "build-rand":
		order := keys
		if op == "build-seq" {
			order = sorted(keys)
		}
		for b.Loop() {
			m := newMap()
			for _, k := range order {
				m.Put(k, val)
			}
		}
		perKey(b, len(keys))
		return
	}

	m := newMap()
	for _, k := range keys {
		m.Put(k, val)
	}
	i := 0
	switch op {
	case "get":
		for b.Loop() {
			if _, ok := m.Get(keys[i]); !ok {
				b.Fatalf("key %v missing", keys[i])
			}
			i = next(i, len(keys))
		}
	case "churn":
		for b.Loop() {
			m.Delete(keys[i])
			m.Put(keys[i], val)
			i = next(i, len(keys))
		}
	case "popmin":
		for b.Loop() {
			m.Put(m.DeleteMin(), val)
		}
	default:
		b.Fatalf("unknown op %q", op)
	}
}

var (
	intMaps = []impl[int]{
		mapImpl("own", newOwnIntMap),
		mapImpl("generic", newGenericMap[int]),
		mapImpl("gollrb", newGollrbMap[int]),
		mapImpl("btype", newBtypeMap[int]),
		mapImpl("tidwall", newTidwallMap[int]),
		mapImpl("google", newGoogleMap[int]),
	}
	stringMaps = []impl[string]{
		mapImpl("own", newOwnStringMap),
		mapImpl("generic", newGenericMap[string]),
		mapImpl("gollrb", newGollrbMap[string]),
		mapImpl("btype", newBtypeMap[string]),
		mapImpl("tidwall", newTidwallMap[string]),
		mapImpl("google", newGoogleMap[string]),
	}
)

func BenchmarkSortedMap(b *testing.B) { suite(b, mapOps, intMaps, stringMaps) }
