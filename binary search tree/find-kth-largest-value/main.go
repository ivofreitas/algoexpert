package main

import "fmt"

// This is an input class. Do not edit.
type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func FindKthLargestValueInBst(tree *BST, k int) int {
	// Write your code here.
	return tree.helper(k, []int{})[k-1]
}

func (tree *BST) helper(k int, array []int) []int {

	if len(array) < k {
		if tree.Right != nil {
			array = tree.Right.helper(k, array)
		}
		array = append(array, tree.Value)
		if tree.Left != nil {
			array = tree.Left.helper(k, array)
		}
	}

	return array
}

func main() {
	root := &BST{Value: 15}
	root.Left = &BST{Value: 5}
	root.Left.Left = &BST{Value: 2}
	root.Left.Left.Left = &BST{Value: 1}
	root.Left.Left.Right = &BST{Value: 3}
	root.Left.Right = &BST{Value: 5}
	root.Right = &BST{Value: 20}
	root.Right.Left = &BST{Value: 17}
	root.Right.Right = &BST{Value: 22}
	fmt.Println(FindKthLargestValueInBst(root, 3))
}
