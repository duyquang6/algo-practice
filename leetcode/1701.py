# https://leetcode.com/problems/average-waiting-time/
from typing import List
class Solution:
    def averageWaitingTime(self, customers: List[List[int]]) -> float:
        cur = 0
        count = 0
        for c in customers:            
            cur = c[0] if c[0] >= cur else cur
            cur += c[1]
            count += cur - c[0]
        return count / len(customers)