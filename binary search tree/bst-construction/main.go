package main

/*
	BINARY SEARCH TREE DIFFERS FROM BINARY TREE ON A ADDITIONAL PROPERTY
	CONSIDER A ROOT NODE THAT HAS A VALUE AND A POINTER TO A LEFT AND A POINTER TO A RIGHT NODE
	> ALL VALUES TO IT'S LEFT ARE SMALLER
	> ALL VALUES TO IT'S RIGHT ARE EQUAL OR GREATER
	> ALL SUB NODES RESPECT THESE RULES
*/

type BST struct {
	Value int
	Left  *BST
	Right *BST
}

func (tree *BST) Insert(value int) *BST {
	node := new(BST)
	node.Value = value

	currentNode := tree

	for {
		if value < currentNode.Value {
			if currentNode.Left == nil {
				currentNode.Left = node
				break
			} else {
				currentNode = currentNode.Left
			}
		} else {
			if currentNode.Right == nil {
				currentNode.Right = node
				break
			} else {
				currentNode = currentNode.Right
			}
		}
	}

	return tree
}

func (tree *BST) Contains(value int) bool {

	currentNode := tree

	for currentNode != nil {
		if value < currentNode.Value {
			currentNode = currentNode.Left
		} else if value > currentNode.Value {
			currentNode = currentNode.Right
		} else {
			return true
		}
	}

	return false
}

func (tree *BST) Remove(value int) *BST {
	tree.remove(value, nil)
	return tree
}

func (tree *BST) remove(value int, parent *BST) *BST {
	currentNode := tree
	for currentNode != nil {
		if value < currentNode.Value {
			parent = currentNode
			currentNode = currentNode.Left
		} else if value > currentNode.Value {
			parent = currentNode
			currentNode = currentNode.Right
		} else {
			if currentNode.Left != nil && currentNode.Right != nil {
				currentNode.Value = currentNode.Right.getMinValue()
				currentNode.Right.remove(currentNode.Value, currentNode)
			} else if parent == nil {
				if currentNode.Left != nil {
					currentNode.Value = currentNode.Left.Value
					currentNode.Right = currentNode.Left.Right
					currentNode.Left = currentNode.Left.Left
				} else if currentNode.Right != nil {
					currentNode.Value = currentNode.Right.Value
					currentNode.Left = currentNode.Right.Left
					currentNode.Right = currentNode.Right.Right
				}
			} else if parent.Left == currentNode {
				if currentNode.Left != nil {
					parent.Left = currentNode.Left
				} else {
					parent.Left = currentNode.Right
				}
			} else if parent.Right == currentNode {
				if currentNode.Left != nil {
					parent.Right = currentNode.Left
				} else {
					parent.Right = currentNode.Right
				}
			}
			break
		}
	}

	return tree
}

func (tree *BST) getMinValue() int {
	currentNode := tree
	for currentNode.Left != nil {
		currentNode = currentNode.Left
	}
	return currentNode.Value
}

func main() {}
