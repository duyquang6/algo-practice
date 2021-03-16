class Solution:
    def knightour(self, m, n):
        tours = [(-2, -1), (-1, -2), (-2, 1), (-1, 2),
                 (2, -1), (1, -2), (2, 1), (1, 2)]
        true_path = []

        def helper(path, i, j):
            if i < 0 or i >= m or j < 0 or j >= n or (i, j) in path:
                return 0
            path.add((i, j))
            if len(path) == m * n:
                return True
            for it, jt in tours:
                ni = i + it 
                nj = j + jt
                if ni < 0 or ni >= m or nj < 0 or nj >= n or (ni, nj) in path:
                    continue
                if helper(path, i + it, j + jt):
                    return True
            path.remove((i, j))
        path = set()
        for i in range(m):
            for j in range(n):
                if helper(path, i, j):
                    print(path)
                    return
                true_path = []


print(Solution().knightour(8, 8))