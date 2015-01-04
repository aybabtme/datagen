// Package redblackbst implements a red black balanced search tree,
// based on the details provided in Algorithms 4th edition, by
// Robert Sedgewick and Kevin Wayne.
//
// A red black bst is useful as a map that keeps its items in
// sorted order, while preserving efficient inserts, lookups and
// deletions.
//
// Some tests were extracted from GoLLRB, a similar implementation by
// Petar Maymounkov.
package redblackbst

// ugly type names to avoid collisions, for easy find/replace.

type KType interface {
	Compare(other KType) int
}
