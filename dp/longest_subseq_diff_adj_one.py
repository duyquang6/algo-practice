class Solution:
    def memoi(self, arr):
        from functools import lru_cache

        @lru_cache(None)
        def helper(n):
            if n == 0:
                return 1
            res = 0
            adj = {arr[n] + 1, arr[n] - 1}
            for i in range(n - 1, -1, -1):
                if arr[i] in adj:
                    res = max(helper(i) + 1, res) 
            return res

        return helper(len(arr) - 1)


print(Solution().memoi([1, 2, 3, 2, 3, 7, 2, 1]))
