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

// Clear all the values in the tree.
func (r *RedBlack) Clear() { r.root = nil }

// Put a value in the tree at key `k`. The old value at `k` is returned
// if the key was already present.
func (r *RedBlack) Put(k KType, v VType) (old VType, overwrite bool) {
	r.root, old, overwrite = put(r.root, k, v)
	return
}

func put(h *node, k KType, v VType) (_ *node, old VType, overwrite bool) {
	if h == nil {
		return newNode(k, v, 1, red), old, overwrite
	}

	cmp := k.Compare(h.key)
	if cmp < 0 {
		h.left, old, overwrite = put(h.left, k, v)
	} else if cmp > 0 {
		h.right, old, overwrite = put(h.right, k, v)
	} else {
		overwrite = true
		old = h.val
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
	return h, old, overwrite
}

// Get a value from the tree at key `k`. Returns false
// if the key doesn't exist.
func (r RedBlack) Get(k KType) (VType, bool) {
	return loopGet(r.root, k)
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

// Has tells if a value exists at key `k`.
func (r RedBlack) Has(k KType) bool {
	_, ok := loopGet(r.root, k)
	return ok
}

// Min returns the smallest key/value in the tree, if it exists.
func (r RedBlack) Min() (KType, VType, bool) {
	if r.root == nil {
		return nil, nil, false
	}
	h := min(r.root)
	return h.key, h.val, true
}

func min(x *node) *node {
	if x.left == nil {
		return x
	}
	return min(x.left)
}

// Max returns the largest key/value in the tree, if it exists.
func (r RedBlack) Max() (KType, VType, bool) {
	if r.root == nil {
		return nil, nil, false
	}
	h := max(r.root)
	return h.key, h.val, true
}

func max(x *node) *node {
	if x.right == nil {
		return x
	}
	return max(x.right)
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

// Ceiling returns the smallest key/value in the tree that is larger than
// `k`.
func (r RedBlack) Ceiling(k KType) (KType, VType, bool) {
	x := ceiling(r.root, k)
	if x == nil {
		return nil, nil, false
	}
	return x.key, x.val, true
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

// Select key of rank k, meaning the k-th biggest KType in the tree.
func (r RedBlack) Select(k int) (KType, VType, bool) {
	x := nodeselect(r.root, k)
	if x == nil {
		return nil, nil, false
	}
	return x.key, x.val, true
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

// Rank is the number of keys less than `k`.
func (r RedBlack) Rank(k KType) int {
	return keyrank(k, r.root)
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

// Keys visit each keys in the tree, in order.
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

// RangedKeys visit each keys between lo and hi in the tree, in order.
// It stops when visit returns false.
func (r RedBlack) RangedKeys(lo, hi KType, visit func(KType, VType) bool) {
	keys(r.root, visit, lo, hi)
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

// DeleteMin removes the smallest key and its value from the tree.
func (r *RedBlack) DeleteMin() (oldk KType, oldv VType, ok bool) {
	r.root, oldk, oldv, ok = deleteMin(r.root)
	if !r.IsEmpty() {
		r.root.color = black
	}
	return
}

func deleteMin(h *node) (_ *node, oldk KType, oldv VType, ok bool) {
	if h == nil {
		return nil, nil, nil, false
	}

	if h.left == nil {
		return nil, h.key, h.val, true
	}
	if !isRed(h.left) && !isRed(h.left.left) {
		h = moveRedLeft(h)
	}
	h.left, oldk, oldv, ok = deleteMin(h.left)
	return balance(h), oldk, oldv, ok
}

// DeleteMax removes the largest key and its value from the tree.
func (r *RedBlack) DeleteMax() (oldk KType, oldv VType, ok bool) {
	r.root, oldk, oldv, ok = deleteMax(r.root)
	if !r.IsEmpty() {
		r.root.color = black
	}
	return
}

func deleteMax(h *node) (_ *node, oldk KType, oldv VType, ok bool) {
	if h == nil {
		return nil, oldk, oldv, ok
	}
	if isRed(h.left) {
		h = rotateRight(h)
	}
	if h.right == nil {
		return nil, h.key, h.val, true
	}
	if !isRed(h.right) && !isRed(h.right.left) {
		h = moveRedRight(h)
	}
	h.right, oldk, oldv, ok = deleteMax(h.right)
	return balance(h), oldk, oldv, ok
}

// Delete key `k` from tree, if it exists.
func (r *RedBlack) Delete(k KType) (old VType, ok bool) {
	if r.root == nil {
		return
	}
	r.root, old, ok = delete(r.root, k)
	if !r.IsEmpty() {
		r.root.color = black
	}
	return
}

func delete(h *node, k KType) (*node, VType, bool) {

	var old VType
	var ok bool

	if h == nil {
		return nil, nil, false
	}

	if k.Compare(h.key) < 0 {
		if h.left == nil {
			return h, nil, false
		}

		if !isRed(h.left) && !isRed(h.left.left) {
			h = moveRedLeft(h)
		}

		h.left, old, ok = delete(h.left, k)
		h = balance(h)
		return h, old, ok
	} else {
		if isRed(h.left) {
			h = rotateRight(h)
		}

		if k.Compare(h.key) == 0 && h.right == nil {
			return nil, h.val, true
		}

		if h.right != nil && !isRed(h.right) && !isRed(h.right.left) {
			h = moveRedRight(h)
		}

		if k.Compare(h.key) == 0 {

			var subk KType
			var subv VType
			h.right, subk, subv, ok = deleteMin(h.right)

			old, h.key, h.val = h.val, subk, subv
			ok = true
		} else {
			h.right, old, ok = delete(h.right, k)
		}
	}
	h = balance(h)
	return h, old, ok
}

// deletions

func moveRedLeft(h *node) *node {
	flipColors(h)
	if isRed(h.right.left) {
		h.right = rotateRight(h.right)
		h = rotateLeft(h)
		flipColors(h)
	}
	return h
}

func moveRedRight(h *node) *node {
	flipColors(h)
	if isRed(h.left.left) {
		h = rotateRight(h)
		flipColors(h)
	}
	return h
}

func balance(h *node) *node {
	if isRed(h.right) {
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
