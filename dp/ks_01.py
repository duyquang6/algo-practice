class Solution:
    def solve(self, values, weights, W):
        from functools import lru_cache
        @lru_cache(None)
        def dp(n, W):
            return max((values[n-1] + dp(i, W-weights[n-1]) for i in range(1, n)), default=0) if weights[n-1] <= W else 0
        return max((dp(n, W) for n in range(1, len(values) + 1)), default=0)


print(Solution().solve([60, 100, 120], [10, 20, 30], 50))
