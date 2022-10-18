package graph

import (
	"bytes"
	"fmt"
	"os"
	"sort"

	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
)

type DFSGraph struct {
	NoOfEdges int
	E         [][]int
	V         []DFSVertex
	ETypes    [][]EdgeType
	CCs       [][]int // TODO - use pred subgraph somehow
}

func (g *DFSGraph) AddEdge(from, to int) {
	g.E[from] = append(g.E[from], to)
	g.ETypes[from] = append(g.ETypes[from], EdgeTypeUnset)
	g.NoOfEdges += 1
}

func NewDFSGraph(noOfNodes int) DFSGraph {
	g := DFSGraph{
		E:      make([][]int, noOfNodes),
		V:      make([]DFSVertex, noOfNodes),
		ETypes: make([][]EdgeType, noOfNodes),
	}
	for i := 0; i < noOfNodes; i++ {
		g.V[i].Id = i
		g.V[i].Color = VertexColorWhite
	}
	return g
}

type DFSVertex struct {
	Id         int
	Color      VertexColor
	Pred       *int
	Discovered Timestamp
	Ended      Timestamp
}

type EdgeType int

const (
	EdgeTypeUnset EdgeType = iota + 1
	EdgeTypeTree
	EdgeTypeBack
	EdgeTypeForward
	EdgeTypeCross
)

// Precondition: For all nodes n, n.color == White and n.pred == nil
func (g *DFSGraph) DFS() {
	t := 0
	var visit func(v int)
	visit = func(v int) {
		t += 1
		g.V[v].Discovered = Timestamp(t)
		g.V[v].Color = VertexColorGray
		for idx, u := range g.E[v] {
			if g.V[u].Color == VertexColorWhite {
				g.ETypes[v][idx] = EdgeTypeTree
				g.V[u].Pred = &v
				g.AddVertexToLastCC(u)
				visit(u)
			} else if g.V[u].Color == VertexColorGray {
				g.ETypes[v][idx] = EdgeTypeBack
			} else if g.V[u].Color == VertexColorBlack {
				g.ETypes[v][idx] = EdgeTypeForward
				// or cross?
			}
		}
		g.V[v].Color = VertexColorBlack
		t += 1
		g.V[v].Ended = Timestamp(t)
	}
	for v := range g.V {
		if g.V[v].Color == VertexColorWhite {
			g.CCs = append(g.CCs, []int{v})
			visit(v)
		}
	}
}

func (g DFSGraph) PredecessorSubgraph() (res DFSGraph) {
	res = NewDFSGraph(len(g.V))
	res.V = g.V
	for idx, v := range g.V {
		if v.Pred != nil {
			res.E[*v.Pred] = append(res.E[*v.Pred], idx)
			res.ETypes[*v.Pred] = append(res.ETypes[*v.Pred], EdgeTypeUnset)
		}
	}
	return
}

func (g DFSGraph) ContainsCycles() bool {
	for v, edges := range g.E {
		for idx := range edges {
			if g.ETypes[v][idx] == EdgeTypeBack {
				return true
			}
		}
	}
	return false
}

func (g *DFSGraph) Clear() {
	for i := 0; i < len(g.V); i++ {
		g.V[i].Color = VertexColorWhite
		g.V[i].Discovered = -1
		g.V[i].Ended = -1
		g.V[i].Pred = nil
	}
	for v, edges := range g.E {
		for idx := range edges {
			g.ETypes[v][idx] = EdgeTypeUnset
		}
	}
	g.CCs = nil
}

func (g DFSGraph) TopologicalSort() (success bool) {
	if g.ContainsCycles() {
		return false
	}
	sort.Sort(g)
	return true
}

func (g DFSGraph) Len() int           { return len(g.V) }
func (g DFSGraph) Less(i, j int) bool { return g.V[i].Ended > g.V[j].Ended }
func (g DFSGraph) Swap(i, j int) {
	g.V[i], g.V[j] = g.V[j], g.V[i]
	g.E[i], g.E[j] = g.E[j], g.E[i]
	g.ETypes[i], g.ETypes[j] = g.ETypes[j], g.ETypes[i]

	for v, edges := range g.E {
		for idx, u := range edges {
			if u == i {
				g.E[v][idx] = j
			} else if u == j {
				g.E[v][idx] = i
			}
		}
	}
}

func (g DFSGraph) AddVertexToLastCC(v int) {
	g.CCs[len(g.CCs)-1] = append(g.CCs[len(g.CCs)-1], v)
}

func (g DFSGraph) CCToGraph(root int) (res DFSGraph) {
	res = NewDFSGraph(len(g.CCs[root]))
	var inv = make(map[int]int)
	for idx, v := range g.CCs[root] {
		res.V[idx] = g.V[v]
		inv[v] = idx
	}
	for idx, v := range g.CCs[root] {
		for _, to := range g.E[v] {
			res.AddEdge(idx, inv[to])
		}
	}
	return
}

func (g DFSGraph) DFSForest() (res []DFSGraph) {
	for i := 0; i < len(g.CCs); i++ {
		res = append(res, g.CCToGraph(i))
	}
	return
}

// ----------------- visualisations

func (g DFSGraph) ToDotfile(name string) {
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
		// if id != n.Id {
		// 	panic("precondition violated")
		// }
		node, err := graph.CreateNode(fmt.Sprintf("%v", n.Id))
		if err != nil {
			panic(err)
		}
		node.SetLabel(fmt.Sprintf("id: %v, value: %v\ntimestamp: %v/%v", n.Id, nil, n.Discovered, n.Ended))
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
		for idx, to := range adjList {
			e, _ := graph.CreateEdge("", nodes[from], nodes[to])
			if g.ETypes[from] != nil {
				t := g.ETypes[from][idx].String()
				e.SetLabel(t)
			}

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

func (t EdgeType) String() string {
	switch t {
	case EdgeTypeUnset:
		return "U"
	case EdgeTypeTree:
		return "T"
	case EdgeTypeBack:
		return "B"
	case EdgeTypeForward:
		return "F"
	case EdgeTypeCross:
		return "C"
	}
	panic("")
}
