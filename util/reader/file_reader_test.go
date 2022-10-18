package reader_test

import (
	"competitive_programming/util/reader"
	"testing"
)

func TestFileReader(t *testing.T) {
	r := reader.NewFileReader("test-input1")
	l1, _ := r.NextInt()
	if l1 != 42 {
		t.Errorf("expected 42; was %v", l1)
	}
	l2, _ := r.SplitNextIntLine(" ")
	if l2[0] != 43 || l2[1] != 44 {
		t.Errorf("expected 43 44; was %v", l2)
	}
	l3, _ := r.SplitNextIntLine(",")
	if l3[0] != 43 || l3[1] != 44 {
		t.Errorf("expected 43,44; was %v", l3)
	}
	l4, _ := r.NextLine()
	if l4 != "hello world" {
		t.Errorf("expected 'hello world'; was %v", l4)
	}
	l5, _ := r.SplitNextLine(" ")
	if l5[0] != "foo" || l5[1] != "bar" {
		t.Errorf("expected 'foo bar'; was %v", l5)
	}
}
