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

import (
	"bytes"
	"fmt"
	"io"
)

// ugly type names to avoid collisions, for easy find/replace.

type KType interface {
	Compare(other KType) int
}

// debugging

// DotGraph exports the sorted set into DOT format.
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
