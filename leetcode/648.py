# https://leetcode.com/problems/replace-words/

from typing import List

class TreeNode:
    def __init__(self):
        self.children = {}
        self.is_end = False

class Trie:
    def __init__(self):
        self.root = TreeNode()
    
    def insert(self, s):
        node = self.root
        for char in s:
            if char not in node.children:
                node.children[char] = TreeNode()
            node = node.children[char]
        node.is_end = True
    
    def searchPrefix(self, s):
        node = self.root
        for i, char in enumerate(s):
            if char not in node.children:
                return
            node = node.children[char]
            if node.is_end:
                return s[:i+1]
class Solution:
    def replaceWords(self, dictionary: List[str], sentence: str) -> str:
        trie = Trie()
        for word in dictionary:
            trie.insert(word)
        
        sentence = sentence.split(' ')
        
        for i, word in enumerate(sentence):
            ans = trie.searchPrefix(word)
            if ans:
                sentence[i] = ans
        return ' '.join(sentence)
        