package main

import "fmt"

func StaircaseTraversal(height int, maxSteps int) int {
	// Write your code here.
	memory := map[int]int{0: 1, 1: 1}
	return helper(height, maxSteps, memory)
}

func helper(height int, maxSteps int, memory map[int]int) int {
	if value, ok := memory[height]; ok {
		return value
	}
	for step := 1; step <= min(height, maxSteps); step++ {
		memory[height] += helper(height-step, maxSteps, memory)
	}
	return memory[height]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(StaircaseTraversal(4, 3))
}
