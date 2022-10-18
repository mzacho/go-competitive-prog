package types_test

import (
	"competitive_programming/util/types"
	"testing"
)

func TestAddToSet1(t *testing.T) {
	a := types.NewSet(1, 2, 3, 5)
	updated := a.Add(6)

	if len(a) != 5 {
		t.Errorf("Expected len(a) = 5; was %v", len(a))
	}
	if !updated {
		t.Errorf("Expected updated = true; was %v", len(a))
	}
}

func TestAddToSet2(t *testing.T) {
	a := types.NewSet(1, 2, 3, 5)
	updated := a.Add(5)

	if len(a) != 4 {
		t.Errorf("Expected len(a) = 4; was %v", len(a))
	}
	if updated {
		t.Errorf("Expected updated = false; was %v", len(a))
	}
}

func TestRemoveFromSet1(t *testing.T) {
	a := types.NewSet(1, 2, 3, 5)
	updated := a.Remove(2)

	if len(a) != 3 {
		t.Errorf("Expected len(a) = 3; was %v", len(a))
	}
	if !updated {
		t.Errorf("Expected updated = true; was %v", len(a))
	}
}

func TestRemoveFromSet2(t *testing.T) {
	a := types.NewSet(1, 2, 3, 5)
	updated := a.Remove(6)

	if len(a) != 4 {
		t.Errorf("Expected len(a) = 4; was %v", len(a))
	}
	if updated {
		t.Errorf("Expected updated = false; was %v", len(a))
	}
}

func TestSetUnion(t *testing.T) {
	a := types.NewSet(1, 2, 3, 5)
	b := types.NewSet(1, 2, 4)

	if len(a.Union(b)) != 5 {
		t.Errorf("Expected len(a.Union(b)) = 3; was %v", len(a.Union(b)))
	}
}

func TestMutableSetUnion(t *testing.T) {
	a := types.NewSet(1, 2, 3, 5)
	b := types.NewSet(1, 2, 4)
	c := a.Union(b)

	if len(c) != 5 {
		t.Errorf("Expected len(a.Union(b)) = 3; was %v", len(c))
	}
}

func TestSetMinus(t *testing.T) {
	a := types.NewSet(1, 2, 3, 5)
	b := types.NewSet(1, 2, 4)

	if len(a.Minus(b)) != 2 {
		t.Errorf("Expected len(a.Minus(b)) = 2; was %v", len(a.Minus(b)))
	}
}

func TestMutableSetMinus(t *testing.T) {
	a := types.NewSet(1, 2, 3, 5)
	b := types.NewSet(1, 2, 4)
	c := a.Minus(b)

	if len(c) != 2 {
		t.Errorf("Expected len(a.Minus(b)) = 2; was %v", len(c))
	}
}

func TestSetIntersection(t *testing.T) {
	a := types.NewSet(1, 2, 3, 5)
	b := types.NewSet(1, 2, 4)

	a2 := a.Intersect(b).Union(a.Minus(b))

	if !a2.Equals(a) {
		t.Errorf("Expected a2 == a; was %v", a2.Equals(a))
	}
}

func TestMutableSetIntersection(t *testing.T) {
	a := types.NewSet(1, 2, 3, 5)
	a2 := types.NewSet(1, 2, 3, 5)
	a3 := types.NewSet(1, 2, 3, 5)
	b := types.NewSet(1, 2, 4)

	a.MutableIntersect(b)
	a2.MutableMinus(b)
	a.MutableUnion(a2)

	if !a.Equals(a3) {
		t.Errorf("Expected a == a3; was %v", a.Equals(a3))
	}
}