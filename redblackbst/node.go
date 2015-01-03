package redblackbst

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

func newNode(k KType, v VType, n int, color bool) *node {
	return &node{key: k, val: v, n: n, color: color}
}

func isRed(x *node) bool { return (x != nil) && (x.color == red) }

func rotateLeft(h *node) *node {
	x := h.right
	h.right = x.left
	x.left = h
	x.color = h.color
	h.color = red
	x.n = h.n
	h.n = 1 + size(h.left) + size(h.right)
	return x
}

func rotateRight(h *node) *node {
	x := h.left
	h.left = x.right
	x.right = h
	x.color = h.color
	h.color = red
	x.n = h.n
	h.n = 1 + size(h.left) + size(h.right)
	return x
}

func flipColors(h *node) {
	h.color = red
	h.left.color = black
	h.right.color = black
}

func size(x *node) int {
	if x == nil {
		return 0
	}
	return x.n
}
