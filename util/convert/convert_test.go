package convert

import (
	"testing"
)

func TestConvertByteToInt(t *testing.T) {
	if RuneToInt('a') != 1 {
		t.Errorf("Expected RuneToInt('a') == 1; was %v", RuneToInt('a'))
	}
	if RuneToInt('z') != 26 {
		t.Errorf("Expected RuneToInt('z') == 26; was %v", RuneToInt('z'))
	}
	if RuneToInt('A') != 27 {
		t.Errorf("Expected RuneToInt('A') == 27; was %v", RuneToInt('A'))
	}
	if RuneToInt('Z') != 52 {
		t.Errorf("Expected RuneToInt('Z') == 52; was %v", RuneToInt('Z'))
	}
}