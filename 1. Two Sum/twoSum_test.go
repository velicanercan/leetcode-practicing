package twoSum

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	t.Run("twoSum", func(t *testing.T) {
		nums := []int{2, 7, 11, 15}
		target := 9
		got := twoSum(nums, target)
		want := []int{0, 1}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("twoSum", func(t *testing.T) {
		nums := []int{3, 2, 4}
		target := 6
		got := twoSum(nums, target)
		want := []int{1, 2}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("twoSum", func(t *testing.T) {
		nums := []int{3, 3}
		target := 6
		got := twoSum(nums, target)
		want := []int{0, 1}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
