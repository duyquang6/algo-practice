# https://leetcode.com/problems/compare-strings-by-frequency-of-the-smallest-character/
import functools, collections
from typing import List
class Solution:
    def numSmallerByFrequency(self, queries: List[str], words: List[str]) -> List[int]:
        @functools.lru_cache(None)
        def calcFreqSmallest(s):
            freq = 0
            wc = collections.Counter(s)
            for c in range(ord('a'), ord('z')+1):
                if chr(c) in wc:
                    return wc[chr(c)]
        
        queries = [calcFreqSmallest(q) for q in queries]
        words = [calcFreqSmallest(q) for q in words]
        return [sum(1 for w in words if w > q) for q in queries]