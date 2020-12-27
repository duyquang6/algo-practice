// https://leetcode.com/problems/populating-next-right-pointers-in-each-node-ii/
package leetcode

import "container/list"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	queue := list.New()
	queue.PushFront(root)
	for queue.Len() > 0 {
		size := queue.Len()
		for size > 0 {
			elem := queue.Back()
			node := elem.Value.(*Node)
			if node.Left != nil {
				queue.PushFront(node.Left)
			}
			if node.Right != nil {
				queue.PushFront(node.Right)
			}
			size--
			if size > 0 {
				nextElem := elem.Prev()
				node.Next = nextElem.Value.(*Node)
			}
			queue.Remove(elem)
		}
	}
	return root
}
