// Package redblackbst implements a red black balanced search tree,
// based on the details provided in Algorithms 4th edition, by
// Robert Sedgewick and Kevin Wayne.
package redblackbst

// ugly type names to avoid collisions, for easy find/replace.

type KType interface {
	Compare(other KType) int
}
type VType interface{}

// RedBlack holds the state of a red black balanced search tree.
type RedBlack struct {
	root *node
}

// New creates a red black balanced search tree.
func New() *RedBlack { return &RedBlack{} }

// IsEmpty tells if the tree contains no key/value.
func (r RedBlack) IsEmpty() bool {
	return r.root == nil
}

// Size of the tree.
func (r RedBlack) Size() int { return size(r.root) }

// Put a value in the tree at key `k`.
func (r *RedBlack) Put(k KType, v VType) {
	r.root = put(r.root, k, v)
}

// Get a value from the tree at key `k`. Returns false
// if the key doesn't exist.
func (r RedBlack) Get(k KType) (VType, bool) {
	return loopGet(r.root, k)
	//return recurGet(r.root, k)
}

// Min returns the smallest key/value in the tree, if it exists.
func (r RedBlack) Min() (KType, VType, bool) {
	h := min(r.root)
	if h == nil {
		return nil, nil, false
	}
	return h.key, h.val, true
}

// Max returns the largest key/value in the tree, if it exists.
func (r RedBlack) Max() (KType, VType, bool) {
	h := max(r.root)
	if h == nil {
		return nil, nil, false
	}
	return h.key, h.val, true
}

// Floor returns the largest key/value in the tree that is smaller than
// `k`.
func (r RedBlack) Floor(k KType) (KType, VType, bool) {
	x := floor(r.root, k)
	if x == nil {
		return nil, nil, false
	}
	return x.key, x.val, true
}

// Ceiling returns the smallest key/value in the tree that is larger than
// `k`.
func (r RedBlack) Ceiling(k KType) (KType, VType, bool) {
	x := ceiling(r.root, k)
	if x == nil {
		return nil, nil, false
	}
	return x.key, x.val, true
}

// Select key of rank k, meaning the k-th biggest KType in the tree.
func (r RedBlack) Select(k int) (KType, VType, bool) {
	x := nodeselect(r.root, k)
	if x == nil {
		return nil, nil, false
	}
	return x.key, x.val, true
}

// Rank is the number of keys less than `k`.
func (r RedBlack) Rank(k KType) int {
	return keyrank(k, r.root)
}

// Keys visit each keys in the tree, in order.
// It stops when visit returns false.
func (r RedBlack) Keys(visit func(KType, VType) bool) {
	min, _, ok := r.Min()
	if !ok {
		return
	}
	max, _, ok := r.Max()
	if !ok {
		return
	}
	r.RangedKeys(min, max, visit)
}

// RangedKeys visit each keys between lo and hi in the tree, in order.
// It stops when visit returns false.
func (r RedBlack) RangedKeys(lo, hi KType, visit func(KType, VType) bool) {
	keys(r.root, visit, lo, hi)
}

// DeleteMin removes the smallest key and its value from the tree.
func (r *RedBlack) DeleteMin() {
	if !isRed(r.root.left) && !isRed(r.root.right) {
		r.root.color = red
	}
	r.root = deleteMin(r.root)
	if !r.IsEmpty() {
		r.root.color = black
	}
}

// DeleteMax removes the largest key and its value from the tree.
func (r *RedBlack) DeleteMax() {
	if !isRed(r.root.left) && !isRed(r.root.right) {
		r.root.color = red
	}
	r.root = deleteMax(r.root)
	if !r.IsEmpty() {
		r.root.color = black
	}
}

// Delete key `k` from tree, if it exists.
func (r *RedBlack) Delete(k KType) (ok bool) {
	if !isRed(r.root.left) && !isRed(r.root.right) {
		r.root.color = red
	}
	r.root, ok = delete(r.root, k)
	if !r.IsEmpty() {
		r.root.color = black
	}
	return
}

// Clear all the values in the tree.
func (r *RedBlack) Clear() { r.root = nil }

// maximum,minimum,floor,ceiling,select,rank,range

func put(h *node, k KType, v VType) *node {
	if h == nil {
		return newNode(k, v, 1, red)
	}

	cmp := k.Compare(h.key)
	if cmp < 0 {
		h.left = put(h.left, k, v)
	} else if cmp > 0 {
		h.right = put(h.right, k, v)
	} else {
		h.val = v
	}

	if isRed(h.right) && !isRed(h.left) {
		h = rotateLeft(h)
	}
	if isRed(h.left) && isRed(h.left.left) {
		h = rotateRight(h)
	}
	if isRed(h.left) && isRed(h.right) {
		flipColors(h)
	}
	h.n = size(h.left) + size(h.right) + 1
	return h
}

