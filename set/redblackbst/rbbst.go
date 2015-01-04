package redblackbst

import (
	"bytes"
	"fmt"
	"io"
)

// RedBlack is a sorted map built on a left leaning red black balanced
// search sorted map. It stores VType values, keyed by KType.
type RedBlack struct {
	root *treenode
}

func (r RedBlack) compare(a, b KType) int { return a.Compare(b) }

// NewRedBlack creates a sorted map.
func NewRedBlack() *RedBlack { return &RedBlack{} }

// IsEmpty tells if the sorted map contains no key/value.
func (r RedBlack) IsEmpty() bool {
	return r.root == nil
}

// Size of the sorted map.
func (r RedBlack) Size() int { return r.root.size() }

// Clear all the values in the sorted map.
func (r *RedBlack) Clear() { r.root = nil }

// Put the key `k` in the sorted set. If the value was already there,
// true is returned.
func (r *RedBlack) Put(k KType) (already bool) {
	r.root, already = r.put(r.root, k)
	return
}

func (r *RedBlack) put(h *treenode, k KType) (_ *treenode, already bool) {
	if h == nil {
		n := &treenode{key: k, n: 1, colorRed: true}
		return n, already
	}

	cmp := r.compare(k, h.key)
	if cmp < 0 {
		h.left, already = r.put(h.left, k)
	} else if cmp > 0 {
		h.right, already = r.put(h.right, k)
	} else {
		already = true
	}

	if h.right.isRed() && !h.left.isRed() {
		h = r.rotateLeft(h)
	}
	if h.left.isRed() && h.left.left.isRed() {
		h = r.rotateRight(h)
	}
	if h.left.isRed() && h.right.isRed() {
		r.flipColors(h)
	}
	h.n = h.left.size() + h.right.size() + 1
	return h, already
}

// Contains tells if `k` is a member of the set.
func (r RedBlack) Contains(k KType) bool {
	return r.loopContains(r.root, k)
}

func (r RedBlack) loopContains(h *treenode, k KType) (ok bool) {
	for h != nil {
		cmp := r.compare(k, h.key)
		if cmp == 0 {
			return true
		} else if cmp < 0 {
			h = h.left
		} else if cmp > 0 {
			h = h.right
		}
	}
	return
}

// Min returns the smallest key in the sorted set, if it exists.
func (r RedBlack) Min() (k KType, ok bool) {
	if r.root == nil {
		return
	}
	h := r.min(r.root)
	return h.key, true
}

func (r RedBlack) min(x *treenode) *treenode {
	if x.left == nil {
		return x
	}
	return r.min(x.left)
}

// Max returns the largest key in the sorted set, if it exists.
func (r RedBlack) Max() (k KType, ok bool) {
	if r.root == nil {
		return
	}
	h := r.max(r.root)
	return h.key, true
}

func (r RedBlack) max(x *treenode) *treenode {
	if x.right == nil {
		return x
	}
	return r.max(x.right)
}

// Floor returns the largest key in the sorted set that is smaller than
// `k`.
func (r RedBlack) Floor(key KType) (k KType, ok bool) {
	x := r.floor(r.root, key)
	if x == nil {
		return
	}
	return x.key, true
}

func (r RedBlack) floor(h *treenode, k KType) *treenode {
	if h == nil {
		return nil
	}
	cmp := r.compare(k, h.key)
	if cmp == 0 {
		return h
	}
	if cmp < 0 {
		return r.floor(h.left, k)
	}
	t := r.floor(h.right, k)
	if t != nil {
		return t
	}
	return h
}

// Ceiling returns the smallest key in the sorted set that is larger than
// `k`.
func (r RedBlack) Ceiling(key KType) (k KType, ok bool) {
	x := r.ceiling(r.root, key)
	if x == nil {
		return
	}
	return x.key, true
}

func (r RedBlack) ceiling(h *treenode, k KType) *treenode {
	if h == nil {
		return nil
	}
	cmp := r.compare(k, h.key)
	if cmp == 0 {
		return h
	}
	if cmp > 0 {
		return r.ceiling(h.right, k)
	}
	t := r.ceiling(h.left, k)
	if t != nil {
		return t
	}
	return h
}

