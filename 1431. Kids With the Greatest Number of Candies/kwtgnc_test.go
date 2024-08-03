package kwtgnc

import "testing"

func Test_kidsWithCandies(t *testing.T) {
	tests := []struct {
		candies      []int
		extraCandies int
		want         []bool
	}{
		{[]int{2, 3, 5, 1, 3}, 3, []bool{true, true, true, false, true}},
		{[]int{4, 2, 1, 1, 2}, 1, []bool{true, false, false, false, false}},
		{[]int{12, 1, 12}, 10, []bool{true, false, true}},
		{[]int{1, 1, 1, 1, 1}, 0, []bool{true, true, true, true, true}},
		{[]int{1, 1, 1, 1, 1}, 1, []bool{true, true, true, true, true}},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			got := kidsWithCandies(test.candies, test.extraCandies)
			for i := range got {
				if got[i] != test.want[i] {
					t.Errorf("kidsWithCandies() = %v, want %v", got, test.want)
					break
				}
			}
		})
	}

}
