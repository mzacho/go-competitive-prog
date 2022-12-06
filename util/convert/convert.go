package convert

import (
	"fmt"
	"strconv"
	"unicode"
)

func UnsafeToInt(base int, s string) int {
	i, _ := strconv.ParseInt(s, base, 0)
	return int(i)
}

func UnsafeToIntBase10(s string) int {
	i, _ := strconv.ParseInt(s, 10, 0)
	return int(i)
}

func IntToString(i int) string {
	return fmt.Sprint(i)
}

func IntToStringWithDefault(i int, default_ string) string {
	if i == -1 {
		return default_
	}
	return fmt.Sprint(i)
}

// Converts [a-z] to 1..26 and [A-Z] to 27..52
func RuneToInt(b rune) int {
	if unicode.IsLower(b) {
		return int(b) - 96
	} else {
		return int(b) - 38
	}
}
