package main

import "fmt"

func MinRewards(scores []int) int {
	minReward := 0
	allRewards := initReward(len(scores))

	fmt.Printf("scores: %v\n", scores)

	for i := 0; i < len(scores)-1; i++ {
		if scores[i] > scores[i+1] {
			continue
		}

		allRewards[i+1] = allRewards[i] + 1
	}
	fmt.Printf("foward change: %v\n", allRewards)

	for j := len(scores) - 1; j > 0; j-- {
		if scores[j] > scores[j-1] {
			continue
		}

		allRewards[j-1] = max(allRewards[j-1], allRewards[j]+1)
	}
	fmt.Printf("backward change: %v\n", allRewards)

	for _, reward := range allRewards {
		minReward += reward
	}

	return minReward
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func initReward(size int) []int {
	result := make([]int, size)
	for idx := range result {
		result[idx] = 1
	}
	return result
}

func main() {
	fmt.Printf("result: %v", MinRewards([]int{8, 4, 2, 1, 3, 6, 7, 9, 5}))
}
