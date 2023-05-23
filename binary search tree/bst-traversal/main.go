package main

import "fmt"

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) InOrderTraverse(array []int) []int {
	// Write your code here.
	if tree.Left != nil {
		array = tree.Left.InOrderTraverse(array)
	}
	array = append(array, tree.Value)
	if tree.Right != nil {
		array = tree.Right.InOrderTraverse(array)
	}
	return array
}

func (tree *BST) PreOrderTraverse(array []int) []int {
	// Write your code here.
	array = append(array, tree.Value)
	if tree.Left != nil {
		array = tree.Left.PreOrderTraverse(array)
	}
	if tree.Right != nil {
		array = tree.Right.PreOrderTraverse(array)
	}
	return array
}

func (tree *BST) PostOrderTraverse(array []int) []int {
	// Write your code here.
	if tree.Left != nil {
		array = tree.Left.PostOrderTraverse(array)
	}
	if tree.Right != nil {
		array = tree.Right.PostOrderTraverse(array)
	}
	array = append(array, tree.Value)
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
	fmt.Println(root.InOrderTraverse([]int{}))
	fmt.Println(root.PreOrderTraverse([]int{}))
	fmt.Println(root.PostOrderTraverse([]int{}))
}
