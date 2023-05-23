package main

import "fmt"

func GetNthFib(n int) int {
	//return recursiveFib(map[int]int{2: 1, 1: 0}, n)
	return iterativeFib(n)
}

// time: O(n) | space: O(1)
func iterativeFib(n int) int {
	if n < 1 {
		return 0
	}
	first, second, counter := 0, 1, 3
	for counter <= n {
		next := first + second
		first = second
		second = next
		counter++
	}

	return second
}

// time: O(n) | space: O(n)
func recursiveFib(memory map[int]int, n int) int {
	if n < 1 {
		return 0
	}
	if value, ok := memory[n]; ok {
		return value
	}
	memory[n] = recursiveFib(memory, n-1) + recursiveFib(memory, n-2)
	return memory[n]
}

func main() {
	fmt.Println(GetNthFib(8))
}
