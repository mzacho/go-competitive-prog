package solutions

import (
	"competitive_programming/util/graph"
	"competitive_programming/util/reader"
	"fmt"
)

func Run_codejam2022_4(reader reader.FileReader) {
	noOfTests, _ := reader.NextInt()
	for i := 0; i < noOfTests; i++ {
		l1, _ := reader.NextInt()
		l2, _ := reader.SplitNextIntLine(" ")
		l3, _ := reader.SplitNextIntLine(" ")
		g := buildGraph(l1, l2, l3)
		fmt.Printf("Case #%v: %v\n", i+1, solve(g))
	}
}

func solve(g graph.DFSGraph) int {
	g.ToDotfile("g.dot")
	g.DFS()
	g.ToDotfile("dfs-g.dot")
	if g.TopologicalSort() {
		g.Clear()
		g.DFS()
		for idx, c := range g.DFSForest() {
			c.ToDotfile(fmt.Sprintf("component-%v.dot", idx))
		}
	}
	return 42
}

func buildGraph(n int, funs, edges []int) (res graph.DFSGraph) {
	res = graph.NewDFSGraph(n)
	for from, to := range edges {
		if to != 0 {
			res.AddEdge(to-1, from)
		}
	}
	return
}
