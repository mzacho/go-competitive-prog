package graph

type Timestamp int
type VertexColor int

const (
	VertexColorWhite VertexColor = iota + 1
	VertexColorGray
	VertexColorBlack
)

// Reverses all edges
// Removes BFS/ DFS data
// func (g *Graph[T]) Transpose() Graph[T] {
// 	var edges = make([][]int, len(g.E))
// 	for from, adjList := range g.E {
// 		for _, to := range adjList {
// 			edges[to] = append(edges[to], from)
// 		}
// 	}
// 	var V = make([]*Vertex[T], len(g.V))
// 	for idx, n := range g.V {
// 		V[idx] = &Vertex[T]{
// 			Value: n.Value,
// 			Id:    n.Id,
// 		}
// 	}
// 	return Graph[T]{
// 		Root:      g.Root,
// 		E:         edges,
// 		V:         V,
// 		NoOfEdges: g.NoOfEdges,
// 	}
// }
