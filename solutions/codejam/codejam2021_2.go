package solutions

import (
	"competitive_programming/util/convert"
	"competitive_programming/util/graph"
	"competitive_programming/util/reader"
	"fmt"
)

func Run_codejam2021_2(reader reader.FileReader) {
	t, _ := reader.NextInt()
	for i := 0; i < t; i++ {
		l, _ := reader.NextLineSplit(" ")
		l1 := convert.UnsafeToIntBase10(l[0])
		l2 := convert.UnsafeToIntBase10(l[1])
		l3 := l[2]
		fmt.Printf("Case #%v: %v\n", i+1, cost(build_bintree(l1, l2, l3)))
	}
}

type Data struct {
	Choice string
	Cost   int
}

func build_bintree(X int, Y int, l3 string) graph.BinaryTree[Data] {
	var visit func(*graph.BinaryTree[Data], Data, Data)
	visit = func(t *graph.BinaryTree[Data], left, right Data) {
		if t.Left == nil {
			t.Left = &graph.BinaryTree[Data]{
				V: left,
			}
			t.Right = &graph.BinaryTree[Data]{
				V: right,
			}
		} else {
			visit(t.Left, left, right)
			visit(t.Right, left, right)
		}
	}
	root := graph.BinaryTree[Data]{}
	for _, char := range l3 {
		if char == '?' {
			left := Data{
				Choice: "C",
			}
			right := Data{
				Choice: "J",
			}
			visit(&root, left, right)
		}
	}
	return root
}

func cost(tree graph.BinaryTree[Data]) int {
	// var visit func (*graph.BinaryTree[Data]) int
	// visit = func (t *graph.BinaryTree[Data]) int {
	// 	if t.Left == nil {
	// 		return 0
	// 	}

	// }

	//tree.ToDotfile("bintree.dot")
	return 42
}
