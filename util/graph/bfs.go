package graph

import (
	"bytes"
	"competitive_programming/util/queue"
	"fmt"
	"os"

	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
)

type BFSGraph struct {
	Root      int
	NoOfEdges int
	E         [][]int
	V         []BFSVertex
}

func (g *BFSGraph) AddEdge(from, to int) {
	g.E[from] = append(g.E[from], to)
}

func (g *BFSGraph) AppendNode(n BFSVertex) {
	g.V = append(g.V, n)
}

func NewBFSGraph[T any](noOfNodes int) *BFSGraph {
	g := &BFSGraph{
		E: make([][]int, noOfNodes),
		V: make([]BFSVertex, noOfNodes),
	}
	for v := range g.V {
		g.V[v].Id = v
		g.V[v].Color = VertexColorWhite
		g.V[v].Discovered = -1
	}
	return g
}

type BFSVertex struct {
	Id         int
	Color      VertexColor
	Pred       int
	Discovered Timestamp
}

// Precondition: For all nodes n: n.color == White, n.pred == nil and n.discovered = -1
func (g *BFSGraph) BFS(sources ...int) {
	q := queue.NewQueue[int](len(g.V))
	for _, s := range sources {
		g.V[s].Color = VertexColorGray
		g.V[s].Discovered = 0
		q.Enqueue(&s)
	}

	for v := q.Dequeue(); v != nil; v = q.Dequeue() {
		for _, u := range g.E[*v] {
			if g.V[u].Color == VertexColorWhite {
				g.V[u].Color = VertexColorGray
				g.V[u].Discovered = g.V[*v].Discovered + 1
				g.V[u].Pred = *v
				q.Enqueue(&u)
			}
		}
		g.V[*v].Color = VertexColorBlack
	}
}

// ----------------- visualisations

func (g *BFSGraph) ToDotfile(name string) {
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
	var nodes = make(map[int]*cgraph.Node)
	for id, n := range g.V {
		if id != n.Id {
			panic("precondition violated")
		}
		node, err := graph.CreateNode(fmt.Sprintf("%v", n.Id))
		if err != nil {
			panic(err)
		}
		node.SetLabel(fmt.Sprintf("id: %v, value: %v\ndiscovered: %v", n.Id, nil, n.Discovered))
		switch n.Color {
		case VertexColorWhite:
			node.SetFillColor("white")
		case VertexColorGray:
			node.SetFillColor("gray")
		case VertexColorBlack:
			node.SetFillColor("black")
		}
		nodes[id] = node
	}
	for from, adjList := range g.E {
		for _, to := range adjList {
			graph.CreateEdge("", nodes[from], nodes[to])
		}
	}
	var buf bytes.Buffer
	if err := gv.Render(graph, "dot", &buf); err != nil {
		panic(err)
	}
	if err := os.WriteFile(name, buf.Bytes(), 0666); err != nil {
		panic(err)
	}
}
