package statemachine

import "fmt"

type SM struct {
	Current State
}

type State int

const (
	A State = iota + 1
	B
	C
)

func (sm *SM) Process(s State) {
	switch sm.Current {
	case A:
		switch s {
		case A:
			panic(fmt.Errorf("transitioning from %v to %v", A, A))
		case B:
			panic(fmt.Errorf("transitioning from %v to %v", A, B))
		case C:
			panic(fmt.Errorf("transitioning from %v to %v", A, C))
		}
	case B:
		switch s {
		case A:
			panic(fmt.Errorf("transitioning from %v to %v", B, A))
		case B:
			panic(fmt.Errorf("transitioning from %v to %v", B, B))
		case C:
			panic(fmt.Errorf("transitioning from %v to %v", B, C))
		}
	case C:
		switch s {
		case A:
			panic(fmt.Errorf("transitioning from %v to %v", C, A))
		case B:
			panic(fmt.Errorf("transitioning from %v to %v", C, B))
		case C:
			panic(fmt.Errorf("transitioning from %v to %v", C, C))
		}
	}
}
