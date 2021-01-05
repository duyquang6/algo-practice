# https://leetcode.com/problems/maximum-number-of-balloons/

import collections

ex = collections.Counter('balloon')
class Solution:
    def maxNumberOfBalloons(self, text: str) -> int:
        wc = collections.Counter(text)
        res = float('inf')
        for k, v in ex.items():
            if k not in wc:
                return 0
            res = min(res, wc[k] // v)
        return res