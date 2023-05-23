package main

import "fmt"

func GetPermutations(array []int) [][]int {
	// Write your code here.
	permutations := [][]int{}

	helper(0, array, &permutations)

	return permutations
}

func helper(i int, array []int, permutations *[][]int) {
	if i == len(array)-1 {
		newPerm := make([]int, len(array))
		copy(newPerm, array)
		*permutations = append(*permutations, newPerm)
	} else {
		for j := i; j < len(array); j++ {
			swap(array, i, j)
			helper(i+1, array, permutations)
			swap(array, i, j)
		}
	}
}

func swap(array []int, i, j int) {
	array[i], array[j] = array[j], array[i]
}

func main() {
	fmt.Println(GetPermutations([]int{1, 2, 3, 4}))
}
