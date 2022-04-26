package linkedlist

import "fmt"

type Node struct {
	Data int
	Next *Node
}

// FromArray takes an array of integers and converts it to a linked list
func FromArray(arr []int) *Node {
	if len(arr) == 0 {
		return nil
	}

	dummyHead := Node{}
	iter := &dummyHead
	for _, val := range arr {
		iter.Next = &Node{Data: val}
		iter = iter.Next
	}
	fmt.Println()

	return dummyHead.Next
}

// ToArray takes a node (head of linked list) and returns array representation
// of a linked list start form that node
func ToArray(node *Node) []int {
	if node == nil {
		return []int{}
	}

	arr := []int{node.Data}

	iter := node
	for iter.Next != nil {
		arr = append(arr, iter.Next.Data)
		iter = iter.Next
	}

	return arr
}
