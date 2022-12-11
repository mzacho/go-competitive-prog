package types

type Pair[T, R any] struct {
	N1 T
	N2 R
}

type IntPair = Pair[int, int]
