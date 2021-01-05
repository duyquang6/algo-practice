# https://leetcode.com/problems/arithmetic-subarrays/

from typing import List
class Solution:
    def checkArithmeticSubarrays(self, nums: List[int], l: List[int], r: List[int]) -> List[bool]:
        res = []
        for (ql, qr) in zip(l,r):
            if qr - ql + 1 < 2:
                res.append(False)
                continue
            arr = sorted(nums[ql: qr + 1])
            leap = arr[1] - arr[0]
            for i in range(1, qr-ql):
                if arr[i+1] - arr[i] != leap:
                    res.append(False)
                    break
            else:
                res.append(True)
        return res