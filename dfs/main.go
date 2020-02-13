package main

import "log"

// node is a graph node
type Vertex struct {
	ID         int
	Neighbours map[int]*Vertex
}

func main() {
	edgeLists := [][]int{{0, 1}, {0, 2}, {2, 3}, {1, 3}, {3, 4}, {3, 5}, {5, 6}}
	log.Println(findDFS(edgeLists))
}

func findDFS(edgeLists [][]int) bool {
	graph := makeGraph(edgeLists)
	visited := map[int]bool{}
	log.Println(graph)
	return dfs(graph, 0, visited, 5)
}

func makeGraph(edgeLists [][]int) map[int][]int {
	graph := make(map[int][]int)
	for _, edge := range edgeLists {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	return graph
}

func dfs(graph map[int][]int, at int, visited map[int]bool, k int) bool {
	if _, ok := visited[at]; ok {
		return false
	}
	visited[at] = true
	if k == at {
		return true
	}
	for _, neighbour := range graph[at] {
		if dfs(graph, neighbour, visited, k) {
			return true
		}
	}
	return false
}
