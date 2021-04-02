

class Solution:
    def solve(self, mat):
        from functools import lru_cache
        neighbors = [[0, 1], [0, -1], [1, 0], [-1, 0]]
        M, N = len(mat), len(mat[0])

        @lru_cache(None)
        def dp(i, j):
            res = 1
            val = mat[i][j]
            for neigh in neighbors:
                next_i, next_j = i + neigh[0], j + neigh[1]
                if not (0 <= next_i < M) or not (0 <= next_j < N) or mat[next_i][next_j] != val + 1:
                    continue
                res = max(res, 1 + dp(next_i, next_j))
            return res
        res = 0
        for i in range(M):
            for j in range(N):
                res = max(res, dp(i, j))
        return res


print(Solution().solve([[1, 2, 9],
                        [5, 3, 8],
                        [4, 6, 7]]))
