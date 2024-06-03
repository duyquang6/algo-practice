# https://leetcode.com/problems/sort-list/
#
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def sortList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        def merge_sort(root):
            if root is None:
                return root
            if root.next is None:
                return root
            if root.next.next is None:
                if root.val > root.next.val:
                    root.val, root.next.val = root.next.val, root.val
                return root

            left = root
            slow, fast = root, root
            prev_slow = None
            while fast is not None and fast.next is not None:
                prev_slow = slow
                slow = slow.next
                fast = fast.next.next

            prev_slow.next = None
            right = slow
            left = merge_sort(left)
            right = merge_sort(right)
            p, q = left, right
            if p.val > q.val:
                p, q = q, p
            prev_p, prev_q = None, None
            while p is not None or q is not None:
                if p is None:
                    prev_p.next = q
                    break
                if q is None:
                    break
                if p.val < q.val:
                    prev_p = p
                    p = p.next
                else:
                    if prev_p is not None:
                        prev_p.next = q
                    q_next = q.next
                    q.next = p
                    prev_p = q
                    q = q_next
            return left if left.val < right.val else right

        return merge_sort(head)
