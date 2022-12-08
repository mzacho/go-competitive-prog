package day7

import (
	"competitive_programming/util/convert"
	"competitive_programming/util/graph"
	"competitive_programming/util/reader"
	"fmt"
	"reflect"
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

func SOLVE(fl firstLine, tcs []testcase) int {
	in := tcs[0].lines
	var res graph.Dir
	var curr *graph.Dir
	for i:=0; i<len(in);i++ {
		l := in[i]
		if reflect.DeepEqual(l, line{"$", "cd", "/"}) {
			res = graph.Dir{Name: "/"}
			curr = &res
			continue
		}
		if l[0] == "$" {
			if l[1] == "ls" {
				for j:=i+1; j<len(in) && in[j][0] != "$"; j++ {
					if in[j][0] == "dir" {
						curr.Dirs = append(curr.Dirs,
							graph.Dir{
								Name: in[j][1],
								Parent: curr,
							},
						)
					} else {
						curr.Files = append(curr.Files,
							graph.File{
								Size: convert.UnsafeToIntBase10(in[j][0]),
								Name: in[j][1],
							},
						)
					}
					i = j
				}
			} else if l[1] == "cd" {
				if l[2] == ".." {
					curr = curr.Parent
					continue
				} else {
					for i, d := range curr.Dirs {
						if d.Name == l[2] {
							curr = &curr.Dirs[i]
							break
						}
					}
				}
			}
		}
	}
	size(&res)
	// return sum_mink(100000, res)
	_, y := min_above(30000000 - (70000000 - res.Size), res)
	return y
}

func size(d *graph.Dir) int {
	for _, f := range d.Files {
		d.Size += f.Size
	}
	for i := range d.Dirs {
		d.Size += size(&d.Dirs[i])
	}
	return d.Size
}

func sum_mink(i int, d graph.Dir) int {
	var res int
	if d.Size <= i {
		res += d.Size
	}
	for _, dir := range d.Dirs {
		res += sum_mink(i, dir)
	}
	return res
}

func min_above(i int, d graph.Dir) (string, int) {
	best_name := ""
	best_val := -1
	if d.Size >= i {
		best_name = d.Name
		best_val = d.Size
		for _, dir := range d.Dirs {
			x, y := min_above(i, dir)
			if y != -1 && y < best_val {
				best_name = x
				best_val = y
			}
		}
	}
	return best_name, best_val
}