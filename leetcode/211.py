# https://leetcode.com/problems/design-add-and-search-words-data-structure/


class TreeNode:
    def __init__(self):
        self.children = {}
        self.is_end = False

class WordDictionary:

    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.root = TreeNode()

    def addWord(self, word: str) -> None:
        node = self.root
        for char in word:
            if char not in node.children:
                node.children[char] = TreeNode()
            node = node.children[char]
        node.is_end = True

    def search(self, word: str) -> bool:
        node = self.root
        def recur(node, word):
            for i, char in enumerate(word):
                if char == '.':
                    return any(recur(node, k + word[i+1:] if len(word) > i + 1 else k) for k in node.children.keys())
                if char not in node.children:
                    return False
                node = node.children[char]
            return node.is_end
        return recur(node, word)


# Your WordDictionary object will be instantiated and called as such:
# obj = WordDictionary()
# obj.addWord(word)
# param_2 = obj.search(word)