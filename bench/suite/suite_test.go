package suite

import (
	"cmp"
	"math/rand/v2"
	"slices"
	"strconv"
	"testing"
)

// sizes are the number of elements held by a structure while it is
// measured: one that fits in cache, and one that doesn't.
var sizes = []int{1_000, 1_000_000}

// key is the set of key types the suite runs with.
type key interface{ ~int | ~string }

// val is the value stored in sorted maps.
const val = "v"

// intKeys returns 0..n-1 in a random (but deterministic) order.
func intKeys(n int) []int {
	keys := make([]int, n)
	for i := range keys {
		keys[i] = i
	}
	rand.New(rand.NewPCG(1, 2)).Shuffle(n, func(i, j int) { keys[i], keys[j] = keys[j], keys[i] })
	return keys
}

// stringKeys returns intKeys(n) formatted as strings.
func stringKeys(n int) []string {
	keys := make([]string, n)
	for i, k := range intKeys(n) {
		keys[i] = strconv.Itoa(k)
	}
	return keys
}

// impl is one implementation of a datastructure, ready to run any of the
// operations of that datastructure's benchmark, or to check that its adapter
// behaves (see TestImpls).
type impl[K key] struct {
	name  string
	run   func(b *testing.B, op string, keys []K)
	check func(t *testing.T, keys []K)
}

// suite runs every op against every impl, for every size and key type.
// Implementations are the innermost level so that they run back to back.
func suite(b *testing.B, ops []string, ints []impl[int], strs []impl[string]) {
	b.Run("key=int", func(b *testing.B) {
		for _, n := range sizes {
			runOps(b, n, ops, intKeys(n), ints)
		}
	})
	b.Run("key=string", func(b *testing.B) {
		for _, n := range sizes {
			runOps(b, n, ops, stringKeys(n), strs)
		}
	})
}

func runOps[K key](b *testing.B, n int, ops []string, keys []K, impls []impl[K]) {
	b.Run("n="+strconv.Itoa(n), func(b *testing.B) {
		for _, op := range ops {
			b.Run("op="+op, func(b *testing.B) {
				for _, im := range impls {
					b.Run("impl="+im.name, func(b *testing.B) { im.run(b, op, keys) })
				}
			})
		}
	})
}

// perKey reports the time per key for ops where one iteration processes
// every key.
func perKey(b *testing.B, n int) {
	b.ReportMetric(float64(b.Elapsed().Nanoseconds())/float64(b.N)/float64(n), "ns/key")
}

func next(i, n int) int {
	if i++; i == n {
		return 0
	}
	return i
}

func sorted[K cmp.Ordered](keys []K) []K { return slices.Sorted(slices.Values(keys)) }
