package main

type Node struct {
	Value      int
	Prev, Next *Node
}

type DoublyLinkedList struct {
	Head, Tail *Node
}

func NewDoublyLinkedList() *DoublyLinkedList {
	// Write your code here.
	return &DoublyLinkedList{Head: nil}
}

func (ll *DoublyLinkedList) SetHead(node *Node) {
	// Write your code here.
	node.Next = ll.Head
	node.Prev = ll.Tail
	ll.Head = node
}

func (ll *DoublyLinkedList) SetTail(node *Node) {
	// Write your code here.

	ll.Tail = node
}

func (ll *DoublyLinkedList) InsertBefore(node, nodeToInsert *Node) {
	// Write your code here.

}

func (ll *DoublyLinkedList) InsertAfter(node, nodeToInsert *Node) {
	// Write your code here.
}

func (ll *DoublyLinkedList) InsertAtPosition(position int, nodeToInsert *Node) {
	// Write your code here.
}

func (ll *DoublyLinkedList) RemoveNodesWithValue(value int) {
	// Write your code here.
}

func (ll *DoublyLinkedList) Remove(node *Node) {
	// Write your code here.
}

func (ll *DoublyLinkedList) ContainsNodeWithValue(value int) bool {
	// Write your code here.
	return false
}
