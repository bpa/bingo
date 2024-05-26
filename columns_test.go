package main

import (
	"reflect"
	"testing"
)

func TestCombos(t *testing.T) {
	combos := generateCombinations(3, 2)
	if len(combos) != 3 {
		t.Fatalf(`Expected 3 combos, got %d`, len(combos))
	}
	expected := [3][]int{{1, 0}, {2, 0}, {2, 1}}
	for i := range len(combos) {
		if !reflect.DeepEqual(combos[i], expected[i]) {
			t.Fatalf(`Actual %v != expected %v`, combos, expected)
		}
	}
}

func TestPermutations(t *testing.T) {
	for i := range 5 {
		if i > 1 {
			perms := permutations(i)
			if len(perms) != fact(i) {
				t.Fatalf(`len(permutations(%d)) [%d] != %d`, i, len(perms), fact(i))
			}
		}
	}
}
