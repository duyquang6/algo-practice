class Node:
    def __init__(self, start, end, val=0, left=None, right=None) -> None:
        self.left = left
        self.right = right
        self.start = start
        self.end = end
        self.val = val

    def print(self) -> None:
        def print_tree(node):
            print(f"[{node.start}-{node.end}]: {node.val}")
            if node.left:
                print_tree(node.left)
            if node.right:
                print_tree(node.right)

        print_tree(self)


class Solution:
    def buildSegmentTree(self, arr) -> Node:
        N = len(arr)

        def build(start, end) -> Node:
            if end < start:
                return None

            node = Node(
                start,
                end,
            )

            if start == end:
                node.val = arr[start]
                return node

            n = end + start
            mid = n // 2

            left_start = start
            left_end = mid
            node.left = build(left_start, left_end)

            right_start = mid + 1
            right_end = end
            node.right = build(right_start, right_end)

            if node.left:
                node.val += node.left.val

            if node.right:
                node.val += node.right.val

            return node

        return build(0, N - 1)

    def updateIndex(self, tree, index, value) -> None:
        def recur(node):
            start, end = node.start, node.end
            if index < start or index > end:
                return 0
            if start == end and start == index:
                diff = value - node.val
                node.val = value
                return diff
            mid = (end + start) // 2
            diff = 0
            if index <= mid:
                diff = recur(node.left)
            else:
                diff = recur(node.right)

            node.val += diff
            return diff

        recur(tree)


tree = Solution().buildSegmentTree([0, 1, 2, 3, 5, 6, 7])
tree.print()
Solution().updateIndex(tree, 1, 10)
print("After update")
tree.print()
