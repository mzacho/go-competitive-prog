package solutions

// import (
// 	"competitive_programming/util/array"
// 	"competitive_programming/util/reader"
// 	"competitive_programming/util/types"
// )

// func Run_knapsack_0_1(reader reader.FileReader) {
// 	values := reader.SplitNextIntLine(" ")
// 	weights := reader.SplitNextIntLine(" ")
// 	var items types.ItemSlice
// 	for _, p := range array.Zip(values, weights) {
// 		items = append(items, types.Item{
// 			Value:  p.N1,
// 			Weight: p.N2,
// 		})
// 	}
// 	max_weight, _ := reader.NextInt()
// 	Knapsack_0_1(max_weight, items)
// }

// func Knapsack_0_1(W int, items types.ItemSlice) {
// 	var c = make([][]int, len(items))
// 	for i := 0; i < len(items); i++ {
// 		c[i] = make([]int, W)
// 	}
// 	for w := 0; w < W; w++ {
// 		c[0][w] = 0
// 	}
// 	for i := 1; i < len(items); i++ {
// 		c[i][0] = 0
// 		for w := 1; w < W; w++ {
// 			if items[i].Weight <= w {
// 				if items[i].Value+c[i-1][w-items[i].Weight] > c[i-1][w] {
// 					c[i][w] = items[i].Value + c[i-1][w-items[i].Weight]
// 				} else {
// 					c[i][w] = c[i-1][w]
// 				}
// 			} else {
// 				c[i][w] = c[i-1][w]
// 			}
// 		}
// 	}
// }

// // func Memoized_cut_rod(l1 int, l2 []int) int {
// // 	var results = make([]int, l1)
// // 	for i := 0; i < len(results); i++ {
// // 		results[i] = -1
// // 	}
// // 	return memoized_cut_rod_aux(l1, l2, results)
// // }

// // func memoized_cut_rod_aux(n int, prices []int, results []int) int {
// // 	if n == 0 {
// // 		return 0
// // 	}
// // 	if results[n-1] >= 0 {
// // 		return results[n-1]
// // 	} else {
// // 		var res int
// // 		for i := 1; i <= n; i++ {
// // 			_, res = array.Max(res, prices[i-1]+memoized_cut_rod_aux(n-i, prices, results))
// // 		}
// // 		results[n-1] = res
// // 		return res
// // 	}
// // }

// // func Bottom_up_cut_rod(n int, prices []int) int {
// // 	var result = make([]int, n+1)
// // 	result[0] = 0
// // 	for j := 1; j <= n; j++ {
// // 		solution := -1
// // 		for i := 0; i < j; i++ {
// // 			_, solution = array.Max(solution, prices[i]+result[j-i-1])
// // 		}
// // 		result[j] = solution
// // 	}
// // 	return result[n]
// // }

// // func extended_bottom_up_cut_rod(n int, prices []int) ([]int, []int) {
// // 	var result = make([]int, n+1)
// // 	var solution = make([]int, n+1)
// // 	result[0] = 0
// // 	for j := 1; j <= n; j++ {
// // 		cost := -1
// // 		for i := 0; i < j; i++ {
// // 			if cost < prices[i]+result[j-i-1] {
// // 				cost = prices[i] + result[j-i-1]
// // 				solution[j] = i + 1
// // 			}
// // 		}
// // 		result[j] = cost
// // 	}
// // 	return result, solution
// // }

// // func print_solution(n int, prices []int) {
// // 	var res []string
// // 	r, s := extended_bottom_up_cut_rod(n, prices)
// // 	for n > 0 {
// // 		res = append(res, fmt.Sprint(s[n]))
// // 		n -= s[n]
// // 	}
// // 	fmt.Printf("Optimal cuts: %v\n", strings.Join(res, " "))
// // 	fmt.Printf("Optimal cost: %v\n", r[n])
// // }
