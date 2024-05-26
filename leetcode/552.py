# https://leetcode.com/problems/student-attendance-record-ii/submissions/


class Solution:
    memo = {
        (2, 0): 8,
        (2, 1): 2,
        (2, 2): 1,
        (2, 3): 1,
        (2, 4): 3,
        (2, 5): 1,
        (2, 6): 0,
        (1, 0): 3,
        (1, 1): 1,
        (1, 2): 1,
        (1, 3): 0,
        (1, 4): 1,
        (1, 5): 0,
        (1, 6): 0,
    }

    def checkRecord(self, n: int) -> int:
        # i=0: all cases
        # i=1: has 0A, 0L
        # i=2: has 0A, 1L
        # i=3: has 0A, 2L
        # i=4: has 1A, 0L
        # i=5: 1A, 1L
        # i=6: 1A, 2L
        def dp(n, i) -> int:
            if (n, i) in self.memo:
                return self.memo[(n, i)]

            if i == 0:
                res = (
                    dp(n - 1, 1) * 3
                    + dp(n - 2, 1) * 3
                    + dp(n - 1, 3) * 2
                    + dp(n - 1, 4) * 2
                    + dp(n - 2, 4) * 2
                    + dp(n - 1, 6)
                ) % (10**9 + 7)
                self.memo[(n, i)] = res
                return res
            if i == 1:
                res = (dp(n - 1, 1) + dp(n - 2, 1) + dp(n - 1, 3)) % (10**9 + 7)
                self.memo[(n, i)] = res
                return res
            if i == 3:
                res = dp(n - 2, 1) % (10**9 + 7)
                self.memo[(n, i)] = res
                return res
            if i == 4:
                res = (
                    dp(n - 1, 1)
                    + dp(n - 2, 1)
                    + dp(n - 1, 3)
                    + dp(n - 1, 4)
                    + dp(n - 2, 4)
                    + dp(n - 1, 6)
                ) % (10**9 + 7)
                self.memo[(n, i)] = res
                return res
            if i == 6:
                res = dp(n - 2, 4) % (10**9 + 7)
                self.memo[(n, i)] = res
                return res
            return 0

        return dp(n, 0) % (10**9 + 7)
