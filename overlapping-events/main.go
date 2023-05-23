package main

import (
	"fmt"
	"sort"
)

func OverlappingEvents(events [][]int) [][]int {

	sort.Slice(events[:], func(i, j int) bool {
		for x := range events[i] {
			if events[i][x] == events[j][x] {
				continue
			}
			return events[i][x] < events[j][x]
		}
		return false
	})

	result := make([][]int, 0)
	result = append(result, events[0])

	for i := 1; i < len(events); i++ {
		current := events[i]
		predecessor := result[len(result)-1]
		if predecessor[1] >= current[0] {
			limit := Max(predecessor[1], current[1])
			result[len(result)-1] = []int{predecessor[0], limit}
		} else {
			result = append(result, events[i])
		}
	}

	return result
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x

}

func main() {
	fmt.Println(OverlappingEvents([][]int{{1, 2}, {3, 5}, {4, 7}, {6, 8}, {9, 10}}))
}
