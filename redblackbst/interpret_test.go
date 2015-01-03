package redblackbst

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

var tests = []struct {
	want graph
	prog P
}{
	{
		prog: P{Name: "simple put", Ops: []Op{
			{"Put", "b"},
			{"Put", "a"},
		}},
		want: makeGraph([]edge{
			{from: "b",
				to:    "a",
				dir:   "left",
				color: "red"},
		}),
	},
	{
		prog: P{Name: "simple put 2", Ops: []Op{
			{"Put", "a"},
			{"Put", "b"},
		}},
		want: makeGraph([]edge{
			{from: "b",
				to:    "a",
				dir:   "left",
				color: "red"},
		}),
	},
	{
		prog: P{Name: "two blacks", Ops: []Op{
			{"Put", "a"},
			{"Put", "e"},
			{"Put", "s"},
		}},
		want: makeGraph([]edge{
			{from: "e",
				to:    "a",
				dir:   "left",
				color: "black"},
			{from: "e",
				to:    "s",
				dir:   "right",
				color: "black"},
		}),
	},

	{
		prog: P{Name: "two blacks, one red", Ops: []Op{
			{"Put", "a"},
			{"Put", "e"},
			{"Put", "s"},
			{"Put", "r"},
		}},
		want: makeGraph([]edge{
			{from: "e",
				to:    "a",
				dir:   "left",
				color: "black"},
			{from: "e",
				to:    "s",
				dir:   "right",
				color: "black"},
			{from: "s",
				to:    "r",
				dir:   "left",
				color: "red"},
		}),
	},

	{
		prog: P{Name: "two blacks, two red", Ops: []Op{
			{"Put", "a"},
			{"Put", "e"},
			{"Put", "s"},
			{"Put", "r"},
			{"Put", "c"},
		}},
		want: makeGraph([]edge{
			{from: "e",
				to:    "c",
				dir:   "left",
				color: "black"},
			{from: "e",
				to:    "s",
				dir:   "right",
				color: "black"},
			{from: "s",
				to:    "r",
				dir:   "left",
				color: "red"},
			{from: "c",
				to:    "a",
				dir:   "left",
				color: "red"},
		}),
	},
	{
		prog: P{Name: "two blacks, three red", Ops: []Op{
			{"Put", "a"},
			{"Put", "e"},
			{"Put", "s"},
			{"Put", "r"},
			{"Put", "c"},
			{"Put", "h"},
		}},
		want: makeGraph([]edge{
			{from: "r",
				to:    "e",
				dir:   "left",
				color: "red"},
			{from: "r",
				to:    "s",
				dir:   "right",
				color: "black"},
			{from: "e",
				to:    "c",
				dir:   "left",
				color: "black"},
			{from: "e",
				to:    "h",
				dir:   "right",
				color: "black"},
			{from: "c",
				to:    "a",
				dir:   "left",
				color: "red"},
		}),
	},

	{
		prog: P{Name: "standard indexing client", Ops: []Op{
			{"Put", "s"},
			{"Put", "e"},
			{"Put", "a"},
			{"Put", "r"},
			{"Put", "c"},
			{"Put", "h"},
			{"Put", "x"},
			{"Put", "m"},
			{"Put", "p"},
			{"Put", "l"},
		}},
		want: makeGraph([]edge{
			{from: "m",
				to:    "e",
				dir:   "left",
				color: "black"},
			{from: "m",
				to:    "r",
				dir:   "right",
				color: "black"},
			{from: "e",
				to:    "c",
				dir:   "left",
				color: "black"},
			{from: "e",
				to:    "l",
				dir:   "right",
				color: "black"},
			{from: "l",
				to:    "h",
				dir:   "left",
				color: "red"},
			{from: "c",
				to:    "a",
				dir:   "left",
				color: "red"},
			{from: "r",
				to:    "p",
				dir:   "left",
				color: "black"},
			{from: "r",
				to:    "x",
				dir:   "right",
				color: "black"},
			{from: "x",
				to:    "s",
				dir:   "left",
				color: "red"},
		}),
	},

	{
		prog: P{Name: "same keys, in increasing order", Ops: []Op{
			{"Put", "a"},
			{"Put", "c"},
			{"Put", "e"},
			{"Put", "h"},
			{"Put", "l"},
			{"Put", "m"},
			{"Put", "p"},
			{"Put", "r"},
			{"Put", "s"},
			{"Put", "x"},
		}},
		want: makeGraph([]edge{
			{from: "h",
				to:    "c",
				dir:   "left",
				color: "black"},
			{from: "h",
				to:    "r",
				dir:   "right",
				color: "black"},
			{from: "r",
				to:    "m",
				dir:   "left",
				color: "red"},
			{from: "r",
				to:    "x",
				dir:   "right",
				color: "black"},
			{from: "m",
				to:    "l",
				dir:   "left",
				color: "black"},
			{from: "m",
				to:    "p",
				dir:   "right",
				color: "black"},
			{from: "x",
				to:    "s",
				dir:   "left",
				color: "red"},
			{from: "c",
				to:    "a",
				dir:   "left",
				color: "black"},
			{from: "c",
				to:    "e",
				dir:   "right",
				color: "black"},
		}),
	},
}

