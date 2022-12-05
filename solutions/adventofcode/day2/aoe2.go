package solutions

import (
	"competitive_programming/util/reader"
	"fmt"
	"strings"
)

func Run_AoC(reader reader.FileReader) {
	var lines [][]string
	for {
		s, err := reader.NextLine()
		if err != nil {
			break
		}
		l := strings.Split(s, " ")
		lines = append(lines, []string{l[0], l[1]})
	}
	fmt.Println(max_score(lines))
}

func max_score(input [][]string) int {
	var t_sum int
	for _, l := range input {
		if l[0] == "A" { // rock
			if l[1] == "X" {
				t_sum += 3 + 0
				} else if l[1] == "Y" {
					t_sum += 1 + 3
				} else if l[1] == "Z" {
					t_sum += 2 + 6
			}
		} else if l[0] == "B" { // paper
			if l[1] == "X" {
				t_sum += 1 + 0
				} else if l[1] == "Y" {
					t_sum += 2 + 3
				} else if l[1] == "Z" {
					t_sum += 3 + 6
			}
		} else if l[0] == "C" { // scissor
			if l[1] == "X" {
				t_sum += 2 + 0
				} else if l[1] == "Y" { // paper
					t_sum += 3 + 3
				} else if l[1] == "Z" {
					t_sum += 1 + 6
			}
		}
	}
	return t_sum
}