// Select key of rank k, meaning the k-th biggest KType in the sorted set.
func (r RedBlack) Select(key int) (k KType, ok bool) {
	x := r.nodeselect(r.root, key)
	if x == nil {
		return
	}
	return x.key, true
}

func (r RedBlack) nodeselect(x *treenode, k int) *treenode {
	if x == nil {
		return nil
	}
	t := x.left.size()
	if t > k {
		return r.nodeselect(x.left, k)
	} else if t < k {
		return r.nodeselect(x.right, k-t-1)
	} else {
		return x
	}
}

// Rank is the number of keys less than `k`.
func (r RedBlack) Rank(k KType) int {
	return r.keyrank(k, r.root)
}

func (r RedBlack) keyrank(k KType, h *treenode) int {
	if h == nil {
		return 0
	}
	cmp := r.compare(k, h.key)
	if cmp < 0 {
		return r.keyrank(k, h.left)
	} else if cmp > 0 {
		return 1 + h.left.size() + r.keyrank(k, h.right)
	} else {
		return h.left.size()
	}
}

// Keys visit each keys in the sorted set, in order.
// It stops when visit returns false.
func (r RedBlack) Keys(visit func(KType) bool) {
	min, ok := r.Min()
	if !ok {
		return
	}
	// if the min exists, then the max must exist
	max, _ := r.Max()
	r.RangedKeys(min, max, visit)
}

// RangedKeys visit each keys between lo and hi in the sorted set, in order.
// It stops when visit returns false.
func (r RedBlack) RangedKeys(lo, hi KType, visit func(KType) bool) {
	r.keys(r.root, visit, lo, hi)
}

func (r RedBlack) keys(h *treenode, visit func(KType) bool, lo, hi KType) bool {
	if h == nil {
		return true
	}
	cmplo := r.compare(lo, h.key)
	cmphi := r.compare(hi, h.key)
	if cmplo < 0 {
		if !r.keys(h.left, visit, lo, hi) {
			return false
		}
	}
	if cmplo <= 0 && cmphi >= 0 {
		if !visit(h.key) {
			return false
		}
	}
	if cmphi > 0 {
		if !r.keys(h.right, visit, lo, hi) {
			return false
		}
	}
	return true
}

// DeleteMin removes the smallest key from the sorted set.
func (r *RedBlack) DeleteMin() (oldk KType, ok bool) {
	r.root, oldk, ok = r.deleteMin(r.root)
	if !r.IsEmpty() {
		r.root.colorRed = false
	}
	return
}

func (r *RedBlack) deleteMin(h *treenode) (_ *treenode, oldk KType, ok bool) {
	if h == nil {
		return nil, oldk, false
	}

	if h.left == nil {
		return nil, h.key, true
	}
	if !h.left.isRed() && !h.left.left.isRed() {
		h = r.moveRedLeft(h)
	}
	h.left, oldk, ok = r.deleteMin(h.left)
	return r.balance(h), oldk, ok
}

// DeleteMax removes the largest key from the sorted set.
func (r *RedBlack) DeleteMax() (oldk KType, ok bool) {
	r.root, oldk, ok = r.deleteMax(r.root)
	if !r.IsEmpty() {
		r.root.colorRed = false
	}
	return
}

func (r *RedBlack) deleteMax(h *treenode) (_ *treenode, oldk KType, ok bool) {
	if h == nil {
		return nil, oldk, ok
	}
	if h.left.isRed() {
		h = r.rotateRight(h)
	}
	if h.right == nil {
		return nil, h.key, true
	}
	if !h.right.isRed() && !h.right.left.isRed() {
		h = r.moveRedRight(h)
	}
	h.right, oldk, ok = r.deleteMax(h.right)
	return r.balance(h), oldk, ok
}

// Delete key `k` from sorted set, if it exists.
func (r *RedBlack) Delete(k KType) (ok bool) {
	if r.root == nil {
		return
	}
	r.root, ok = r.delete(r.root, k)
	if !r.IsEmpty() {
		r.root.colorRed = false
	}
	return
}

