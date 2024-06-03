# https://leetcode.com/problems/merge-k-sorted-lists/

# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:
        def merge(left, right):
            if left > right:
                return None
            if left == right:
                return lists[left]

            mid = left + (right - left) // 2
            left = merge(left, mid)
            right = merge(mid + 1, right)

            l = ListNode()
            # merge sort on left_list and right_list
            p = l
            while left is not None and right is not None:
                if left.val < right.val:
                    p.next = left
                    left = left.next
                else:
                    p.next = right
                    right = right.next
                p = p.next

            if left is not None:
                p.next = left
            if right is not None:
                p.next = right

            return l.next

        return merge(0, len(lists) - 1)
