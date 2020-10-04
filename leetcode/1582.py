# https://leetcode.com/problems/special-positions-in-a-binary-matrix/
from typing import List
class Solution:
    def numSpecial(self, mat: List[List[int]]) -> int:
        t = [0] * len(mat[0])
        t1 = [0] * len(mat)
        for i,v1 in enumerate(mat):
            for j in range(len(v1)):
                t[j] += v1[j]
        ans = 0
        for i,v1 in enumerate(mat):
            t1[i] = sum(v1)
        for i, v in enumerate(t):
            if v == 1:
                for j in range(len(mat)):
                    if mat[j][i] == 1 and t1[j] == 1:
                         ans+=1
        return ans
            
import unittest

class SolutionTest(unittest.TestCase):
    def test_numSpecial(self):
        self.assertEqual(Solution().numSpecial([[0,0,0,0,0],[1,0,0,0,0],[0,1,0,0,0],[0,0,1,0,0],[0,0,0,1,1]]), 3)
        self.assertEqual(Solution().numSpecial([[1,0,0],[0,0,1],[1,0,0]]), 1)
        self.assertEqual(Solution().numSpecial([[0,0,0,1],[1,0,0,0],[0,1,1,0],[0,0,0,0]]), 2)
        self.assertEqual(Solution().numSpecial([[0,0,1,0],[0,0,0,0],[0,0,0,0],[0,1,0,0]]), 2)

if __name__ == '__main__':
    unittest.main()