# https://leetcode.com/problems/linked-list-cycle-ii/

# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None


class Solution:
    def detectCycle(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if head is None or head.next is None:
            return None

        slow = head
        fast = head

        while (
            slow.next is not None
            and fast.next is not None
            and fast.next.next is not None
        ):
            slow = slow.next
            fast = fast.next.next
            if slow == fast:
                break
        else:
            return None

        while head != slow:
            head = head.next
            slow = slow.next

        return head
