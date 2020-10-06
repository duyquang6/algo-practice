# https://leetcode.com/problems/special-array-with-x-elements-greater-than-or-equal-x/
from typing import List
class Solution:
    def specialArray(self, nums: List[int]) -> int:
        counter = [0] * 101
        for l in nums:
            if l >= 100:
                counter[100] += 1
            else:
                counter[l] += 1
        start = max(nums)
        if start > 100:
            start = 100
        count = 0
        while start > 0:
            count += counter[start]
            if count == start:
                return start
            if count > start:
                return -1
            start -= 1
        return -1
            
import unittest
class SolutionTest(unittest.TestCase):
    def test_specialArray(self):
        self.assertEqual(Solution().specialArray([3,5]), 2)
        self.assertEqual(Solution().specialArray([0,4,3,0,4]), 3)
        self.assertEqual(Solution().specialArray([3,6,7,7,0]), -1)
        self.assertEqual(Solution().specialArray([0,0]), -1)

if __name__ == '__main__':
    unittest.main()