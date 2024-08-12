package main

import "testing"

func TestMaxArea(t *testing.T) {
	testCases := []struct {
		Input  []int
		Expect int
	}{
		{Input: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, Expect: 49},
		{Input: []int{1, 1}, Expect: 1},
		{Input: []int{4, 3, 2, 1, 4}, Expect: 16},
		{Input: []int{1, 2, 1}, Expect: 2},
		{Input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, Expect: 25},
	}

	for _, c := range testCases {
		result := maxArea(c.Input)
		if result != c.Expect {
			t.Fatalf("Expect: %d, Get: %d\n", c.Expect, result)
		}
	}

}
