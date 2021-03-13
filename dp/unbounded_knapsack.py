import functools


class Solution:
    def memoization(self, W, wt, val):
        best = 0
        min_weight = min(wt)

        @functools.lru_cache(None)
        def helper(weight_avail):
            res = 0
            for i, w in enumerate(wt):
                if weight_avail - w < 0:
                    continue
                res = max(res, helper(weight_avail - w) + val[i])
            return res

        return helper(W)

    def tabulation(self, W, wt, val):
        dp = [0] * (W+1)
        min_weight = min(wt)
        for w in range(min_weight, W + 1):
            for i, wi in enumerate(wt):
                if w - wi < 0:
                    continue
                dp[w] = max(dp[w], dp[w - wi] + val[i])

        return dp[W]


print(Solution().tabulation(100, [1, 50], [1, 30])
      == Solution().memoization(100, [1, 50], [1, 30]))
