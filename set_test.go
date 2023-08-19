package go_set_test

import (
	gs "github.com/yycsec1/go-set"
	"testing"
)

func TestSetLen(t *testing.T) {
	s := gs.New[string]()
	s.Add("One")
	s.Add("Two")
	result := s.Len()
	if result != 2 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 2)
	}
	s.Discard("One")
	result = s.Len()
	if result != 1 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 1)
	}
	s.Discard("Two")
	result = s.Len()
	if result != 0 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 0)
	}
}

func TestIsSubsetAndIsSuperset(t *testing.T) {
	list1 := []string{"a", "b", "c"}
	list2 := []string{"f", "e", "d", "c", "b", "a"}
	s1 := gs.NewFrom[string](&list1)
	s2 := gs.NewFrom[string](&list2)
	if !s1.IsSubset(s2) {
		t.Errorf("Result was incorrect, got: %t, want: %t.", false, true)
	}
	if !s2.IsSuperset(s1) {
		t.Errorf("Result was incorrect, got: %t, want: %t.", false, true)
	}
}
