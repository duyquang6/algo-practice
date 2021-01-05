# https://leetcode.com/problems/map-sum-pairs/

class TreeNode:
    def __init__(self):
        self.children = {}
        self.value = 0
class MapSum:

    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.root = TreeNode()
        self.map = {}

    def insert(self, key: str, val: int) -> None:
        diff = val
        if key in self.map:
            diff = val - self.map[key]
        self.map[key] = val
        node = self.root
        for char in key:
            if char not in node.children:
                node.children[char] = TreeNode()
            node.value += diff
            node = node.children[char]
        node.value += diff

    def sum(self, prefix: str) -> int:
        node = self.root
        for char in prefix:
            if char not in node.children:
                return 0
            node = node.children[char]
        return node.value


# Your MapSum object will be instantiated and called as such:
# obj = MapSum()
# obj.insert(key,val)
# param_2 = obj.sum(prefix)