package main

func SameBsts(arrayOne, arrayTwo []int) bool {
	if len(arrayOne) != len(arrayTwo) {
		return false
	}

	if len(arrayOne) == 0 {
		return true
	}

	if arrayOne[0] != arrayTwo[0] {
		return false
	}

	leftOne := smallerValues(arrayOne)
	leftTwo := smallerValues(arrayTwo)
	rightOne := greaterOrEqual(arrayOne)
	rightTwo := greaterOrEqual(arrayTwo)

	return SameBsts(leftOne, leftTwo) && SameBsts(rightOne, rightTwo)
}

func smallerValues(array []int) []int {
	var result []int
	for i := 1; i < len(array); i++ {
		if array[i] < array[0] {
			result = append(result, array[i])
		}
	}
	return result
}

func greaterOrEqual(array []int) []int {
	var result []int
	for i := 1; i < len(array); i++ {
		if array[i] >= array[0] {
			result = append(result, array[i])
		}
	}
	return result
}

func main() {}
