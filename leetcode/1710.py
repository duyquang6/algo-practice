# https://leetcode.com/problems/maximum-units-on-a-truck/

from typing import List
class Solution:
    def maximumUnits(self, boxTypes: List[List[int]], truckSize: int) -> int:
        boxTypes.sort(key=lambda x: -x[1])
        res = 0
        for p in boxTypes:
            prev = truckSize
            truckSize -= min(truckSize, p[0])
            taken = prev - truckSize
            if taken == 0:
                break
            res += taken * p[1]
        return res