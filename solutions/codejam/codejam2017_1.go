package solutions

import (
	"competitive_programming/util/convert"
	"competitive_programming/util/reader"
	"fmt"
	"strings"
)

func Run_codejam2017_1(reader reader.FileReader) {
	t, _ := reader.NextInt()
	for i := 0; i < t; i++ {
		l, _ := reader.NextLineSplit(" ")
		l1 := convert.UnsafeToIntBase10(l[1])
		l2 := l[0]
		fmt.Printf("Case #%v: %v\n", i+1, convert.IntToStringWithDefault(FlipPancacesAux(l1, strings.Split(l2, ""), 0), "IMPOSSIBLE"))
		// fmt.Printf("Case #%v:\n", i+1)
		// fmt.Println(MY_SOLUTIONY(l1, l2))
	}
}

func FlipPancacesAux(k int, l2 []string, s int) int {
	if k > len(l2) {
		return -1
	}
	if k == len(l2) {
		seenPlus := false
		for i := 0; i < len(l2); i++ {
			if l2[i] == "+" {
				seenPlus = true
			}
			if l2[i] == "-" && seenPlus {
				return -1
			}
		}
		if seenPlus {
			return s
		}
		return s + 1
	}
	if l2[0] == "+" {
		return FlipPancacesAux(k, l2[1:], s)
	} else {
		var r string
		for i := 1; i < k; i++ {
			if l2[i] == "+" {
				r += "-"
			} else {
				r += "+"
			}
		}
		p := l2[k:]
		return FlipPancacesAux(k, append(strings.Split(r, ""), p...), s+1)
	}
}
