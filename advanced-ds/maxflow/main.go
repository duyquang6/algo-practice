package main

import (
	"fmt"
	"math"

	llq "github.com/emirpasic/gods/v2/queues/linkedlistqueue"
)

func maxflow(s, t int, n int, graph [][]int) int {
	res := 0
	parents := make([]int, n)
	for i := range parents {
		parents[i] = -1
	}

	newFlow := bfs(graph, s, t, parents)
	for newFlow != 0 {
		res += newFlow
		cur := t

		for cur != s {
			p := parents[cur]
			graph[p][cur] -= newFlow
			graph[cur][p] += newFlow
			cur = p
		}

		parents = make([]int, n)
		for i := range parents {
			parents[i] = -1
		}

		newFlow = bfs(graph, s, t, parents)
	}

	return res
}

func bfs(graph [][]int, s, t int, parents []int) int {
	q := llq.New[int]()
	q.Enqueue(s)

	found := false
	for !q.Empty() {
		node, _ := q.Dequeue()
		if node == t {
			found = true
			break
		}
		for neinode := range graph[node] {
			if parents[neinode] != -1 || graph[node][neinode] == 0 {
				continue
			}
			parents[neinode] = node
			q.Enqueue(neinode)
		}
	}

	if !found {
		return 0
	}

	cur := t
	res := math.MaxInt
	for cur != s {
		prev := parents[cur]
		res = min(res, graph[prev][cur])
		cur = prev
	}

	return res
}

func main() {
	fmt.Println(
		maxflow(0, 5, 6, [][]int{
			{0, 16, 13, 0, 0, 0},
			{0, 0, 10, 12, 0, 0},
			{0, 4, 0, 0, 14, 0},
			{0, 0, 9, 0, 0, 20},
			{0, 0, 0, 7, 0, 4},
			{0, 0, 0, 0, 0, 0},
		}),
	)
}
