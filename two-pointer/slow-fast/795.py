# https://leetcode.com/problems/number-of-subarrays-with-bounded-maximum/
class Solution:
    def numSubarrayBoundedMax(self, nums: List[int], left: int, right: int) -> int:
        res = 0
        i = 0
        dp = 0
        for j in range(len(nums)):
            if nums[j] < left:
                res += dp
                continue
            if nums[j] > right:
                i = j + 1
                dp = 0
                continue
            dp = j - i + 1
            res += dp
        return res
