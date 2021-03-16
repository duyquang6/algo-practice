from collections import defaultdict


class Graph:
    def __init__(self, vertices):
        self.graph = defaultdict(list)
        self.V = vertices

    def addEdge(self, u, v):
        self.graph[u].append(v)

    def topologicalSorting(self):
        pass

    def isCyclicDFS(self):
        visited = set()

        def helper(v, path):
            if v in path:
                return True
            visited.add(v)
            path.add(v)
            for neigh in self.graph[v]:
                if helper(neigh, path):
                    return True
            path.remove(v)
            return False

        for v in range(self.V):
            if helper(v, set()):
                return True
        return False

    def isCyclic(self):
        # return self.isCyclicTopo()
        return self.isCyclicDFS()


g = Graph(4)
g.addEdge(0, 1)
g.addEdge(0, 2)
g.addEdge(1, 2)
g.addEdge(2, 0)
g.addEdge(2, 3)
g.addEdge(3, 3)
print(g.isCyclic())
