package main

import "fmt"

type Trie struct {
	prefix      string
	sub         []*Trie
	isEndOfWord bool
}

func compressTrie(strs []string) Trie {
	root := Trie{sub: make([]*Trie, 26)}
	for _, s := range strs {
		build(s, &root)
	}

	return root
}

func build(s string, root *Trie) {
	i := 0
	cur := root
	for i < len(s) {
		cidx := s[i] - 'a'
		parent := cur
		cur = cur.sub[cidx]
		if cur == nil {
			newNode := Trie{
				prefix:      s[i:],
				sub:         make([]*Trie, 26),
				isEndOfWord: true,
			}
			parent.sub[cidx] = &newNode
			return
		}
		commonLength := getLongestCommonPrefix(cur.prefix, s[i:])
		i += commonLength

		if commonLength < len(cur.prefix) {
			newNode := Trie{
				prefix: cur.prefix[:commonLength],
				sub:    make([]*Trie, 26),
			}

			newNode.sub[cur.prefix[commonLength]-'a'] = cur
			cur.prefix = cur.prefix[commonLength:]
			parent.sub[cidx] = &newNode
		}
		if i == len(s) {
			cur.isEndOfWord = true
		}
	}
}

func getLongestCommonPrefix(s1, s2 string) int {
	length := min(len(s1), len(s2))
	for i := range length {
		if s1[i] == s2[i] {
			continue
		}
		return i
	}

	return length
}

func print(root *Trie, level int) {
	fmt.Printf("level:%d,prefix:%s,isEndOfWord:%t\n", level, root.prefix, root.isEndOfWord)
	for i := range root.sub {
		if root.sub[i] != nil {
			print(root.sub[i], level+1)
		}
	}
}

func main() {
	root := compressTrie([]string{"book", "boar"})
	print(&root, 0)
}
