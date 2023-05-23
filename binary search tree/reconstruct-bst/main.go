package main

// This is an input class. Do not edit.
type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func ReconstructBst(preOrderTraversalValues []int) *BST {
	if len(preOrderTraversalValues) == 0 {
		return nil
	}

	current := preOrderTraversalValues[0]
	rightTreeIdx := len(preOrderTraversalValues)

	for i := 1; i < len(preOrderTraversalValues); i++ {
		if preOrderTraversalValues[i] >= current {
			rightTreeIdx = i
			break
		}
	}

	leftNode := ReconstructBst(preOrderTraversalValues[1:rightTreeIdx])
	rightNode := ReconstructBst(preOrderTraversalValues[rightTreeIdx:])

	return &BST{current, leftNode, rightNode}
}

type TreeInfo struct {
	RootIdx int
}

func reconstructBstWithRange(lowerBound, upperBound int, preOrderTraversalValues []int, currentSubtreeInfo *TreeInfo) *BST {
	if len(preOrderTraversalValues) == currentSubtreeInfo.RootIdx {
		return nil
	}

	current := preOrderTraversalValues[currentSubtreeInfo.RootIdx]
	if current < lowerBound || current >= upperBound {
		return nil
	}

	currentSubtreeInfo.RootIdx++
	leftNode := reconstructBstWithRange(lowerBound, current, preOrderTraversalValues, currentSubtreeInfo)
	rightNode := reconstructBstWithRange(current, upperBound, preOrderTraversalValues, currentSubtreeInfo)

	return &BST{current, leftNode, rightNode}
}

func main() {}
