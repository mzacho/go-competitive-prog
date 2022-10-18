package solutions

import (
	"competitive_programming/util/reader"
	"fmt"
	"strings"
)

func Run_codejam2019_2(reader reader.FileReader) {
	t, _ := reader.NextInt()
	for i := 0; i < t; i++ {
		n, _ := reader.NextInt()
		l2, _ := reader.NextLine()
		fmt.Printf("Case #%v: %v\n", i+1, MazeWalk(n, l2))
		// fmt.Printf("Case #%v:\n", i+1)
		// fmt.Println(MY_SOLUTIONY(l1, l2))
	}
}

func MazeWalk(n int, l2 string) (res string) {
	for _, c := range strings.Split(l2, "") {
		if c == "E" {
			res += "S"
		} else {
			res += "E"
		}
	}
	return
}
