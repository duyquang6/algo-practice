# https://leetcode.com/problems/palindrome-linked-list/description/

# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def isPalindrome(self, head: Optional[ListNode]) -> bool:
        slow, fast = head, head

        while fast is not None and fast.next is not None:
            slow = slow.next
            fast = fast.next.next

        p, q, r = None, slow, slow.next
        while q is not None:
            q.next = p
            p = q
            q = r
            if r is not None:
                r = r.next

        q = p
        p = head
        while q is not None:
            if p.val != q.val:
                return False
            p = p.next
            q = q.next

        return True
