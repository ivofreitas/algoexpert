package main

import "fmt"

func Powerset(array []int) [][]int {
	// Write your code here.
	powerSet := [][]int{{}}
	powerSetIdx := 0
	powerSetSize := len(powerSet) - 1
	for _, element := range array {
		for powerSetIdx <= powerSetSize {
			newArray := []int{element}
			newArray = append(newArray, powerSet[powerSetIdx]...)
			powerSet = append(powerSet, newArray)
			powerSetIdx++
		}
		powerSetIdx = 0
		powerSetSize = len(powerSet) - 1
	}

	return powerSet
}

func powersetRecursive(array []int, index int) [][]int {
	if index < 0 {
		return [][]int{{}}
	}
	subsets := powersetRecursive(array, index-1)
	element := array[index]
	length := len(subsets)
	for i := 0; i < length; i++ {
		currentSubset := subsets[i]
		newsubset := append([]int{}, currentSubset...)
		newsubset = append(newsubset, element)
		subsets = append(subsets, newsubset)
	}
	return subsets
}

func main() {
	fmt.Println(powersetRecursive([]int{1, 2, 3}, 2))
}
