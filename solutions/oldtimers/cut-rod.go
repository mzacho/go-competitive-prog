package solutions

// import (
// 	"competitive_programming/util/array"
// 	"competitive_programming/util/reader"
// 	"fmt"
// 	"strings"
// )

// func Run_cutrod(reader reader.FileReader) {
// 	t, _ := reader.NextInt()
// 	for i := 0; i < t; i++ {
// 		l1 := reader.NextInt()
// 		l2 := reader.SplitNextIntLine(" ")
// 		fmt.Printf("Case #%v:\n", i+1)
// 		print_solution(l1, l2)
// 	}
// }

// func Memoized_cut_rod(l1 int, l2 []int) int {
// 	var results = make([]int, l1)
// 	for i := 0; i < len(results); i++ {
// 		results[i] = -1
// 	}
// 	return memoized_cut_rod_aux(l1, l2, results)
// }

// func memoized_cut_rod_aux(n int, prices []int, results []int) int {
// 	if n == 0 {
// 		return 0
// 	}
// 	if results[n-1] >= 0 {
// 		return results[n-1]
// 	} else {
// 		var res int
// 		for i := 1; i <= n; i++ {
// 			_, res = array.Max(res, prices[i-1]+memoized_cut_rod_aux(n-i, prices, results))
// 		}
// 		results[n-1] = res
// 		return res
// 	}
// }

// func Bottom_up_cut_rod(n int, prices []int) int {
// 	var result = make([]int, n+1)
// 	result[0] = 0
// 	for j := 1; j <= n; j++ {
// 		solution := -1
// 		for i := 0; i < j; i++ {
// 			_, solution = array.Max(solution, prices[i]+result[j-i-1])
// 		}
// 		result[j] = solution
// 	}
// 	return result[n]
// }

// func extended_bottom_up_cut_rod(n int, prices []int) ([]int, []int) {
// 	var result = make([]int, n+1)
// 	var solution = make([]int, n+1)
// 	result[0] = 0
// 	for j := 1; j <= n; j++ {
// 		cost := -1
// 		for i := 0; i < j; i++ {
// 			if cost < prices[i]+result[j-i-1] {
// 				cost = prices[i] + result[j-i-1]
// 				solution[j] = i + 1
// 			}
// 		}
// 		result[j] = cost
// 	}
// 	return result, solution
// }

// func print_solution(n int, prices []int) {
// 	var res []string
// 	r, s := extended_bottom_up_cut_rod(n, prices)
// 	for n > 0 {
// 		res = append(res, fmt.Sprint(s[n]))
// 		n -= s[n]
// 	}
// 	fmt.Printf("Optimal cuts: %v\n", strings.Join(res, " "))
// 	fmt.Printf("Optimal cost: %v\n", r[n])
// }
