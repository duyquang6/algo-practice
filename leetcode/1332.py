# https://leetcode.com/problems/remove-palindromic-subsequences/

class Solution:
    def removePalindromeSub(self, s: str) -> int:
        return 0 if not s else 1 if all(s[i] ==s[-i - 1] for i in range(len(s)//2)) else 2