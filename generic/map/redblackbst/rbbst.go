package redblackbst

// Generic port of github.com/aybabtme/datagen/map/redblackbst.

func (r RedBlack[K, V]) compare(a, b K) int { return r.cmp(a, b) }

// RedBlack is a sorted map built on a left leaning red black balanced
// search sorted map. It stores V values, keyed by K.
type RedBlack[K, V any] struct {
	root *mapnode[K, V]
	cmp  func(a, b K) int
}

// NewRedBlack creates a sorted map, ordering keys with `cmp`.
func NewRedBlack[K, V any](cmp func(a, b K) int) *RedBlack[K, V] {
	return &RedBlack[K, V]{cmp: cmp}
}

// IsEmpty tells if the sorted map contains no key/value.
func (r RedBlack[K, V]) IsEmpty() bool {
	return r.root == nil
}

// Size of the sorted map.
func (r RedBlack[K, V]) Size() int { return r.root.size() }

// Clear all the values in the sorted map.
func (r *RedBlack[K, V]) Clear() { r.root = nil }

// Put a value in the sorted map at key `k`. The old value at `k` is returned
// if the key was already present.
func (r *RedBlack[K, V]) Put(k K, v V) (old V, overwrite bool) {
	r.root, old, overwrite = r.put(r.root, k, func() V { return v }, func(_ V) V { return v })
	return
}

// Mutate is like a Put when `k` isn't defined, but allows you to create or mutate the value found at the location of `k`.
func (r *RedBlack[K, V]) Mutate(k K, creator func() V, mutator func(old V) V) {
	r.root, _, _ = r.put(r.root, k, creator, mutator)
}

