

def findWays(num_dice, num_face, target):
    from functools import lru_cache

    @lru_cache(None)
    def dp(target, d):
        if target == 0 and d == 0:
            return 1
        if d == 0:
            return 0
        res = 0
        for i in range(1, num_face + 1):
            res += dp(target - i, d-1)
        return res

    return dp(target, num_dice)


print(findWays(4, 2, 1))
print(findWays(2, 2, 3))
print(findWays(6, 3, 8))
print(findWays(4, 2, 5))
print(findWays(4, 3, 5))
