def solve(arr):
    from functools import lru_cache
    arr_3 = []
    for d in arr:
        a = sorted(d)
        arr_3.append((a[0], a[1], a[2]))
        arr_3.append((a[1], a[2], a[0]))
        arr_3.append((a[0], a[2], a[1]))
    arr_3.sort(key=lambda x: x[0] * x[1])
    print(arr_3)

    @lru_cache(None)
    def dp(n):
        res = arr_3[n-1][2]
        for k in range(n-1, 0, -1):
            if arr_3[k-1][0] < arr_3[n-1][0] and arr_3[k-1][1] < arr_3[n-1][1]:
                res = max(res, arr_3[n-1][2] + dp(k))
        return res
    return max((dp(l) for l in range(1, len(arr_3)+1)), default=0)


print(solve([[4, 6, 7], [1, 2, 3], [4, 5, 6], [10, 12, 32]]))
