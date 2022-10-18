package solutions

import (
	"competitive_programming/util/array"
	"competitive_programming/util/reader"
	"fmt"
)

func Run_codejam2022_2(reader reader.StdInReader) {
	noOfTests := reader.NextInt()
	for i := 0; i < noOfTests; i++ {
		l1 := reader.NextIntLine(" ")
		l2 := reader.NextIntLine(" ")
		l3 := reader.NextIntLine(" ")
		fmt.Printf("Case #%v:\n", i+1)
		fmt.Println(findColor(l1, l2, l3))
	}
}

func findColor(l1, l2, l3 []int) string {
	_, n1 := array.Min(l1[0], l2[0], l3[0])
	if n1 >= 1000000 {
		return fmt.Sprintf("%v 0 0 0", n1)
	}
	_, n2 := array.Min(l1[1], l2[1], l3[1])
	if n1+n2 >= 1000000 {
		return fmt.Sprintf("%v %v 0 0", n1, 1000000-n1)
	}
	_, n3 := array.Min(l1[2], l2[2], l3[2])
	if n1+n2+n3 >= 1000000 {
		return fmt.Sprintf("%v %v %v 0", n1, n2, 1000000-n1-n2)
	}
	_, n4 := array.Min(l1[3], l2[3], l3[3])
	if n1+n2+n3+n4 >= 1000000 {
		return fmt.Sprintf("%v %v %v %v", n1, n2, n3, 1000000-n1-n2-n3)
	}
	return "IMPOSSIBLE"
}