func loopGet(x *node, k KType) (VType, bool) {
	for x != nil {
		cmp := k.Compare(x.key)
		if cmp == 0 {
			return x.val, true
		} else if cmp < 0 {
			x = x.left
		} else if cmp > 0 {
			x = x.right
		}
	}
	return nil, false
}

func recurGet(x *node, k KType) (VType, bool) {
	if x == nil {
		return nil, false
	}
	cmp := k.Compare(x.key)
	if cmp == 0 {
		return x.val, true
	} else if cmp < 0 {
		return recurGet(x.left, k)
	} else if cmp > 0 {
		return recurGet(x.right, k)
	}
	panic("unreachable")
}

func min(x *node) *node {
	if x.left == nil {
		return x
	}
	return min(x.left)
}

func max(x *node) *node {
	if x.right == nil {
		return x
	}
	return max(x.right)
}

func floor(x *node, k KType) *node {
	if x == nil {
		return nil
	}
	cmp := k.Compare(x.key)
	if cmp == 0 {
		return x
	}
	if cmp < 0 {
		return floor(x.left, k)
	}
	t := floor(x.right, k)
	if t != nil {
		return t
	}
	return x
}

func ceiling(x *node, k KType) *node {
	if x == nil {
		return nil
	}
	cmp := k.Compare(x.key)
	if cmp == 0 {
		return x
	}
	if cmp > 0 {
		return ceiling(x.right, k)
	}
	t := ceiling(x.left, k)
	if t != nil {
		return t
	}
	return x
}

func nodeselect(x *node, k int) *node {
	if x == nil {
		return nil
	}
	t := size(x.left)
	if t > k {
		return nodeselect(x.left, k)
	} else if t < k {
		return nodeselect(x.right, k-t-1)
	} else {
		return x
	}
}

func keyrank(k KType, x *node) int {
	if x == nil {
		return 0
	}
	cmp := k.Compare(x.key)
	if cmp < 0 {
		return keyrank(k, x.left)
	} else if cmp > 0 {
		return 1 + size(x.left) + keyrank(k, x.right)
	} else {
		return size(x.left)
	}
}

func keys(x *node, visit func(KType, VType) bool, lo, hi KType) bool {
	if x == nil {
		return true
	}
	cmplo := lo.Compare(x.key)
	cmphi := hi.Compare(x.key)
	if cmplo < 0 {
		if !keys(x.left, visit, lo, hi) {
			return false
		}
	}
	if cmplo <= 0 && cmphi >= 0 {
		if !visit(x.key, x.val) {
			return false
		}
	}
	if cmphi > 0 {
		if !keys(x.right, visit, lo, hi) {
			return false
		}
	}
	return true
}

// deletions

func moveRedLeft(h *node) *node {
	flipColors(h)
	if isRed(h.right.left) {
		h.right = rotateRight(h.right)
		h = rotateLeft(h)
	}
	return h
}

func moveRedRight(h *node) *node {
	flipColors(h)
	if !isRed(h.left.left) {
		h = rotateRight(h)
	}
	return h
}

func deleteMin(h *node) *node {
	if h.left == nil {
		return nil
	}
	if !isRed(h.left) && !isRed(h.left.left) {
		h = moveRedLeft(h)
	}
	h.left = deleteMin(h.left)
	return balance(h)
}

func deleteMax(h *node) *node {
	if isRed(h.left) {
		h = rotateRight(h)
	}
	if h.right == nil {
		return nil
	}
	if !isRed(h.right) && !isRed(h.right.left) {
		h = moveRedRight(h)
	}
	h.right = deleteMax(h.right)
	return balance(h)
}

func delete(h *node, k KType) (*node, bool) {
	if k.Compare(h.key) < 0 {
		if !isRed(h.left) && !isRed(h.left.left) {
			h = moveRedLeft(h)
		}
		var ok bool
		h.left, ok = delete(h.left, k)
		return balance(h), ok
	}

	if isRed(h.left) {
		h = rotateRight(h)
	}

	if k.Compare(h.key) == 0 && h.right == nil {
		return nil, true
	}

	if !isRed(h.right) && !isRed(h.right.left) {
		h = moveRedRight(h)
	}

	var ok bool
	if k.Compare(h.key) == 0 {
		h.key = min(h.right).key
		h.val, _ = recurGet(h.right, h.key)
		h.right = deleteMin(h.right)
		ok = true
	} else {
		h.right, ok = delete(h.right, k)
	}

	return balance(h), ok
}

func balance(h *node) *node {
	if isRed(h.right) {
		h = rotateLeft(h)
	}
	if isRed(h.right) && !isRed(h.left) {
		h = rotateLeft(h)
	}
	if isRed(h.left) && isRed(h.left.left) {
		h = rotateRight(h)
	}
	if isRed(h.left) && isRed(h.right) {
		flipColors(h)
	}
	h.n = size(h.left) + size(h.right) + 1
	return h
}
