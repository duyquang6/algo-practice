# https://leetcode.com/problems/find-words-that-can-be-formed-by-characters/
from typing import List
import collections
class Solution:
    def countCharacters(self, words: List[str], chars: str) -> int:
        wc = collections.Counter(chars)
        res = 0
        for w in words:
            _wc = collections.Counter(w)            
            if all(wc[k] >= v for k,v in _wc.items()):
                res += len(w)
                
        return res