class Solution:
    def memoi(self, s):
        from functools import lru_cache
        res = ''
        @lru_cache(None)
        def helper(m, n, data):
            nonlocal res
            if m == -1 or n == -1:
                return 0
            if m == n:
                return max(helper(m - 1, n, data), helper(m, n-1, data))
            if s[m] == s[n]:
                data += s[m]
                res = max(res, data)
                return 1 + helper(m - 2, n - 2, data)
            else:
                return max(helper(m - 1, n, data), helper(m, n - 1, data))
        
        helper(len(s) - 2, len(s) - 1, '')
        return res[::-1]


print(Solution().memoi("AABEBCDD"))
