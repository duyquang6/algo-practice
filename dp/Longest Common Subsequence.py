class Solution:
    def solve(self, str1, str2):

        from functools import lru_cache

        @lru_cache(None)
        def dp(m, n):
            if m == 0 or n == 0:
                return 0

            if str1[m-1] == str2[n-1]:
                return dp(m-1, n-1) + 1

            return max(dp(m-1, n), dp(m, n-1))

        return dp(len(str1), len(str2))


print(Solution().solve("AGGTAB", "GXTXAYB"))
