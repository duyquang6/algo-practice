class Solution:
    def solve(self, arr):
        from functools import lru_cache
        _sum = sum(arr)

        @lru_cache(None)
        def dp(n, _cur):
            nonlocal _sum
            if _cur == _sum - _cur:
                return True
            if n == 0:
                return False

            return dp(n-1, _cur) or dp(n-1, _cur + arr[n-1])

        return any(dp(k, 0) for k in range(1, len(arr)+1))


print(Solution().solve([1,1]))
