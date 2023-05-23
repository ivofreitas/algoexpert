package main

import "fmt"

type SpecialArray []interface{}

// Tip: Each element of `array` can either be cast
// to `SpecialArray` or `int`.
func ProductSum(array []interface{}) int {
	// Write your code here.

	return productSum(array, 1)
}

func productSum(array []interface{}, multi int) int {
	sum := 0
	for _, element := range array {
		if value, ok := element.([]interface{}); ok {
			sum += productSum(value, multi+1)
		} else {
			sum += element.(int)
		}
	}
	return sum * multi
}

func main() {
	fmt.Println(ProductSum([]interface{}{5, 2, []interface{}{7, -1}, 3, []interface{}{6, []interface{}{-13, 8}, 4}}))
}
