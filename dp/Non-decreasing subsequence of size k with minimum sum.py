class Solution:
    def solve(self, arr, k):
        from functools import lru_cache

        @lru_cache(None)
        def dp(n, k):
            if k == 1:
                return arr[n-1]
            res = float('inf')
            for i in range(n-1, 0, -1):
                if arr[i-1] <= arr[n-1]:
                    res = min(dp(i, k-1) + arr[n-1], res)
            return res
        res = float('inf')
        for i in range(1, len(arr) + 1):
            res = min(res, dp(i, k))
        return res


print(Solution().solve([58, 12, 11, 12, 82, 30, 20, 77, 16, 86], 5))
