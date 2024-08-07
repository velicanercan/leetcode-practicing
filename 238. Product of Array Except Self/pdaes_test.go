package main

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Test 1", args{[]int{1, 2, 3, 4}}, []int{24, 12, 8, 6}},
		{"Test 2", args{[]int{-1, 1, 0, -3, 3}}, []int{0, 0, 9, 0, 0}},
		{"Test 3", args{[]int{1, 2, 3, 4, 5}}, []int{120, 60, 40, 30, 24}},
		{"Test 4", args{[]int{1, 2, 3, 4, 5, 6}}, []int{720, 360, 240, 180, 144, 120}},
		{"Test 5", args{[]int{1, -1, 1, -1, 1, -1}}, []int{-1, 1, -1, 1, -1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := productExceptSelf(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productExceptSelf() = %v, want %v", got, tt.want)
			}
		})
	}
}
