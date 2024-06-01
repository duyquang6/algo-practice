# https://leetcode.com/problems/capacity-to-ship-packages-within-d-days/description/
#
#
class Solution:
    def shipWithinDays(self, weights: List[int], days: int) -> int:
        left, right = max(weights), sum(weights)

        while left < right:
            mid = left + (right - left) // 2
            cond = True

            d = 1
            w = 0
            for i in range(len(weights)):
                w += weights[i]
                if w > mid:
                    print(w)
                    w = weights[i]
                    d += 1
                    if d > days:
                        cond = False
                        break

            if not cond:
                left = mid + 1
            else:
                right = mid

        return left
