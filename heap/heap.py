def recur(arr, i):
    n = len(arr)
    if i >= n:
        return False
    left = 2 * i + 1
    right = 2 * i + 2
    is_left = recur(arr, left)
    is_right = recur(arr, right)
    if not is_left and not is_right:
        return True
    min_i = left

    if is_right and arr[right] < arr[left]:
        min_i = right

    if arr[min_i] < arr[i]:
        arr[i], arr[min_i] = arr[min_i], arr[i]
    return True


class MinHeap:
    def __init__(self, arr=[]):
        self.heapify(arr)
        self.arr = arr

    def heapify(self, arr):
        n = len(arr)

        for i in range(n):
            recur(arr, i)

    def pop(self):
        arr = self.arr
        n = len(arr)
        if n == 0:
            return None
        val = arr[0]
        arr[0], arr[n - 1] = arr[n - 1], arr[0]
        arr.pop()
        recur(arr, 0)

        return val

    def push(self, val):
        self.arr.append(val)
        recur(self.arr, 0)


class MaxHeap:
    def __init__(self, arr=[]):
        for i in range(len(arr)):
            arr[i] = -arr[i]
        self.heapify(arr)
        self.arr = arr

    def heapify(self, arr):
        n = len(arr)

        for i in range(n):
            recur(arr, i)

    def pop(self):
        arr = self.arr
        n = len(arr)
        if n == 0:
            return None
        val = arr[0]
        arr[0], arr[n - 1] = arr[n - 1], arr[0]
        arr.pop()
        recur(arr, 0)

        return -val

    def push(self, val):
        self.arr.append(-val)
        recur(self.arr, 0)


h = [100, 41, 51, 41, 31, 16, 13]
heap = MinHeap(h)
print(h)

pop_val = heap.pop()
print("pop_val", pop_val)
print(h)

pop_val = heap.pop()
print("pop_val", pop_val)
print(h)

heap.push(10)
print("after_push", h)


print("-------------Max Heap--------------")

h = [100, 41, 51, 41, 31, 16, 13]
heap = MaxHeap(h)
print(h)

pop_val = heap.pop()
print("pop_val", pop_val)
print(h)

pop_val = heap.pop()
print("pop_val", pop_val)
print(h)

heap.push(10)
print("after_push", h)
