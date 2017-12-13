package singlyLinkedList

import "linked_lists/nodes"

// Storing the length on the lists uses extra memory, but can give us the flexibility to
// perform additional operations on the lists.
type SinglyLinkedList struct {
	head *nodes.SinglyLinkedNode
	length int
}

func IsEmpty(s *SinglyLinkedList) bool {
	if s.head == nil {
		return true
	}

	return false
}

// Initialize an empty singly-linked list
func (s *SinglyLinkedList) Init() *SinglyLinkedList {
	s.length = 0
	return s
}

// New returns a pointer to an initialized Singly-Linked List
func New() *SinglyLinkedList {
	return new(SinglyLinkedList).Init()
}

// Head returns the first node in the list s
func (s *SinglyLinkedList) Head() *nodes.SinglyLinkedNode{
	return s.head
}

// Tail returns the last node in the list s
func (s *SinglyLinkedList) Tail() *nodes.SinglyLinkedNode{
	currentNode := s.head
	for currentNode != nil && currentNode.Next != nil {
		currentNode = currentNode.Next
	}
	return currentNode
}

//Append adds a new node to the end of list s
func (s *SinglyLinkedList) Append(n *nodes.SinglyLinkedNode) {
	// If the head is empty, this is a new list, place the the node into the head of the list
	// It has no next, yet...
	if s.head == nil {
		s.head = n
	} else {
		// Otherwise, we need to iterate over the list and find the end, and then set the next pointer
		// of the last item to the incoming/appended node
		currentNode := s.head

		for currentNode.Next != nil {
			currentNode = currentNode.Next
		}

		currentNode.Next = n
	}

	s.length++
}

//Prepend adds a new node to the front of list s
func (s *SinglyLinkedList) Prepend(n *nodes.SinglyLinkedNode) {
	if IsEmpty(s) {
		s.head = n
	} else {
		n.Next = s.head
		s.head = n
	}

	s.length++
}

//InsertBefore inserts node insert before node in list s
func (s *SinglyLinkedList) InsertBefore(inserted *nodes.SinglyLinkedNode, before *nodes.SinglyLinkedNode) {
	// If the head is the pointer to the before node, put the item in the front of the list
	if s.head == before {
		inserted.Next = before
		s.head = inserted
		s.length++
	} else {
		//Otherwise, loop through the nodes until we find the node BEFORE the before node:
		currentNode := s.head
		for currentNode != nil && currentNode.Next != nil && currentNode.Next != before {
			// this is the node before the "before" node, step currentNode to the next node
			currentNode = currentNode.Next
		}

		if currentNode.Next == before {
			inserted.Next = before
			currentNode.Next = inserted
			s.length++
		}
	}
}







