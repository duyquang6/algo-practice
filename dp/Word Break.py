class Solution:
    def solve(self, words, input):
        words = set(words)
        from functools import lru_cache

        @lru_cache(None)
        def dp(input):
            if not input:
                return True
            for l in range(1, len(input)+1):
                if input[:l] in words and dp(input[l:]):
                    return True
            return False

        return dp(input)


print(Solution().solve(['i', 'like', 'sam', 'sung', 'samsung',
                        'mobile', 'ice', 'cream', 'icecream', 'man', 'go', 'mango'], 'ilike'))
print(Solution().solve(['i', 'like', 'sam', 'sung', 'samsung',
                        'mobile', 'ice', 'cream', 'icecream', 'man', 'go', 'mango'], 'ilikesamsung'))
