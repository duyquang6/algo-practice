# https://leetcode.com/problems/determine-if-string-halves-are-alike/

class Solution:
    def halvesAreAlike(self, s: str) -> bool:                
        return sum(1 for c in s[:len(s) // 2] if c in set('aieouAIEOU')) == sum(1 for c in s[len(s)//2:] if c in set('aieouAIEOU'))
