package day9

import (
	"competitive_programming/util/array"
	"competitive_programming/util/convert"
	m "competitive_programming/util/mathz"
	"competitive_programming/util/reader"
	"competitive_programming/util/types"
	"fmt"
	"strings"
)

type line types.Pair[types.Direction, int]

// type line []int
// type line [][]string

func parseLine(s string) line {
	x := strings.Split(s, " ")
	return line(types.Pair[types.Direction, int]{
		N1: types.ParseDirection(x[0]),
		N2: convert.UnsafeToIntBase10(x[1]),
	})
	// return functional.Map(
	// 	strings.Split(s, " "),
	// 	convert.UnsafeToIntBase10,
	// )
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
	var tcs []testcase
	var fl firstLine

	// Some days the first line needs special attention ...
	if false {
		// s, _ := reader.NextLineSplit(",")
		s, _ := reader.SplitNextIntLine(",")
		fl = firstLine(s)
	}

	for {
		s, err := reader.NextLine()
		if err != nil {
			break
		}
		var t testcase
		var lines []line
		for s != "" && err == nil {
			lines = append(lines, parseLine(s))
			s, err = reader.NextLine()
		}
		if len(lines) > 0 {
			t = parseTestCase(lines)
			tcs = append(tcs, t)
		}
	}
	fmt.Println(SOLVE(fl, tcs, 9))
}

func move(p *types.IntPair, d types.Direction) {
	switch d {
	case types.Up:
		p.N2++
	case types.Down:
		p.N2--
	case types.Right:
		p.N1++
	case types.Left:
		p.N1--
	}
}

func SOLVE(fl firstLine, tcs []testcase, k int) int {
	seen := make(types.IntPairSet)
	l := array.Make(types.IntPair{N1: 1, N2: 1}, k+1)
	for _, p := range tcs[0].lines {
		for i := 0; i < p.N2; i++ {
			move(&l[0], p.N1)
			for j := 1; j < len(l); j++ {
				if m.AbsInt(l[j-1].N1-l[j].N1) < 2 && m.AbsInt(l[j].N2-l[j-1].N2) < 2 {
					continue
				}
				if l[j].N1 < l[j-1].N1 {
					move(&l[j], types.Right)
				} else if l[j].N1 > l[j-1].N1 {
					move(&l[j], types.Left)
				}
				if l[j].N2 < l[j-1].N2 {
					move(&l[j], types.Up)
				} else if l[j].N2 > l[j-1].N2 {
					move(&l[j], types.Down)
				}
				if j == len(l)-1 {
					seen[l[j]] = true
				}
			}
		}
	}
	return len(seen)
}
