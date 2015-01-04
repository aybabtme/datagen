package redblackbst

import (
	"bytes"
	"fmt"
	"io"
)

// RedBlack is a sorted map built on a left leaning red black balanced
// search sorted map. It stores VType values, keyed by KType.
type RedBlack struct {
	root *node
}

func (r RedBlack) compareKType(a, b KType) int { return a.Compare(b) }

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

// Put a value in the sorted map at key `k`. The old value at `k` is returned
// if the key was already present.
func (r *RedBlack) Put(k KType, v VType) (old VType, overwrite bool) {
	r.root, old, overwrite = r.put(r.root, k, v)
	return
}

func (r *RedBlack) put(h *node, k KType, v VType) (_ *node, old VType, overwrite bool) {
	if h == nil {
		n := &node{key: k, val: v, n: 1, color: red}
		return n, old, overwrite
	}

	cmp := r.compareKType(k, h.key)
	if cmp < 0 {
		h.left, old, overwrite = r.put(h.left, k, v)
	} else if cmp > 0 {
		h.right, old, overwrite = r.put(h.right, k, v)
	} else {
		overwrite = true
		old = h.val
		h.val = v
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
	return h, old, overwrite
}

// Get a value from the sorted map at key `k`. Returns false
// if the key doesn't exist.
func (r RedBlack) Get(k KType) (VType, bool) {
	return r.loopGet(r.root, k)
}

func (r RedBlack) loopGet(h *node, k KType) (v VType, ok bool) {
	for h != nil {
		cmp := r.compareKType(k, h.key)
		if cmp == 0 {
			return h.val, true
		} else if cmp < 0 {
			h = h.left
		} else if cmp > 0 {
			h = h.right
		}
	}
	return
}

// Has tells if a value exists at key `k`. This is short hand for `Get.
func (r RedBlack) Has(k KType) bool {
	_, ok := r.loopGet(r.root, k)
	return ok
}

// Min returns the smallest key/value in the sorted map, if it exists.
func (r RedBlack) Min() (k KType, v VType, ok bool) {
	if r.root == nil {
		return
	}
	h := r.min(r.root)
	return h.key, h.val, true
}

func (r RedBlack) min(x *node) *node {
	if x.left == nil {
		return x
	}
	return r.min(x.left)
}

// Max returns the largest key/value in the sorted map, if it exists.
func (r RedBlack) Max() (k KType, v VType, ok bool) {
	if r.root == nil {
		return
	}
	h := r.max(r.root)
	return h.key, h.val, true
}

func (r RedBlack) max(x *node) *node {
	if x.right == nil {
		return x
	}
	return r.max(x.right)
}

// Floor returns the largest key/value in the sorted map that is smaller than
// `k`.
func (r RedBlack) Floor(key KType) (k KType, v VType, ok bool) {
	x := r.floor(r.root, key)
	if x == nil {
		return
	}
	return x.key, x.val, true
}

func (r RedBlack) floor(h *node, k KType) *node {
	if h == nil {
		return nil
	}
	cmp := r.compareKType(k, h.key)
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

// Ceiling returns the smallest key/value in the sorted map that is larger than
// `k`.
func (r RedBlack) Ceiling(key KType) (k KType, v VType, ok bool) {
	x := r.ceiling(r.root, key)
	if x == nil {
		return
	}
	return x.key, x.val, true
}

func (r RedBlack) ceiling(h *node, k KType) *node {
	if h == nil {
		return nil
	}
	cmp := r.compareKType(k, h.key)
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

// Select key of rank k, meaning the k-th biggest KType in the sorted map.
func (r RedBlack) Select(key int) (k KType, v VType, ok bool) {
	x := r.nodeselect(r.root, key)
	if x == nil {
		return
	}
	return x.key, x.val, true
}

func (r RedBlack) nodeselect(x *node, k int) *node {
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

func (r RedBlack) keyrank(k KType, h *node) int {
	if h == nil {
		return 0
	}
	cmp := r.compareKType(k, h.key)
	if cmp < 0 {
		return r.keyrank(k, h.left)
	} else if cmp > 0 {
		return 1 + h.left.size() + r.keyrank(k, h.right)
	} else {
		return h.left.size()
	}
}

// Keys visit each keys in the sorted map, in order.
// It stops when visit returns false.
func (r RedBlack) Keys(visit func(KType, VType) bool) {
	min, _, ok := r.Min()
	if !ok {
		return
	}
	// if the min exists, then the max must exist
	max, _, _ := r.Max()
	r.RangedKeys(min, max, visit)
}

// RangedKeys visit each keys between lo and hi in the sorted map, in order.
// It stops when visit returns false.
func (r RedBlack) RangedKeys(lo, hi KType, visit func(KType, VType) bool) {
	r.keys(r.root, visit, lo, hi)
}

func (r RedBlack) keys(h *node, visit func(KType, VType) bool, lo, hi KType) bool {
	if h == nil {
		return true
	}
	cmplo := r.compareKType(lo, h.key)
	cmphi := r.compareKType(hi, h.key)
	if cmplo < 0 {
		if !r.keys(h.left, visit, lo, hi) {
			return false
		}
	}
	if cmplo <= 0 && cmphi >= 0 {
		if !visit(h.key, h.val) {
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

// DeleteMin removes the smallest key and its value from the sorted map.
func (r *RedBlack) DeleteMin() (oldk KType, oldv VType, ok bool) {
	r.root, oldk, oldv, ok = r.deleteMin(r.root)
	if !r.IsEmpty() {
		r.root.color = black
	}
	return
}

func (r *RedBlack) deleteMin(h *node) (_ *node, oldk KType, oldv VType, ok bool) {
	if h == nil {
		return nil, oldk, oldv, false
	}

	if h.left == nil {
		return nil, h.key, h.val, true
	}
	if !h.left.isRed() && !h.left.left.isRed() {
		h = r.moveRedLeft(h)
	}
	h.left, oldk, oldv, ok = r.deleteMin(h.left)
	return r.balance(h), oldk, oldv, ok
}

// DeleteMax removes the largest key and its value from the sorted map.
func (r *RedBlack) DeleteMax() (oldk KType, oldv VType, ok bool) {
	r.root, oldk, oldv, ok = r.deleteMax(r.root)
	if !r.IsEmpty() {
		r.root.color = black
	}
	return
}

func (r *RedBlack) deleteMax(h *node) (_ *node, oldk KType, oldv VType, ok bool) {
	if h == nil {
		return nil, oldk, oldv, ok
	}
	if h.left.isRed() {
		h = r.rotateRight(h)
	}
	if h.right == nil {
		return nil, h.key, h.val, true
	}
	if !h.right.isRed() && !h.right.left.isRed() {
		h = r.moveRedRight(h)
	}
	h.right, oldk, oldv, ok = r.deleteMax(h.right)
	return r.balance(h), oldk, oldv, ok
}

// Delete key `k` from sorted map, if it exists.
func (r *RedBlack) Delete(k KType) (old VType, ok bool) {
	if r.root == nil {
		return
	}
	r.root, old, ok = r.delete(r.root, k)
	if !r.IsEmpty() {
		r.root.color = black
	}
	return
}

func (r *RedBlack) delete(h *node, k KType) (_ *node, old VType, ok bool) {

	if h == nil {
		return h, old, false
	}

	if r.compareKType(k, h.key) < 0 {
		if h.left == nil {
			return h, old, false
		}

		if !h.left.isRed() && !h.left.left.isRed() {
			h = r.moveRedLeft(h)
		}

		h.left, old, ok = r.delete(h.left, k)
		h = r.balance(h)
		return h, old, ok
	}

	if h.left.isRed() {
		h = r.rotateRight(h)
	}

	if r.compareKType(k, h.key) == 0 && h.right == nil {
		return nil, h.val, true
	}

	if h.right != nil && !h.right.isRed() && !h.right.left.isRed() {
		h = r.moveRedRight(h)
	}

	if r.compareKType(k, h.key) == 0 {

		var subk KType
		var subv VType
		h.right, subk, subv, ok = r.deleteMin(h.right)

		old, h.key, h.val = h.val, subk, subv
		ok = true
	} else {
		h.right, old, ok = r.delete(h.right, k)
	}

	h = r.balance(h)
	return h, old, ok
}

// deletions

func (r *RedBlack) moveRedLeft(h *node) *node {
	r.flipColors(h)
	if h.right.left.isRed() {
		h.right = r.rotateRight(h.right)
		h = r.rotateLeft(h)
		r.flipColors(h)
	}
	return h
}

func (r *RedBlack) moveRedRight(h *node) *node {
	r.flipColors(h)
	if h.left.left.isRed() {
		h = r.rotateRight(h)
		r.flipColors(h)
	}
	return h
}

func (r *RedBlack) balance(h *node) *node {
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

func (r *RedBlack) rotateLeft(h *node) *node {
	x := h.right
	h.right = x.left
	x.left = h
	x.color = h.color
	h.color = red
	x.n = h.n
	h.n = 1 + h.left.size() + h.right.size()
	return x
}

func (r *RedBlack) rotateRight(h *node) *node {
	x := h.left
	h.left = x.right
	x.right = h
	x.color = h.color
	h.color = red
	x.n = h.n
	h.n = 1 + h.left.size() + h.right.size()
	return x
}

func (r *RedBlack) flipColors(h *node) {
	h.color = !h.color
	h.left.color = !h.left.color
	h.right.color = !h.right.color
}

// nodes

const (
	red   = true
	black = false
)

type node struct {
	key         KType
	val         VType
	left, right *node
	n           int
	color       bool
}

func (x *node) isRed() bool { return (x != nil) && (x.color == red) }

func (x *node) size() int {
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

func (r RedBlack) dotGraph(h *node, out io.Writer, name string) (n int, err error) {
	nodes := bytes.NewBuffer(nil)
	edges := bytes.NewBuffer(nil)

	fmt.Fprintf(nodes, "digraph %q {\n", name)
	fmt.Fprintf(edges, "\t%q -> ", name)
	r.dotvisit(h, nodes, edges, true)
	fmt.Fprintf(edges, "}\n")

	edges.WriteTo(nodes)

	return out.Write(nodes.Bytes())
}

func (r RedBlack) dotvisit(x *node, nodes io.Writer, edges *bytes.Buffer, isLeft bool) {

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