type Op struct {
	Op  string
	Key string
}

type P struct {
	Name string
	Ops  []Op
}

func (p *P) Run(tree *RedBlack) {
	for step, op := range p.Ops {
		switch op.Op {
		case "Put":
			tree.Put(K(op.Key), struct{}{})
		case "Delete":
			tree.Delete(K(op.Key))
		case "DeleteMin":
			tree.DeleteMin()
		case "Print":
			printTreeStats(tree, fmt.Sprintf("%s-%d", p.Name, step))
		}
	}
}

func TestScenarios(t *testing.T) {

	for _, tt := range tests {
		tree := New()
		tt.prog.Run(tree)

		want := tt.want
		got := toGraph(tree)

		if !reflect.DeepEqual(want.nodes, got.nodes) {
			openDot(want.toDot())
			openDot(got.toDot())
			t.Logf("want=%#v", want.nodes)
			t.Logf("got =%#v", got.nodes)
			t.Fatalf("%q: failed,", tt.prog.Name)
		}
	}
}

type nd struct {
	name  string
	edges []edge
}

type edge struct {
	from  string
	to    string
	dir   string
	color string
}

type graph struct {
	name  string
	nodes map[string]nd
}

func toGraph(tree *RedBlack) graph {
	g := graph{
		name:  "got",
		nodes: make(map[string]nd),
	}

	lo, _, _ := tree.Min()
	hi, _, _ := tree.Max()
	nodes(tree.root, func(n *node) bool {

		h := nd{
			name: string(n.key.(K)),
		}

		if n.left != nil {
			var leftColor string
			if isRed(n.left) {
				leftColor = "red"
			} else {
				leftColor = "black"
			}

			h.edges = append(h.edges, edge{
				from:  h.name,
				to:    string(n.left.key.(K)),
				color: leftColor,
				dir:   "left",
			})
		}
		if n.right != nil {
			var rightColor string
			if isRed(n.right) {
				rightColor = "red"
			} else {
				rightColor = "black"
			}
			h.edges = append(h.edges, edge{
				from:  h.name,
				to:    string(n.right.key.(K)),
				color: rightColor,
				dir:   "right",
			})
		}

		g.nodes[h.name] = h
		return true
	}, lo, hi)

	return g
}

func makeGraph(edges []edge) graph {
	g := graph{
		name:  "want",
		nodes: make(map[string]nd),
	}
	for _, e := range edges {
		n, ok := g.nodes[e.from]
		if !ok {
			n = nd{
				name:  e.from,
				edges: []edge{e},
			}
		} else {
			n.edges = append(n.edges, e)
		}
		g.nodes[e.from] = n
		n, ok = g.nodes[e.to]
		if !ok {
			n = nd{
				name: e.to,
			}
			g.nodes[e.to] = n
		}

	}
	return g
}

func (g graph) toDot() *bytes.Buffer {
	w := bytes.NewBuffer(nil)

	fmt.Fprintf(w, "digraph %q {", g.name)

	for _, node := range g.nodes {
		fmt.Fprintf(w, "\t%q [shape = circle];\n", node.name)
	}

	for _, node := range g.nodes {
		for _, edge := range node.edges {
			fmt.Fprintf(w, "\t%q -> %q [label=%q, color=%s];\n", node.name, edge.to, edge.dir, edge.color)
		}
	}

	fmt.Fprintf(w, "}\n")

	return w
}
