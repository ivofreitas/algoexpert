package main

// look for the middle value until there is no value left
func MinHeightBST(array []int) *BST {
	// Write your code here.

	return constructMinHeightBST(array, 0, len(array)-1)
}

func constructMinHeightBST(array []int, startIdx, endIdx int) *BST {
	if endIdx < startIdx {
		return nil
	}
	midIdx := (startIdx + endIdx) / 2
	bst := &BST{Value: array[midIdx]}

	bst.Left = constructMinHeightBST(array, startIdx, midIdx-1)
	bst.Right = constructMinHeightBST(array, midIdx+1, endIdx)

	return bst
}

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) Insert(value int) *BST {
	if value < tree.Value {
		if tree.Left == nil {
			tree.Left = &BST{Value: value}
		} else {
			tree.Left.Insert(value)
		}
	} else {
		if tree.Right == nil {
			tree.Right = &BST{Value: value}
		} else {
			tree.Right.Insert(value)
		}
	}
	return tree
}

func main() {
}
