class FenwickTree:
    def __init__(self, arr):
        self.F = [0] * (len(arr))

        # NOTE: build fenwick O(n)
        for i in range(1, len(arr)):
            self.F[i] += arr[i]
            j = i + (i & -i)
            if j < len(arr):
                self.F[j] += self.F[i]

    def add(self, idx, X):
        while idx < len(self.F):
            self.F[idx] += X
            right_most_bit = idx & (-idx)
            idx += right_most_bit

    def sum(self, idx):
        cur = 0
        while idx > 0:
            cur += self.F[idx]
            right_most_bit = idx & (-idx)
            idx -= right_most_bit
        return cur

    def range_query(self, l, r):
        return self.sum(r) - self.sum(l - 1)


arr = [-1e9, 1, 2, 3, 4, 5]
ft = FenwickTree(arr)
print(ft.range_query(1, 3))
print(ft.range_query(2, 5))
