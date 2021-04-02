class Solution:
    def solve(self, arr, target):
        arr.sort()
        i, j = 0, 0
        _sum = arr[0]
        while j < len(arr):
            if _sum == target:
                return True
            elif _sum < target:
                j += 1
                if j < len(arr):
                    _sum += arr[j]
            else:
                _sum -= arr[i]
                i += 1
        return False

    def solution2(self, arr, target):
        from functools import lru_cache

        @lru_cache(None)
        def dp(n, _sum):
            if _sum == 0:
                return True
            if n == 1:
                return arr[0] == _sum
            if _sum < 0:
                return False
            if _sum == arr[n-1]:
                return True
            for i in range(1, n):
                if dp(i, _sum-arr[n-1]):
                    return True
            return False
        for i in range(1, len(arr)+1):
            if dp(i, target):
                return True
        return False


print(Solution().solve([3, 34, 4, 12, 5, 2], 9))
print(Solution().solve([3, 34, 4, 12, 5, 2], 30))

print(Solution().solution2([3, 34, 4, 12, 5, 2], 9))
print(Solution().solution2([3, 34, 4, 12, 5, 2], 30))
