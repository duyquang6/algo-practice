def egg_drop(n, k):
    from functools import lru_cache

    @lru_cache(None)
    def dp(n, k):
        if n == 1:
            return k
        if k <= 0:
            return float('-inf')
        res = float('inf')
        for i in range(1, k+1):
            res = min(res, max(1 + dp(n, k-i), 1 + dp(n-1, i-1)))
        return res

    return dp(n, k)


print(egg_drop(2, 10))
