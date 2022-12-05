package day3

import (
	"competitive_programming/util/convert"
	"competitive_programming/util/reader"
	"competitive_programming/util/types"
	"fmt"
)

type line string

// type line []int
// type line [][]string

func parseLine(s string) line {
	return line(s)
	// return functional.Map(
	// 	strings.Split(s, " "),
	// 	convert.UnsafeToIntBase10,
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
	sum := 0
	var a types.Set[rune]
	var b types.Set[rune]
	var c types.Set[rune]
	for i, l := range tcs[0].lines {
		if i%3 == 0 {
			a = types.NewEmptySet[rune]()
			b = types.NewEmptySet[rune]()
			c = types.NewEmptySet[rune]()
			for _, s := range l {
				a.Add(s)
			}
		} else if i%3 == 1 {
			for _, s := range l {
				b.Add(s)
			}
		} else {
			for _, s := range l {
				c.Add(s)
			}
			sum += convert.RuneToInt(a.Intersect(b).Intersect(c).ToList()[0])
		}
	}

	return sum
}
