
class Solution:
    def solve(self, n):

        from functools import lru_cache

        def dp(n):
            if n == 1:
                return 1
            res = 0
            for k in range(1, n+1):
                res = max(k*(n-k), dp(n-k) * k, res)
            return res

        return dp(n)


print(Solution().solve(2))
print(Solution().solve(3))
print(Solution().solve(4))
print(Solution().solve(5))
print(Solution().solve(10))
