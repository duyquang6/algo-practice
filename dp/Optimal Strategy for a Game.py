
class Solution:
    def solution1(self, arr):
        from functools import lru_cache
        @lru_cache(None)
        def dp(i, j):
            if i > j:
                return float('inf')
            if i == j:
                return arr[i]
            if j - i == 1:
                return max(arr[i], arr[j])
            return max(arr[j] + min(dp(i, j-2), dp(i+1, j-1)), arr[i] + min(dp(i+2, j), dp(i+1, j-1)))
        return dp(0, len(arr)-1)


print(Solution().solution1([5, 3, 7, 10]))
print(Solution().solution1([8, 15, 3, 7]))
