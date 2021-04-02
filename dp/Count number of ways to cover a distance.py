class Solution:
    def solve(self, n):
        
        from functools import lru_cache

        @lru_cache(None)
        def dp(n):
            if n < 0:
                return 0
            if n == 0:
                return 1
            res = 0
            for step in range(1,4):
                res += dp(n-step)
            return res

        return dp(n)

print(Solution().solve(4))