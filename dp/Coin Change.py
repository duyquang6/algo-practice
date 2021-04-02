class Solution:
    def solve(self, N, cents):
        cents.sort()
        from functools import lru_cache

        @lru_cache(None)
        def dp(start, n):
            if n == 0:
                return 1
            res = 0
            for i in range(start, len(cents)):
                if cents[i] > n:
                    break
                res += dp(i, n-cents[i])
            return res
        return dp(0, N)


print(Solution().solve(4, [1, 2, 3]))
print(Solution().solve(10, [2, 5, 3, 6]))
