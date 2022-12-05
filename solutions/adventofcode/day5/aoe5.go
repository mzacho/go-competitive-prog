package day5

import (
	"competitive_programming/util/convert"
	"competitive_programming/util/reader"
	"fmt"
	"strings"
)

type line []string
// type line []int
// type line [][]string

func parseLine(s string) line {
	return strings.Split(s, " ")
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
	fmt.Println(SOLVE(fl, tcs))
}

// [N] [G]                     [Q]
// [H] [B]         [B] [R]     [H]
// [S] [N]     [Q] [M] [T]     [Z]
// [J] [T]     [R] [V] [H]     [R] [S]
// [F] [Q]     [W] [T] [V] [J] [V] [M]
// [W] [P] [V] [S] [F] [B] [Q] [J] [H]
// [T] [R] [Q] [B] [D] [D] [B] [N] [N]
// [D] [H] [L] [N] [N] [M] [D] [D] [B]
//  1   2   3   4   5   6   7   8   9

func SOLVE(fl firstLine, tcs []testcase) int {
	var stacks = [][]string{
		{"D", "T", "W" , "F", "J", "S", "H", "N"},
		{"H", "R", "P" , "Q", "T", "N", "B", "G"},
		{"L", "Q", "V"},
		{"N", "B", "S", "W", "R", "Q"},
		{"N", "D", "F", "T", "V", "M", "B"},
		{"M", "D", "B", "V", "H", "T", "R"},
		{"D", "B", "Q", "J"},
		{"D", "N", "J", "V", "R", "Z", "H", "Q"},
		{"B", "N", "H", "M", "S"},
	}
	for _, l := range tcs[0].lines {
		move := convert.UnsafeToIntBase10(l[1])
		from := convert.UnsafeToIntBase10(l[3])
		to := convert.UnsafeToIntBase10(l[5])
		// for i:=0; i<move; i++ {
		// 	x := stacks[from-1][len(stacks[from-1])-1]
		// 	stacks[from-1] = stacks[from-1][:len(stacks[from-1])-1]
		// 	stacks[to-1] = append(stacks[to-1], x)
		// }
		x := stacks[from-1][len(stacks[from-1])-move:]
			stacks[from-1] = stacks[from-1][:len(stacks[from-1])-move]
			stacks[to-1] = append(stacks[to-1], x...)
	}
	return 42
}