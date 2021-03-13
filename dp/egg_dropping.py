class Solution:
    def memoization(self, k: int, n: int) -> int:
        from functools import lru_cache
        
        @lru_cache(None)
        def helper(k, n):
            if n == 0:
                return 0
            if k == 1:
                return n
            res = float('inf')
            for x in range(1, n + 1):
                res = min(1 + max(helper(k - 1, x - 1), helper(k, n - x)), res)
            return res
        
        return helper(k, n)