package mergeSortedLists

import (
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	t.Run("mergeTwoLists", func(t *testing.T) {
		list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
		list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
		got := mergeTwoLists(list1, list2)
		want := &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4}}}}}}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("mergeTwoLists", func(t *testing.T) {
		list1 := &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
		list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
		got := mergeTwoLists(list1, list2)
		want := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4}}}}}}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("mergeTwoLists", func(t *testing.T) {
		list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
		list2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 7}}}
		got := mergeTwoLists(list1, list2)
		want := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 7}}}}}}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

}
