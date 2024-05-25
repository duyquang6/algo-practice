# https://leetcode.com/problems/course-schedule-ii/
#

from typing import List


class Solution:
    def findOrder(self, numCourses: int, prerequisites: List[List[int]]) -> List[int]:
        adj = {}
        for edge in prerequisites:
            u, v = edge[0], edge[1]
            if v not in adj:
                adj[v] = []
            adj[v].append(u)

        stack = []

        def dfs(visited, path, u):
            if u in visited:
                return True
            path.add(u)
            visited.add(u)

            for neigh in adj.get(u, []):
                if neigh in path:
                    # loop
                    return False
                if neigh in visited:
                    continue
                if not dfs(visited, path, neigh):
                    return False
            stack.append(u)
            path.remove(u)

        visited = set()

        for n in range(numCourses):
            if len(visited) == numCourses:
                return stack[::-1]
            if not dfs(visited, set(), n):
                return []
        return stack[::-1]


Solution().findOrder(2, [[1, 0]])
