package singlyLinkedList

import "go-linked-lists/nodes"

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
			currentNode = currentNode.Next
		}

		// the currentNode is now in front of the "before" node.
		if currentNode.Next == before {
			// point the next pointer of the inserted node to the "before" node
			inserted.Next = before
			// then point the next pointer of the currentNode to the inserted node
			currentNode.Next = inserted
			// the node has been successfully inserted into the list
			s.length++
		}
	}
}

func (s *SinglyLinkedList) InsertAfter(inserted *nodes.SinglyLinkedNode, after *nodes.SinglyLinkedNode) {
	// If the head is the pointer to the "after" node, then point the next pointer of the
	// inserted node to the head's next pointer and then point the next pointer of the
	// head to the inserted node
	if s.head == after {
		inserted.Next = s.head.Next
		s.head.Next = inserted
	} else {
		currentNode := s.head.Next
		// Otherwise find the "after" node
		for currentNode != nil && currentNode.Next != nil && currentNode != after {
			currentNode = currentNode.Next
		}

		// the currentNode is now the "after" node.
		if currentNode == after {
			// point the next pointer of the inserted node to the currentNode's next pointer
			inserted.Next = currentNode.Next
			// point the next pointer of the currentNode to the inserted node
			currentNode.Next = inserted
			// the node has successfully been inserted into the list
			s.length++
		}
	}
}

func (s *SinglyLinkedList) Reverse() *SinglyLinkedList {
	reversedList := new(SinglyLinkedList).Init()
	currentNode := s.head

	for currentNode != nil {
		reversedList.Prepend(currentNode)
		currentNode = currentNode.Next
		reversedList.length++
	}

	return reversedList
}

func (s *SinglyLinkedList) Delete(deleted *nodes.SinglyLinkedNode) {
	if s.head == deleted {
		s.head = s.head.Next
		deleted.Next = nil
		s.length--
	} else {
		currentNode := s.head.Next
		for currentNode != nil && currentNode.Next != deleted {
			currentNode = currentNode.Next
		}

		// the currentNode is before the node marked for deletion
		if currentNode.Next == deleted {
			currentNode.Next = deleted.Next
			deleted.Next = nil
			s.length--
		}
	}
}







