class Solution:
    def memoi(self, arr):
        from functools import lru_cache
        res = 0

        @lru_cache(None)
        def helper(n):
            res = 1
            for i in range(n - 1, -1, -1):
                if arr[i] < arr[n]:
                    res = max(res, 1 + helper(i))
            return res
        res = 0
        for i in range(len(arr)):
            res = max(res, helper(i))
        return res


print(Solution().memoi([10, 22, 9, 33, 21, 50, 41, 60, 80]))
print(Solution().memoi([3, 10, 2, 1, 20]))
print(Solution().memoi([3, 2]))
print(Solution().memoi([50, 3, 10, 7, 40, 80]))
