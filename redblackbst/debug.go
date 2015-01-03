package redblackbst

import (
	"bytes"
	"fmt"
	"io"
	"log"
)

func nodes(x *node, visit func(*node) bool, lo, hi KType) bool {
	if x == nil {
		return true
	}
	cmplo := lo.Compare(x.key)
	cmphi := hi.Compare(x.key)
	if cmplo < 0 {
		if !nodes(x.left, visit, lo, hi) {
			return false
		}
	}
	if cmplo <= 0 && cmphi >= 0 {
		if !visit(x) {
			return false
		}
	}
	if cmphi > 0 {
		if !nodes(x.right, visit, lo, hi) {
			return false
		}
	}
	return true
}

func checkIs23(tree *RedBlack) bool {
	var nodeIs23 func(*node) bool

	visited := map[KType]struct{}{}

	nodeIs23 = func(n *node) bool {

		if n == nil {
			return true
		}

		if _, ok := visited[n.key]; ok {
			log.Panicf("cycle detected, already visited %p", n)
		}
		visited[n.key] = struct{}{}

		if isRed(n) && isRed(n.left) {
			log.Panicf("%q is red, %q also is red", n.key, n.left.key)
		}

		if isRed(n.right) {
			log.Panicf("%q: right link %q is red", n.key, n.right.key)
			return false
		}
		return nodeIs23(n.left) && nodeIs23(n.right)
	}
	if tree.root == nil {
		return true
	}

	// the root itself is red by default for
	// simplicity of the code, it changes nothing
	return nodeIs23(tree.root.left) && nodeIs23(tree.root.right)
}

func dotGraph(h *node, out io.Writer, name string) {
	buf := bytes.NewBuffer(nil)
	fmt.Fprintf(out, "digraph %q {\n", name)
	fmt.Fprintf(buf, "\t%q -> ", name)
	dotvisit(h, out, buf, true)
	buf.WriteTo(out)
	fmt.Fprintf(out, "}\n")
}

func dotvisit(x *node, nodes, edges io.Writer, isLeft bool) {

	var color string
	if isRed(x) {
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
	dotvisit(x.left, nodes, edges, true)
	fmt.Fprintf(edges, "\t\"%p\" -> ", x)
	dotvisit(x.right, nodes, edges, false)
}
