package leetcode

import (
	"reflect"
	"testing"
)

func Test_insertionSortList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			args: args{
				head: genLinkedList([]int{4, 2, 1, 3}),
			},
			want: genLinkedList([]int{1, 2, 3, 4}),
		},
		{
			args: args{
				head: genLinkedList([]int{-1, 5, 3, 4, 0}),
			},
			want: genLinkedList([]int{-1, 0, 3, 4, 5}),
		},
		{
			args: args{
				head: genLinkedList([]int{}),
			},
			want: genLinkedList([]int{}),
		},
		{
			args: args{
				head: genLinkedList([]int{1}),
			},
			want: genLinkedList([]int{1}),
		},
		{
			args: args{
				head: genLinkedList([]int{2, 1}),
			},
			want: genLinkedList([]int{1, 2}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertionSortList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertionSortList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func genLinkedList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{Val: arr[0]}
	cur := head
	for _, v := range arr[1:] {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return head
}
