package day8

import (
	"competitive_programming/util/convert"
	"competitive_programming/util/functional"
	"competitive_programming/util/reader"
	"competitive_programming/util/types"
	"fmt"
	"math"
	"strings"
)

type line []int

// type line []int
// type line [][]string

func parseLine(s string) line {
	// return line(s)
	return functional.Map(
		strings.Split(s, " "),
		convert.UnsafeToIntBase10,
	)
	// return functional.Map(
	// 	strings.Split(s, " -> "),
	// 	functional.StrSplit(","),
	// )
}

type testcase struct {
	lines []line
}

func parseTestCase(lines []line) testcase {
	return testcase{lines}
}

type firstLine []int

func Run_AoC(reader reader.FileReader) {

	var grid [][]int

	for {
		s, err := reader.SplitNextIntLine("")
		if len(s) > 0 {
			grid = append(grid, s)
		}
		if err != nil {
			break
		}
	}
	fmt.Println(SOLVE2(grid))
}

func SOLVE1(grid [][]int) int {
	seen := make(map[types.Pair[int, int]]bool)
	count := 0
	for i := 1; i < len(grid)-1; i++ {
		height := grid[i][0]
		for j := 1; j < len(grid)-1; j++ { // left to right
			height = maxInt(height, grid[i][j-1])
			curr := grid[i][j]
			if curr > height {
				if !seen[types.Pair[int, int]{N1: i, N2: j}] {
					seen[types.Pair[int, int]{N1: i, N2: j}] = true
					count++
				}
			}
		}
	}
	for i := 1; i < len(grid)-1; i++ {
		height := grid[i][len(grid)-1]
		for j := len(grid) - 2; j > 0; j-- { // right to left
			height = maxInt(height, grid[i][j+1])
			curr := grid[i][j]
			if curr > height {
				if !seen[types.Pair[int, int]{N1: i, N2: j}] {
					seen[types.Pair[int, int]{N1: i, N2: j}] = true
					count++
				}
			}
		}
	}
	for j := 1; j < len(grid)-1; j++ {
		height := grid[0][j]
		for i := 1; i < len(grid)-1; i++ { // top to bottom
			height = maxInt(height, grid[i-1][j])
			curr := grid[i][j]
			if curr > height {
				if !seen[types.Pair[int, int]{N1: i, N2: j}] {
					seen[types.Pair[int, int]{N1: i, N2: j}] = true
					count++
				}
			}
		}
	}
	for j := 1; j < len(grid)-1; j++ {
		height := grid[len(grid)-1][j]
		for i := len(grid) - 2; i > 0; i-- { // bottom to top
			height = maxInt(grid[i+1][j], height)
			curr := grid[i][j]
			if curr > height {
				if !seen[types.Pair[int, int]{N1: i, N2: j}] {
					seen[types.Pair[int, int]{N1: i, N2: j}] = true
					count++
				}
			}
		}
	}
	return count + (4*len(grid) - 4)
}

func SOLVE2(grid [][]int) int {
	best := 0
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid)-1; j++ {
			curr := grid[i][j]
			var _ = curr
			l := 1
			r := 1
			u := 1
			d := 1
			i2 := i - 2
			for i2 > -1 && grid[i][j] > grid[i2+1][j] { // up
				u++
				i2--
			}
			i2 = i + 2
			for i2 < len(grid) && grid[i][j] > grid[i2-1][j] { // down
				d++
				i2++
			}
			j2 := j - 2
			for j2 > -1 && grid[i][j] > grid[i][j2+1] { // left
				l++
				j2--
			}
			j2 = j + 2
			for j2 < len(grid) && grid[i][j] > grid[i][j2-1] { // right
				r++
				j2++
			}
			best = maxInt(best, l*r*u*d)
		}
	}
	return best
}

func maxInt(x int, y int) int {
	return int(math.Max(float64(x), float64(y)))
}
