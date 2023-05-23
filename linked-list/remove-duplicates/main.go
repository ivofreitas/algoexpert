package main

import "fmt"

// This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

func RemoveDuplicatesFromLinkedList(linkedList *LinkedList) *LinkedList {
	// Write your code here.
	for linkedList.Next != nil {
		next := linkedList.Next
		if linkedList.Value == next.Value {
			linkedList.Next = next.Next
		} else {
			linkedList = linkedList.Next
		}
	}

	return linkedList
}

func addMany(linkedList *LinkedList, values []int) *LinkedList {
	current := linkedList
	for current.Next != nil {
		current = current.Next
	}
	for _, value := range values {
		current.Next = &LinkedList{Value: value}
		current = current.Next
	}
	return linkedList
}

func getValues(linkedList *LinkedList) []int {
	values := []int{}
	current := linkedList
	for current != nil {
		values = append(values, current.Value)
		current = current.Next
	}
	return values
}

func main() {
	input := addMany(&LinkedList{Value: 1}, []int{1, 3, 4, 4, 4, 5, 6, 6})
	RemoveDuplicatesFromLinkedList(input)
	fmt.Println(getValues(input))
}
