package solutions

import (
	"competitive_programming/util/array"
	"competitive_programming/util/reader"
	"fmt"
)

func Run_codejam2020_1(reader reader.FileReader) {
	t, _ := reader.NextInt()
	for i := 0; i < t; i++ {
		l1, _ := reader.NextInt()
		m := make([][]int, l1)
		for j := 0; j < l1; j++ {
			l, _ := reader.SplitNextIntLine(" ")
			m[j] = l
		}
		fmt.Printf("Case #%v: %v\n", i+1, NaiveSolution(l1, m))
		// fmt.Printf("Case #%v:\n", i+1)
		// fmt.Println(MY_SOLUTIONZ(l1, l2, l3))
	}
}

func NaiveSolution(l1 int, matrix [][]int) string {
	M := array.Matrix[int](matrix)
	n, m := M.Dim()
	var repeatedRows []int
	for i := 0; i < n; i++ {
		var row []int
		for j := 0; j < m-1; j++ {
			row = append(row, M[i][j])
			if array.IndexOf(M[i][j+1], row) != -1 {
				repeatedRows = append(repeatedRows, i+1)
				break
			}
		}
	}
	M.Transpose()
	var repeatedColumns []int
	for i := 0; i < m; i++ {
		var row []int
		for j := 0; j < n-1; j++ {
			row = append(row, M[i][j])
			if array.IndexOf(M[i][j+1], row) != -1 {
				repeatedColumns = append(repeatedColumns, i+1)
				break
			}
		}
	}
	return fmt.Sprintf("%v %v", len(repeatedRows), len(repeatedColumns))
}

func Trace(l1 int, m [][]int) int {
	var plus = func(x, y int) int {
		return x + y
	}
	return array.Trace(array.Matrix[int](m), plus)
}
