package day10

import (
	"competitive_programming/util/array"
	"competitive_programming/util/convert"
	"competitive_programming/util/mathz"
	"competitive_programming/util/reader"
	t "competitive_programming/util/types"
	"fmt"
	"strings"
)

type operation int

const (
	addx operation = iota + 1
	noop
)

func parseOperation(s string) operation {
	switch s {
	case "addx":
		return addx
	case "noop":
		return noop
	default:
		panic("cannot parse operation")
	}
}

type line t.Pair[operation, int]

// type line []int
// type line [][]string

func parseLine(s string) line {
	x := strings.Split(s, " ")
	if len(x) > 1 {
		return line(t.Pair[operation, int]{
			N1: parseOperation(x[0]),
			N2: convert.UnsafeToIntBase10(x[1]),
		})
	} else {
		return line(t.Pair[operation, int]{
			N1: parseOperation(x[0]),
			N2: -1,
		})
	}
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
	fmt.Println(SOLVE(fl, tcs))
}

func SOLVE(fl firstLine, tcs []testcase) int {
	// k := []int{20, 60, 100, 140, 180, 220}
	x := 1
	s := []int{}
	cycle := 0
	l := ""
	for i := 0; i < len(tcs[0].lines); i++ {
		cycle++
		// if array.Contains(k, cycle) {
		// 	s = append(s, cycle * x)
		// }
		if mathz.AbsInt(x-len(l)) < 2 {
			l += "#"
		} else {
			l += "."
		}
		if len(l) == 40 {
			fmt.Println(l)
			l = ""
		}
		if tcs[0].lines[i].N1 == addx {
			cycle++
			// if array.Contains(k, cycle) {
			// 	s = append(s, cycle * x)
			// }
			if mathz.AbsInt(x-len(l)) < 2 {
				l += "#"
			} else {
				l += "."
			}
			if len(l) == 40 {
				fmt.Println(l)
				l = ""
			}
		}
		if tcs[0].lines[i].N1 == addx {
			x += tcs[0].lines[i].N2
		}
	}
	return array.Sum(s)
}
