# https://leetcode.com/problems/rotate-list/submissions/1269564112/

# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def rotateRight(self, head: Optional[ListNode], k: int) -> Optional[ListNode]:
        if head is None:
            return None
        slow, fast = head, head

        length = 0
        while fast is not None:
            length += 1
            fast = fast.next

        k %= length

        fast = head
        while k > 0:
            fast = fast.next
            k -= 1

        while fast.next is not None:
            fast = fast.next
            slow = slow.next

        fast.next = head
        head = slow.next
        slow.next = None
        return head
