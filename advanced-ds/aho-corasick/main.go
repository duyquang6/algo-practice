package main

import (
	"fmt"
	"slices"

	"github.com/emirpasic/gods/v2/queues/linkedlistqueue"
)

type ahoCorasickTrie struct {
	childs      [26]*ahoCorasickTrie
	failureLink *ahoCorasickTrie
	outputs     []string
}

func search(root *ahoCorasickTrie, text string) ([][]string, [][]int) {
	outs := make([][]string, 0)
	indexes := make([][]int, 0)
	i := 0

	node := root
	for i < len(text) {
		if text[i] == ' ' {
			i++
			node = root
			continue
		}
		cidx := text[i] - 'a'
		if node.childs[cidx] != nil {
			i++
			node = node.childs[cidx]
			if len(node.outputs) > 0 {
				index := []int{}
				out := []string{}
				for _, s := range node.outputs {
					index = append(index, i-len(s))
					out = append(out, s)
				}
				outs = append(outs, out)
				indexes = append(indexes, index)
			}
		} else if node == root {
			i++
		} else {
			node = node.failureLink
		}
	}

	return outs, indexes
}

func buildAhoCorasick(patterns []string) *ahoCorasickTrie {
	root := &ahoCorasickTrie{}
	for _, s := range patterns {
		cur := root
		for _, w := range s {
			wid := w - 'a'
			if cur.childs[wid] == nil {
				cur.childs[wid] = &ahoCorasickTrie{}
			}
			cur = cur.childs[wid]
		}
		cur.outputs = []string{s}
	}

	root.failureLink = root

	q := linkedlistqueue.New[*ahoCorasickTrie]()

	for _, child := range root.childs {
		if child == nil {
			continue
		}
		child.failureLink = root
		q.Enqueue(child)
	}

	for !q.Empty() {
		for range q.Size() {
			node, _ := q.Dequeue()

			for key, child := range node.childs {
				if child == nil {
					continue
				}
				n := node.failureLink
				for n.childs[key] == nil && n != root {
					n = n.failureLink
				}
				if n.childs[key] != nil {
					child.failureLink = n.childs[key]
				} else {
					child.failureLink = root
				}
				child.outputs = copyOutput(child.outputs, child.failureLink.outputs)
				q.Enqueue(child)
			}
		}
	}

	return root
}

func addOutput(outs []string, data string) []string {
	if !slices.Contains(outs, data) {
		outs = append(outs, data)
	}
	return outs
}

func copyOutput(outs []string, ins []string) []string {
	for _, s := range ins {
		outs = addOutput(outs, s)
	}
	return outs
}

func main() {
	h := buildAhoCorasick([]string{
		"eat",
		"eating",
		"meat",
		"in",
	})
	outs, indexes := search(h, "i am eating meat")
	for i := range outs {
		for j := range outs[i] {
			fmt.Println("index:", indexes[i][j], "data:", outs[i][j])
		}
	}
}
