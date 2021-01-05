// https://leetcode.com/problems/insertion-sort-list/
package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	res := &ListNode{Val: head.Val}
	p := head.Next

	for p != nil {
		val := p.Val
		var q *ListNode
		qNext := res
		for qNext != nil {
			if qNext.Val > val {
				break
			}
			q = qNext
			qNext = qNext.Next
		}
		node := &ListNode{Val: val}
		if q != nil {
			q.Next = node
		} else {
			res = node
		}
		node.Next = qNext
		p = p.Next
	}
	return res
}
