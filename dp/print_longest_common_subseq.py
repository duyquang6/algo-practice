
class Solution:
    def solve(self, str1, str2):

        from functools import lru_cache
        @lru_cache(None)
        def dp(m, n):
            res = set()
            if m == 0 or n == 0:
                return res
            if str1[m-1] == str2[n-1]:
                data = dp(m-1,n-1)
                if data:
                    for i in data:
                        res.add(i + str1[m-1])
                else:
                    res.add(str1[m-1])
            else:
                max_m = max((len(i) for i in dp(m-1, n)), default=0)
                max_n = max((len(i) for i in dp(m, n-1)), default=0)
                if max_m >= max_n:
                    res = dp(m-1, n)
                if max_n >= max_m:
                    tmp = dp(m, n-1)
                    res.update(tmp)
            return res
            
        return dp(len(str1), len(str2))

print(Solution().solve("AGTGATG", "GTTAG"))
print(Solution().solve("AATCC", "ACACG"))
print(Solution().solve("ABCBDAB", "BDCABA"))
