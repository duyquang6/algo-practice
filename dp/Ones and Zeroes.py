from typing import List


class Solution:
    def findMaxForm(self, strs: List[str], m: int, n: int) -> int:
        from collections import Counter
        from functools import lru_cache
        data = [Counter(s) for s in strs]

        @lru_cache(None)
        def dp(k, m, n):
            if k == 0:
                return 0
            need_m, need_n = data[k-1]['0'], data[k-1]['1']
            if need_m > m or need_n > n:
                return dp(k-1, m, n)
            return max(dp(k-1, m, n), 1+dp(k-1, m-need_m, n-need_n))
        return max((dp(i, m, n) for i in range(1, len(strs)+1)), default=0)


print(Solution().findMaxForm(["10", "0001", "111001", "1", "0"], 5, 3))
