

class Solution:
    def solve(self, str1, str2):
        from functools import lru_cache

        @lru_cache(None)
        def dp(m, n):
            if m == 0 and n == 0:
                return 0
            if m == 0 or n == 0:
                return float('inf')            
            if str1[m-1] == str2[n-1]:
                return dp(m-1, n-1)
            return min(dp(m, n-1) + 1, dp(m-1, n-1) + 1, dp(m-1, n) + 1)
        return dp(len(str1), len(str2))


print(Solution().solve("sunday", "saturday"))
print(Solution().solve("cat", "cut"))
print(Solution().solve("geek", "gesek"))
