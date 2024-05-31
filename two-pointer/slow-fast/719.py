# https://leetcode.com/problems/find-k-th-smallest-pair-distance/
class Solution:
    def smallestDistancePair(self, nums: List[int], k: int) -> int:
        def cond(distance):
            count = 0
            i, j = 0, 0
            n = len(nums)
            while i < n or j < n:
                while j < n and nums[j] - nums[i] <= distance:
                    j += 1
                count += j - i - 1
                i += 1
            return count >= k

        nums.sort()
        left, right = 0, nums[-1] - nums[0]
        while left < right:
            mid = left + (right - left) // 2
            if not cond(mid):
                left = mid + 1
            else:
                right = mid
        return left
