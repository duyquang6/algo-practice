# https://leetcode.com/problems/keyboard-row/
from typing import List
dp = [set('qwertyuiop'), set('asdfghjkl'), set('zxcvbnm')]
class Solution:
    def findWords(self, words: List[str]) -> List[str]:
        res = []
        for word in words:
            for _set in dp:
                if all(v in _set for v in word.lower()):
                    res.append(word)
                    break
        return res