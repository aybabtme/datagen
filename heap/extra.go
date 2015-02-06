// Package heap provides a heap container for KType. A heap is a tree
// with the property that each node is the minimum-valued node in its
// subtree.
//
// Heaps are useful when one needs to always retrieve the largest or
// smallest value from a set of values. A common example is in priority
// queues, where elements with the highest priority are pop'd out.
package heap

type KType interface {
	Compare(other KType) int
}
