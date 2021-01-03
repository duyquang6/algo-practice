# https://leetcode.com/problems/widest-vertical-area-between-two-points-containing-no-points/

class Solution:
    def maxWidthOfVerticalArea(self, points: List[List[int]]) -> int:
        points.sort(key=lambda x: x[0])
        diff = 0
        for i in range(len(points)-1):
            diff = max(diff, points[i+1][0] - points[i][0])
        return diff