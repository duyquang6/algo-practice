class Solution:
    def solve(self, arr):
        from functools import lru_cache

        def dp(i, j):
            if j - i == 2:
                return arr[i] * arr[i+1] * arr[j]
            return min((arr[i] * arr[k] * arr[j] + dp(i, k) + dp(k, j) for k in range(i+1, j)), default=0)
        return dp(0, len(arr)-1)


print(Solution().solve([40, 20, 30, 10, 30]))
print(Solution().solve([10, 20, 30, 40, 30]))
print(Solution().solve([10, 20, 30]))
