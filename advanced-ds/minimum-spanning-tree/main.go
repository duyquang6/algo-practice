package main

import (
	"fmt"

	"github.com/emirpasic/gods/v2/queues/priorityqueue"
)

func main() {
	testcase := [][]int{
		{0, 1, 10},
		{0, 2, 6},
		{0, 3, 5},
		{0, 3, 15},
		{2, 3, 4},
	}
	expectedMSTWeight := 19

	mst1 := primMST(4, testcase)
	mst2 := kruskalMST(4, testcase)

	check(expectedMSTWeight, mst1)
	check(expectedMSTWeight, mst2)
}

func check(expected int, edges [][]int) {
	actual := 0
	for _, e := range edges {
		actual += e[2]
	}
	if expected != actual {
		panic("expected is not actual")
	}
}

func primMST(n int, edges [][]int) [][]int {
	adjList := make([][][2]int, n)
	for i := range adjList {
		adjList[i] = make([][2]int, 0)
	}

	for _, e := range edges {
		adjList[e[0]] = append(adjList[e[0]], [2]int{e[1], e[2]})
		adjList[e[1]] = append(adjList[e[0]], [2]int{e[0], e[2]})
	}

	for i := range adjList {
		fmt.Println("node", i, "neighs", adjList[i])
	}

	visited := make([]bool, n)
	pq := priorityqueue.NewWith(func(x, y [3]int) int {
		return x[0] - y[0]
	})

	pq.Enqueue([3]int{0, -1, 0})

	res := [][]int{}
	for !pq.Empty() && n > 0 {
		pqNode, _ := pq.Dequeue()
		dist, src, node := pqNode[0], pqNode[1], pqNode[2]
		if visited[node] {
			continue
		}
		if src != -1 {
			res = append(res, []int{src, node, dist})
		}
		visited[node] = true
		n--
		for _, nei := range adjList[node] {
			neighNode, dist := nei[0], nei[1]

			if visited[neighNode] {
				continue
			}

			pq.Enqueue([3]int{dist, node, neighNode})
		}
	}

	return res
}

func kruskalMST(n int, edges [][]int) [][]int {
	// TBU
	return nil
}