func (r *RedBlack[K, V]) put(h *mapnode[K, V], k K, create func() V, mutate func(old V) V) (_ *mapnode[K, V], old V, overwrite bool) {
	if h == nil {
		n := &mapnode[K, V]{key: k, val: create(), n: 1, colorRed: true}
		return n, old, overwrite
	}

	cmp := r.compare(k, h.key)
	if cmp < 0 {
		h.left, old, overwrite = r.put(h.left, k, create, mutate)
	} else if cmp > 0 {
		h.right, old, overwrite = r.put(h.right, k, create, mutate)
	} else {
		overwrite = true
		old = h.val
		h.val = mutate(old)
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
func (r RedBlack[K, V]) Get(k K) (V, bool) {
	return r.loopGet(r.root, k)
}

func (r RedBlack[K, V]) loopGet(h *mapnode[K, V], k K) (v V, ok bool) {
	for h != nil {
		cmp := r.compare(k, h.key)
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
func (r RedBlack[K, V]) Has(k K) bool {
	_, ok := r.loopGet(r.root, k)
	return ok
}

// Min returns the smallest key/value in the sorted map, if it exists.
func (r RedBlack[K, V]) Min() (k K, v V, ok bool) {
	if r.root == nil {
		return
	}
	h := r.min(r.root)
	return h.key, h.val, true
}

func (r RedBlack[K, V]) min(x *mapnode[K, V]) *mapnode[K, V] {
	if x.left == nil {
		return x
	}
	return r.min(x.left)
}

// Max returns the largest key/value in the sorted map, if it exists.
func (r RedBlack[K, V]) Max() (k K, v V, ok bool) {
	if r.root == nil {
		return
	}
	h := r.max(r.root)
	return h.key, h.val, true
}

func (r RedBlack[K, V]) max(x *mapnode[K, V]) *mapnode[K, V] {
	if x.right == nil {
		return x
	}
	return r.max(x.right)
}

// Floor returns the largest key/value in the sorted map that is smaller than
// `k`.
func (r RedBlack[K, V]) Floor(key K) (k K, v V, ok bool) {
	x := r.floor(r.root, key)
	if x == nil {
		return
	}
	return x.key, x.val, true
}

func (r RedBlack[K, V]) floor(h *mapnode[K, V], k K) *mapnode[K, V] {
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

// Ceiling returns the smallest key/value in the sorted map that is larger than
// `k`.
func (r RedBlack[K, V]) Ceiling(key K) (k K, v V, ok bool) {
	x := r.ceiling(r.root, key)
	if x == nil {
		return
	}
	return x.key, x.val, true
}

func (r RedBlack[K, V]) ceiling(h *mapnode[K, V], k K) *mapnode[K, V] {
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

// Select key of rank k, meaning the k-th biggest K in the sorted map.
func (r RedBlack[K, V]) Select(key int) (k K, v V, ok bool) {
	x := r.nodeselect(r.root, key)
	if x == nil {
		return
	}
	return x.key, x.val, true
}

func (r RedBlack[K, V]) nodeselect(x *mapnode[K, V], k int) *mapnode[K, V] {
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
func (r RedBlack[K, V]) Rank(k K) int {
	return r.keyrank(k, r.root)
}

func (r RedBlack[K, V]) keyrank(k K, h *mapnode[K, V]) int {
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

// Keys visit each keys in the sorted map, in order.
// It stops when visit returns false.
func (r RedBlack[K, V]) Keys(visit func(K, V) bool) {
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
func (r RedBlack[K, V]) RangedKeys(lo, hi K, visit func(K, V) bool) {
	r.keys(r.root, visit, lo, hi)
}

func (r RedBlack[K, V]) keys(h *mapnode[K, V], visit func(K, V) bool, lo, hi K) bool {
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
func (r *RedBlack[K, V]) DeleteMin() (oldk K, oldv V, ok bool) {
	r.root, oldk, oldv, ok = r.deleteMin(r.root)
	if !r.IsEmpty() {
		r.root.colorRed = false
	}
	return
}

func (r *RedBlack[K, V]) deleteMin(h *mapnode[K, V]) (_ *mapnode[K, V], oldk K, oldv V, ok bool) {
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
func (r *RedBlack[K, V]) DeleteMax() (oldk K, oldv V, ok bool) {
	r.root, oldk, oldv, ok = r.deleteMax(r.root)
	if !r.IsEmpty() {
		r.root.colorRed = false
	}
	return
}

func (r *RedBlack[K, V]) deleteMax(h *mapnode[K, V]) (_ *mapnode[K, V], oldk K, oldv V, ok bool) {
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
func (r *RedBlack[K, V]) Delete(k K) (old V, ok bool) {
	if r.root == nil {
		return
	}
	r.root, old, ok = r.delete(r.root, k)
	if !r.IsEmpty() {
		r.root.colorRed = false
	}
	return
}

func (r *RedBlack[K, V]) delete(h *mapnode[K, V], k K) (_ *mapnode[K, V], old V, ok bool) {

	if h == nil {
		return h, old, false
	}

	if r.compare(k, h.key) < 0 {
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

	if r.compare(k, h.key) == 0 && h.right == nil {
		return nil, h.val, true
	}

	if h.right != nil && !h.right.isRed() && !h.right.left.isRed() {
		h = r.moveRedRight(h)
	}

	if r.compare(k, h.key) == 0 {

		var subk K
		var subv V
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

func (r *RedBlack[K, V]) moveRedLeft(h *mapnode[K, V]) *mapnode[K, V] {
	r.flipColors(h)
	if h.right.left.isRed() {
		h.right = r.rotateRight(h.right)
		h = r.rotateLeft(h)
		r.flipColors(h)
	}
	return h
}

func (r *RedBlack[K, V]) moveRedRight(h *mapnode[K, V]) *mapnode[K, V] {
	r.flipColors(h)
	if h.left.left.isRed() {
		h = r.rotateRight(h)
		r.flipColors(h)
	}
	return h
}

func (r *RedBlack[K, V]) balance(h *mapnode[K, V]) *mapnode[K, V] {
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

func (r *RedBlack[K, V]) rotateLeft(h *mapnode[K, V]) *mapnode[K, V] {
	x := h.right
	h.right = x.left
	x.left = h
	x.colorRed = h.colorRed
	h.colorRed = true
	x.n = h.n
	h.n = 1 + h.left.size() + h.right.size()
	return x
}

func (r *RedBlack[K, V]) rotateRight(h *mapnode[K, V]) *mapnode[K, V] {
	x := h.left
	h.left = x.right
	x.right = h
	x.colorRed = h.colorRed
	h.colorRed = true
	x.n = h.n
	h.n = 1 + h.left.size() + h.right.size()
	return x
}

func (r *RedBlack[K, V]) flipColors(h *mapnode[K, V]) {
	h.colorRed = !h.colorRed
	h.left.colorRed = !h.left.colorRed
	h.right.colorRed = !h.right.colorRed
}

// nodes

type mapnode[K, V any] struct {
	key         K
	val         V
	left, right *mapnode[K, V]
	n           int
	colorRed    bool
}

func (x *mapnode[K, V]) isRed() bool { return (x != nil) && (x.colorRed == true) }

func (x *mapnode[K, V]) size() int {
	if x == nil {
		return 0
	}
	return x.n
}
