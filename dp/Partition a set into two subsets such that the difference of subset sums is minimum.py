class Solution:
    def solution1(self, arr):
        from functools import lru_cache
        if not arr:
            return None
        _sum = sum(arr)

        @lru_cache(None)
        def dp(n, weight):
            if n == 1:
                return arr[n-1] if arr[n-1] <= weight else 0
            res = 0
            for i in range(1, n):
                if arr[n-1] <= weight:
                    res = max(res, arr[n-1] + dp(i-1, weight-arr[n-1]))
            return res
        res = 0
        for i in range(1, len(arr)+1):
            res = max(res, dp(i, _sum / 2))
        return _sum - 2 * res

    def solution2(self, arr):
        from functools import lru_cache
        if not arr:
            return None
        _sum = sum(arr)

        @lru_cache(None)
        def dp(n, weight):
            if n == 1:
                return arr[n-1] if arr[n-1] <= weight else 0
            res = 0
            for i in range(1, n):
                if arr[n-1] <= weight:
                    res = max(res, arr[n-1] + dp(i-1, weight-arr[n-1]))
            return res
        res = 0
        for i in range(1, len(arr)+1):
            res = max(res, dp(i, _sum / 2))
        return _sum - 2 * res


print(Solution().solution1([1, 6, 11, 5]))
