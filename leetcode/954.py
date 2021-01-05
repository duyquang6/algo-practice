# https://leetcode.com/problems/array-of-doubled-pairs/
from typing import List
import collections
class Solution:
    def canReorderDoubled(self, A: List[int]) -> bool:
        A.sort()
        freq = collections.Counter(A)
        for v in A:
            if freq[v] == 0:
                continue
            if v * 2 in freq and freq[v*2] > 0:
                freq[v*2] -= 1
                freq[v] -= 1
            elif v % 2 == 0 and v // 2 in freq and freq[v//2] > 0:
                freq[v//2] -= 1
                freq[v] -= 1
        return all(v == 0 for v in freq.values())