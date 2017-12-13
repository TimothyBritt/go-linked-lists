package nodes

// Exploring Linked Lists in Go
// Written by Timothy Britt, 12/12/2017
// In preparation for interviewing engineering candidates
// at Strings, Inc.

// What kind of data will we store inside the nodes of our linked list?
// If we want to store different types of data, we can use the interface{}, but this is not the
// best way for many reasons.

// We'll assume that in these examples we would like to store strings.
type Node struct {
	payload string
}

type SinglyLinkedNode struct {
	Node
	Next *SinglyLinkedNode
}

type DoublyLinkedNode struct {
	Node
	Prev *DoublyLinkedNode
	Next *DoublyLinkedNode
}