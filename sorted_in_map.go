package datagen

import (
	"bytes"
	"fmt"
	"io"
)

// SortedIntToIntMap is a sorted map built on a left leaning red black balanced
// search sorted map. It stores int values, keyed by int.
type SortedIntToIntMap struct {
	root *node
}

func (r SortedIntToIntMap) compareint(a, b int) int { return a - b }

// NewSortedIntToIntMap creates a sorted map.
func NewSortedIntToIntMap() *SortedIntToIntMap { return &SortedIntToIntMap{} }

// IsEmpty tells if the sorted map contains no key/value.
func (r SortedIntToIntMap) IsEmpty() bool {
	return r.root == nil
}

// Size of the sorted map.
func (r SortedIntToIntMap) Size() int { return r.root.size() }

// Clear all the values in the sorted map.
func (r *SortedIntToIntMap) Clear() { r.root = nil }

// Put a value in the sorted map at key `k`. The old value at `k` is returned
// if the key was already present.
func (r *SortedIntToIntMap) Put(k int, v int) (old int, overwrite bool) {
	r.root, old, overwrite = r.put(r.root, k, v)
	return
}

func (r *SortedIntToIntMap) put(h *node, k int, v int) (_ *node, old int, overwrite bool) {
	if h == nil {
		n := &node{key: k, val: v, n: 1, color: red}
		return n, old, overwrite
	}

	cmp := r.compareint(k, h.key)
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
func (r SortedIntToIntMap) Get(k int) (int, bool) {
	return r.loopGet(r.root, k)
}

func (r SortedIntToIntMap) loopGet(h *node, k int) (v int, ok bool) {
	for h != nil {
		cmp := r.compareint(k, h.key)
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
func (r SortedIntToIntMap) Has(k int) bool {
	_, ok := r.loopGet(r.root, k)
	return ok
}

// Min returns the smallest key/value in the sorted map, if it exists.
func (r SortedIntToIntMap) Min() (k int, v int, ok bool) {
	if r.root == nil {
		return
	}
	h := r.min(r.root)
	return h.key, h.val, true
}

func (r SortedIntToIntMap) min(x *node) *node {
	if x.left == nil {
		return x
	}
	return r.min(x.left)
}

// Max returns the largest key/value in the sorted map, if it exists.
func (r SortedIntToIntMap) Max() (k int, v int, ok bool) {
	if r.root == nil {
		return
	}
	h := r.max(r.root)
	return h.key, h.val, true
}

func (r SortedIntToIntMap) max(x *node) *node {
	if x.right == nil {
		return x
	}
	return r.max(x.right)
}

// Floor returns the largest key/value in the sorted map that is smaller than
// `k`.
func (r SortedIntToIntMap) Floor(key int) (k int, v int, ok bool) {
	x := r.floor(r.root, key)
	if x == nil {
		return
	}
	return x.key, x.val, true
}

func (r SortedIntToIntMap) floor(h *node, k int) *node {
	if h == nil {
		return nil
	}
	cmp := r.compareint(k, h.key)
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
func (r SortedIntToIntMap) Ceiling(key int) (k int, v int, ok bool) {
	x := r.ceiling(r.root, key)
	if x == nil {
		return
	}
	return x.key, x.val, true
}

func (r SortedIntToIntMap) ceiling(h *node, k int) *node {
	if h == nil {
		return nil
	}
	cmp := r.compareint(k, h.key)
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

// Select key of rank k, meaning the k-th biggest int in the sorted map.
func (r SortedIntToIntMap) Select(key int) (k int, v int, ok bool) {
	x := r.nodeselect(r.root, key)
	if x == nil {
		return
	}
	return x.key, x.val, true
}

func (r SortedIntToIntMap) nodeselect(x *node, k int) *node {
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
func (r SortedIntToIntMap) Rank(k int) int {
	return r.keyrank(k, r.root)
}

func (r SortedIntToIntMap) keyrank(k int, h *node) int {
	if h == nil {
		return 0
	}
	cmp := r.compareint(k, h.key)
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
func (r SortedIntToIntMap) Keys(visit func(int, int) bool) {
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
func (r SortedIntToIntMap) RangedKeys(lo, hi int, visit func(int, int) bool) {
	r.keys(r.root, visit, lo, hi)
}

func (r SortedIntToIntMap) keys(h *node, visit func(int, int) bool, lo, hi int) bool {
	if h == nil {
		return true
	}
	cmplo := r.compareint(lo, h.key)
	cmphi := r.compareint(hi, h.key)
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
func (r *SortedIntToIntMap) DeleteMin() (oldk int, oldv int, ok bool) {
	r.root, oldk, oldv, ok = r.deleteMin(r.root)
	if !r.IsEmpty() {
		r.root.color = black
	}
	return
}

func (r *SortedIntToIntMap) deleteMin(h *node) (_ *node, oldk int, oldv int, ok bool) {
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
func (r *SortedIntToIntMap) DeleteMax() (oldk int, oldv int, ok bool) {
	r.root, oldk, oldv, ok = r.deleteMax(r.root)
	if !r.IsEmpty() {
		r.root.color = black
	}
	return
}

func (r *SortedIntToIntMap) deleteMax(h *node) (_ *node, oldk int, oldv int, ok bool) {
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
func (r *SortedIntToIntMap) Delete(k int) (old int, ok bool) {
	if r.root == nil {
		return
	}
	r.root, old, ok = r.delete(r.root, k)
	if !r.IsEmpty() {
		r.root.color = black
	}
	return
}

func (r *SortedIntToIntMap) delete(h *node, k int) (_ *node, old int, ok bool) {

	if h == nil {
		return h, old, false
	}

	if r.compareint(k, h.key) < 0 {
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

	if r.compareint(k, h.key) == 0 && h.right == nil {
		return nil, h.val, true
	}

	if h.right != nil && !h.right.isRed() && !h.right.left.isRed() {
		h = r.moveRedRight(h)
	}

	if r.compareint(k, h.key) == 0 {

		var subk int
		var subv int
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

func (r *SortedIntToIntMap) moveRedLeft(h *node) *node {
	r.flipColors(h)
	if h.right.left.isRed() {
		h.right = r.rotateRight(h.right)
		h = r.rotateLeft(h)
		r.flipColors(h)
	}
	return h
}

func (r *SortedIntToIntMap) moveRedRight(h *node) *node {
	r.flipColors(h)
	if h.left.left.isRed() {
		h = r.rotateRight(h)
		r.flipColors(h)
	}
	return h
}

func (r *SortedIntToIntMap) balance(h *node) *node {
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

// nodes

const (
	red   = true
	black = false
)

type node struct {
	key         int
	val         int
	left, right *node
	n           int
	color       bool
}

func newNode(k int, v int, n int, color bool) *node {
	return &node{key: k, val: v, n: n, color: color}
}

func (x *node) isRed() bool { return (x != nil) && (x.color == red) }

func (r *SortedIntToIntMap) rotateLeft(h *node) *node {
	x := h.right
	h.right = x.left
	x.left = h
	x.color = h.color
	h.color = red
	x.n = h.n
	h.n = 1 + h.left.size() + h.right.size()
	return x
}

func (r *SortedIntToIntMap) rotateRight(h *node) *node {
	x := h.left
	h.left = x.right
	x.right = h
	x.color = h.color
	h.color = red
	x.n = h.n
	h.n = 1 + h.left.size() + h.right.size()
	return x
}

func (r *SortedIntToIntMap) flipColors(h *node) {
	h.color = !h.color
	h.left.color = !h.left.color
	h.right.color = !h.right.color
}

func (x *node) size() int {
	if x == nil {
		return 0
	}
	return x.n
}

// debugging

// DotGraph exports the sorted map into DOT format.
func (r SortedIntToIntMap) DotGraph(out io.Writer, name string) (int, error) {
	return r.dotGraph(r.root, out, name)
}

func (r SortedIntToIntMap) dotGraph(h *node, out io.Writer, name string) (n int, err error) {
	var m int
	buf := bytes.NewBuffer(nil)
	m, err = fmt.Fprintf(out, "digraph %q {\n", name)
	n += m
	if err != nil {
		return
	}
	fmt.Fprintf(buf, "\t%q -> ", name)
	m, err = r.dotvisit(h, out, buf, true)
	n += m
	if err != nil {
		return
	}

	var k int64
	k, err = buf.WriteTo(out)
	n += int(k)
	if err != nil {
		return
	}

	m, err = fmt.Fprintf(out, "}\n")
	n += m
	return
}

func (r SortedIntToIntMap) dotvisit(x *node, nodes io.Writer, edges *bytes.Buffer, isLeft bool) (n int, err error) {

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

	var m int
	m, err = fmt.Fprintf(nodes, "\t\"%p\" [label=\"%v\", shape = circle, color=%s];\n", x, x.key, color)
	n += m
	if err != nil {
		return
	}

	fmt.Fprintf(edges, "\t\"%p\" -> ", x)
	m, err = r.dotvisit(x.left, nodes, edges, true)
	n += m
	if err != nil {
		return
	}

	fmt.Fprintf(edges, "\t\"%p\" -> ", x)
	m, err = r.dotvisit(x.right, nodes, edges, false)
	n += m
	return
}
