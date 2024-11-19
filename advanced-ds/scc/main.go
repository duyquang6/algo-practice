package main

import "fmt"

func main() {
	fmt.Println(
		findSCC(5, [][]int{
			{1, 3}, {1, 4}, {2, 1}, {3, 2}, {4, 5},
		}))
}

func findSCC(n int, edges [][]int) [][]int {
	return nil
}
