package redblackbst

import (
	"bytes"
	"fmt"
	"io"
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

func dotGraph(h *node, out io.Writer, name string) {
	buf := bytes.NewBuffer([]byte("\t"))

	fmt.Fprintf(out, "digraph %s {\n", name)
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
	fmt.Fprintf(nodes, "\t\"%p\" [label=\"%p, %v\", shape = circle, color=%s];\n", x, x, x.key, color)

	fmt.Fprintf(edges, "\t\"%p\" -> ", x)
	dotvisit(x.left, nodes, edges, true)
	fmt.Fprintf(edges, "\t\"%p\" -> ", x)
	dotvisit(x.right, nodes, edges, false)
}
