package main

import (
	"github.com/duyquang6/go-bitmap"
)

var visited = bitmap.NewBitmap(100)

func main() {
	adjList := [][]int{
		{},
		{0},
		{0},
		{0},
		{1, 2},
		{1, 3},
		{2, 3},
		{4, 5, 6},
	}

	// do toposort to check if there is cyclic
	if checkCyclic(adjList) {
		panic("cannot do c3 algo due to cyclic")
	}

}

func mro(adjList [][]int, node int) {
	neighbors := adjList[node]

	return
}

func merge() {

}

func checkCyclic(adjList [][]int) bool {
	for node := range adjList {
		ok, _ := visited.Get(uint64(node))
		if ok {
			continue
		}
		if !dfs(adjList, node) {
			return false
		}
	}
	return true
}

func dfs(adjList [][]int, node int) bool {
	visited.Set(uint64(node), true)
	for _, val := range adjList[node] {
		ok, _ := visited.Get(uint64(val))
		if ok {
			return false
		}
		if !dfs(adjList, val) {
			return false
		}
	}
	return true
}
