package graph

import (
	"bytes"
	"fmt"
	"os"

	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
)

type BinaryTree[T any] struct {
	V     T
	Left  *BinaryTree[T]
	Right *BinaryTree[T]
}

// type TreeVisitor[T any] interface {
// 	Apply(*BinaryTree[T])
// }

// func Visit[T any](t *BinaryTree[T], transfer func(*BinaryTree[T])) {
// 	if t == nil {
// 		return
// 	}
// 	transfer(t)
// 	if t.Left != nil {
// 		Visit(t.Left, transfer)
// 	}
// 	if t.Right != nil {
// 		Visit(t.Right, transfer)
// 	}
// }

/// -------------------------------------------------------------------

// type PrintTreeVisitor struct {
// 	id     int
// 	graph  *cgraph.Graph
// 	parent *cgraph.Node
// }

// func (v *PrintTreeVisitor) Apply(t *BinaryTree[T]) {
// 	node, err := v.graph.CreateNode(fmt.Sprintf("%v", v.id))
// 	if err != nil {
// 		panic(err)
// 	}
// 	node.SetLabel(fmt.Sprintf("id: %v, value: %v\n", v.id, t.V))
// 	if v.parent != nil {
// 		v.graph.CreateEdge("", v.parent, node)
// 	}
// 	v.id += 1
// }

func (t *BinaryTree[T]) ToDotfile(name string) {
	gv := graphviz.New()
	graph, err := gv.Graph()
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := graph.Close(); err != nil {
			panic(err)
		}
		gv.Close()
	}()
	var visit func(t *BinaryTree[T], id int, parent *cgraph.Node) int
	visit = func(t *BinaryTree[T], id int, parent *cgraph.Node) int {
		node, err := graph.CreateNode(fmt.Sprintf("%v", id))
		if err != nil {
			panic(err)
		}
		node.SetLabel(fmt.Sprintf("id: %v, value: %v\n", id, t.V))
		if parent != nil {
			graph.CreateEdge("", parent, node)
		}
		if t.Left != nil {
			id = visit(t.Left, id+1, node)
			id = visit(t.Right, id+1, node)
			return id
		}
		return id
	}
	visit(t, 0, nil)

	var buf bytes.Buffer
	if err := gv.Render(graph, "dot", &buf); err != nil {
		panic(err)
	}
	if err := os.WriteFile(name, buf.Bytes(), 0666); err != nil {
		panic(err)
	}
}
