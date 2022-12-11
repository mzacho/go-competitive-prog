package types

import (
	"strings"
)

type Direction int

const (
	Up Direction = iota + 1
	Down
	Left
	Right
)

func ParseDirection(s string) Direction {
	switch (strings.ToLower(s)) {
	case "u":
		return Up
	case "d":
		return Down
	case "l":
		return Left
	case "r":
		return Right
	default:
		panic("cannot parse direction " + s)
	}
}