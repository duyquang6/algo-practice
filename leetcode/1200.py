# https://leetcode.com/problems/minimum-absolute-difference/

from typing import List
class Solution:
    def minimumAbsDifference(self, arr: List[int]) -> List[List[int]]:
        arr.sort()
        diffs = []
        _min = float('inf')
        for i in range(len(arr) - 1):
            if arr[i+1] - arr[i] < _min:
                diffs.clear()
                _min = arr[i+1] - arr[i]
            if arr[i+1] - arr[i] == _min:
                diffs.append(i)
            
        return [[arr[i], arr[i+1]] for i in diffs]