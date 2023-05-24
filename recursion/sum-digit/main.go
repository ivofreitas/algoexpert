package main

import "fmt"

func sumDigitsRecursive(n int) int {
	if n < 10 {
		return n
	}
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sumDigits(sum)
}

func sumDigits(n int) int {
	for n >= 10 {
		sum := 0
		for n > 0 {
			sum += n % 10
			n /= 10
		}
		n = sum
	}
	return n
}

func main() {
	var num int
	fmt.Print("Enter an integer: ")
	fmt.Scanln(&num)
	result := sumDigits(num)
	fmt.Println("Sum of digits:", result)
}
