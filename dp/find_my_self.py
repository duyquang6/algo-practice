from functools import lru_cache


class Solution:
    def so(self, N, W, costs):
        def calc_cost(arr, qty):
            if qty == 0:
                return 0
            return arr[0] * qty + arr[1]

        @lru_cache(None)
        def dp(n, qty):
            if qty == 0:
                return 0
            if n == 0:
                return float('inf')
            res = float("inf")
            for i in range(qty+1):
                res = min(res, dp(n-1, i) + calc_cost(costs[n-1], qty-i))
            return res
        return dp(len(W), N)


print(Solution().so(5, [0, 1, 2], [(1, 2), (2, 1), (1, 3)]))
