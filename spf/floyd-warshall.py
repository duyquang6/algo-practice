def floyd_warshall(n, edges):
    mat = [[float("inf")] * n for _ in range(n)]
    for e in edges:
        src, dest, cost = e
        mat[src][dest] = cost
        mat[dest][src] = cost
    for i in range(n):
        for j in range(n):
            for k in range(n):
                if mat[i][j] > mat[i][k] + mat[k][j]:
                    mat[i][j] = mat[i][k] + mat[k][j]
    return mat


print(
    floyd_warshall(
        6, [[0, 1, 3], [0, 3, 1], [1, 4, 8], [1, 2, 2], [2, 5, 5], [5, 4, 3], [3, 4, 7]]
    )
)
