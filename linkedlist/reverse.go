package linkedlist

// Reverse reverses a singly linked list
func Reverse(node Node) Node {
	dummyHead := Node{Next: &node}

	iter := dummyHead.Next
	for iter.Next != nil {
		temp := iter.Next
		dummyHead.Next, iter.Next, temp.Next = temp, temp.Next, dummyHead.Next
	}

	return *dummyHead.Next
}
