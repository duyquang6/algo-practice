# https://leetcode.com/problems/where-will-the-ball-fall/
import functools

from typing import List


class Solution:
    def findBall(self, grid: List[List[int]]) -> List[int]:
        m = len(grid)
        n = len(grid[0])

        @functools.lru_cache(None)
        def dp(i, j):
            nonlocal m, n
            if i >= m:
                return j
            prev = grid[i][j]
            if not 0 <= j + prev < n:
                return -1
            neighbor = grid[i][j+prev]
            return dp(i + 1, j + prev) if neighbor == prev else -1                

        return [dp(0, i) for i in range(n)]