func (r *RedBlack) delete(h *treenode, k KType) (_ *treenode, ok bool) {

	if h == nil {
		return h, false
	}

	if r.compare(k, h.key) < 0 {
		if h.left == nil {
			return h, false
		}

		if !h.left.isRed() && !h.left.left.isRed() {
			h = r.moveRedLeft(h)
		}

		h.left, ok = r.delete(h.left, k)
		h = r.balance(h)
		return h, ok
	}

	if h.left.isRed() {
		h = r.rotateRight(h)
	}

	if r.compare(k, h.key) == 0 && h.right == nil {
		return nil, true
	}

	if h.right != nil && !h.right.isRed() && !h.right.left.isRed() {
		h = r.moveRedRight(h)
	}

	if r.compare(k, h.key) == 0 {

		var subk KType
		h.right, subk, ok = r.deleteMin(h.right)
		h.key = subk
		ok = true
	} else {
		h.right, ok = r.delete(h.right, k)
	}

	h = r.balance(h)
	return h, ok
}

// deletions

func (r *RedBlack) moveRedLeft(h *treenode) *treenode {
	r.flipColors(h)
	if h.right.left.isRed() {
		h.right = r.rotateRight(h.right)
		h = r.rotateLeft(h)
		r.flipColors(h)
	}
	return h
}

func (r *RedBlack) moveRedRight(h *treenode) *treenode {
	r.flipColors(h)
	if h.left.left.isRed() {
		h = r.rotateRight(h)
		r.flipColors(h)
	}
	return h
}

func (r *RedBlack) balance(h *treenode) *treenode {
	if h.right.isRed() {
		h = r.rotateLeft(h)
	}
	if h.left.isRed() && h.left.left.isRed() {
		h = r.rotateRight(h)
	}
	if h.left.isRed() && h.right.isRed() {
		r.flipColors(h)
	}
	h.n = h.left.size() + h.right.size() + 1
	return h
}

func (r *RedBlack) rotateLeft(h *treenode) *treenode {
	x := h.right
	h.right = x.left
	x.left = h
	x.colorRed = h.colorRed
	h.colorRed = true
	x.n = h.n
	h.n = 1 + h.left.size() + h.right.size()
	return x
}

func (r *RedBlack) rotateRight(h *treenode) *treenode {
	x := h.left
	h.left = x.right
	x.right = h
	x.colorRed = h.colorRed
	h.colorRed = true
	x.n = h.n
	h.n = 1 + h.left.size() + h.right.size()
	return x
}

func (r *RedBlack) flipColors(h *treenode) {
	h.colorRed = !h.colorRed
	h.left.colorRed = !h.left.colorRed
	h.right.colorRed = !h.right.colorRed
}

// nodes

type treenode struct {
	key         KType
	left, right *treenode
	n           int
	colorRed    bool
}

func (x *treenode) isRed() bool { return (x != nil) && (x.colorRed == true) }

func (x *treenode) size() int {
	if x == nil {
		return 0
	}
	return x.n
}

// debugging

// DotGraph exports the sorted map into DOT format.
func (r RedBlack) DotGraph(out io.Writer, name string) (int, error) {
	return r.dotGraph(r.root, out, name)
}

func (r RedBlack) dotGraph(h *treenode, out io.Writer, name string) (n int, err error) {
	nodes := bytes.NewBuffer(nil)
	edges := bytes.NewBuffer(nil)

	fmt.Fprintf(nodes, "digraph %q {\n", name)
	fmt.Fprintf(edges, "\t%q -> ", name)
	r.dotvisit(h, nodes, edges, true)
	fmt.Fprintf(edges, "}\n")

	edges.WriteTo(nodes)

	return out.Write(nodes.Bytes())
}

func (r RedBlack) dotvisit(x *treenode, nodes io.Writer, edges *bytes.Buffer, isLeft bool) {

	var color string
	if x.isRed() {
		color = "red"
	} else {
		color = "black"
	}

	var direction string
	if isLeft {
		direction = "left"
	} else {
		direction = "right"
	}

	if x == nil {
		fmt.Fprintf(edges, "nil [label=%q, color=%s];\n", direction, color)
		return
	}

	fmt.Fprintf(edges, "\"%p\" [label=%q, color=%s];\n", x, direction, color)
	fmt.Fprintf(nodes, "\t\"%p\" [label=\"%v\", shape = circle, color=%s];\n", x, x.key, color)

	fmt.Fprintf(edges, "\t\"%p\" -> ", x)
	r.dotvisit(x.left, nodes, edges, true)
	fmt.Fprintf(edges, "\t\"%p\" -> ", x)
	r.dotvisit(x.right, nodes, edges, false)
}
