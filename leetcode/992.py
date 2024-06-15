# https://leetcode.com/problems/subarrays-with-k-different-integers/
from typing import List


class Solution:
    def subarraysWithKDistinct(self, nums: List[int], k: int) -> int:
        diff = {}
        n = len(nums)
        res = 0
        left = 0
        right = 0
        next_left = None
        for right in range(n):
            val = nums[right]
            diff[val] = right
            if len(diff) < k:
                continue
            while len(diff) > k:
                if diff[nums[left]] == left:
                    del diff[nums[left]]
                left += 1
            if next_left is None:
                next_left = left
            next_left = max(next_left, left)
            while next_left < n and diff[nums[next_left]] != next_left:
                next_left += 1
            res += next_left - left + 1

        return res


print(Solution().subarraysWithKDistinct([1, 2, 1, 2, 3], 2))
print(Solution().subarraysWithKDistinct([1, 2, 1, 3, 4], 3))
