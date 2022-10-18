package convert

import (
	"fmt"
	"strconv"
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
