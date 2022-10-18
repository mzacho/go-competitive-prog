package solutions

import (
	"competitive_programming/util/reader"
	"fmt"
	"strings"
)

func Run_codejam2022_1(reader reader.StdInReader) {
	noOfTests := reader.NextInt()
	for i := 0; i < noOfTests; i++ {
		test_case := reader.NextIntLine(" ")
		r := test_case[0]
		c := test_case[1]
		fmt.Printf("Case #%v:\n", i+1)
		fmt.Println(draw(int(r), int(c)))
	}
}

func draw(r, c int) string {
	var str string
	for i := 0; i < r; i++ {
		var str1, str2 string
		for j := 0; j < c; j++ {
			if i == 0 && j == 0 {
				str1 += ".."
				str2 += ".."
			} else {
				str1 += "+-"
				str2 += "|."
			}
		}
		str1 += "+"
		str2 += "|"
		str += fmt.Sprint(str1, "\n", str2, "\n")
		if i == r-1 {
			str += str1
		}
	}
	return strings.Trim(str, "\n")
}
