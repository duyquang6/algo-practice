class Solution:
    def solve(self, arr):
        from functools import lru_cache

        @lru_cache(None)
        def dp(n):
            res = 1
            for k in range(n):
                if arr[n] > arr[k]:
                    res = max(res, dp(k) + 1)
            return res

        res = 0
        for i in range(len(arr)):
            res = max(res, dp(i))
        return res


print(Solution().solve([10, 22, 9, 33, 21, 50, 41, 60, 80]))
print(Solution().solve([3, 10, 2, 1, 20]))
print(Solution().solve([3, 2]))
print(Solution().solve([50, 3, 10, 7, 40, 80]))
