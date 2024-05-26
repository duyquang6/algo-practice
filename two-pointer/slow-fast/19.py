# https://leetcode.com/problems/remove-nth-node-from-end-of-list/
#
#
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
        slow = head
        fast = head
        while n > 0:
            fast = fast.next
            n -= 1
        prev = None
        while fast is not None:
            fast = fast.next
            prev = slow
            slow = slow.next

        if prev is None:
            return slow.next
        prev.next = slow.next

        return head
