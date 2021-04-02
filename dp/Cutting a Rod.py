class Solution:
    def solve(self, rod_length, prices):

        from functools import lru_cache

        @lru_cache(None)
        def dp(n, weight):
            if weight < 0:
                return float("-inf")
            if n == 0:
                return 0
            return max(dp(n, weight-n) + prices[n-1], dp(n-1, weight))

        return dp(len(prices), rod_length)


print(Solution().solve(8, [1, 5,  8,   9,  10,  17,  17,  20]))
print(Solution().solve(8, [3, 5, 8, 9, 10, 17,  17, 20]))
