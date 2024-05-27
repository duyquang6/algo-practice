# https://leetcode.com/problems/reorder-list/description/

# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def reorderList(self, head: Optional[ListNode]) -> None:
        """
        Do not return anything, modify head in-place instead.
        """
        slow, fast = head, head

        while fast is not None and fast.next is not None:
            fast = fast.next.next
            slow = slow.next

        # reverse half linkedlist
        p, cur, q = None, slow, slow.next

        while cur is not None:
            cur.next = p
            p = cur
            cur = q
            if q is not None:
                q = q.next

        p2 = p
        q2 = p.next

        p = head
        q = head.next

        while p is not None:
            p.next = p2
            if p2 is not None:
                p2.next = q

            p2 = q2
            if q2 is not None:
                q2 = q2.next
            p = q
            if q is not None:
                q = q.next
