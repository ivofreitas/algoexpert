package main

import "fmt"

func main() {
	fmt.Println(TwoNumberSum([]int{3, 5, -4, 8, 11, 1, -1, 6}, 10))
}

func TwoNumberSum(array []int, target int) []int {
	// Write your code here.
	duplets := make(map[int]bool, 0)

	for i := 0; i < len(array); i++ {
		y := target - array[i]
		if duplets[y] {
			return []int{array[i], y}
		}
		duplets[array[i]] = true
	}

	return nil
}
