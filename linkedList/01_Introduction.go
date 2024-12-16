package linkedlist

import "fmt"

type Node struct {
	val  int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) AddToEnd(val int) {
	newNode := &Node{val: val}
	if ll.head == nil {
		ll.head = newNode
		return
	}

	current := ll.head

	for current.next != nil {
		current = current.next
	}

	current.next = &Node{val: val}
}

func (ll *LinkedList) Print() {
	current := ll.head

	for current != nil {
		fmt.Print(current.val, " ")
		current = current.next
	}
}

func ConvertArray2LL(arr []int) *LinkedList {
	ll := &LinkedList{}

	for _, val := range arr {
		ll.AddToEnd(val)
	}

	return ll
}


//------------------------------------------------------------------------------------------------------------
// package linkedlist

// import "fmt"

// // Node represents a single node in the linked list
// type Node[T any] struct {
// 	val  T
// 	next *Node[T]
// }

// // LinkedList represents a generic linked list
// type LinkedList[T any] struct {
// 	head *Node[T]
// }

// // AddToEnd adds a new node with the given value to the end of the linked list
// func (ll *LinkedList[T]) AddToEnd(val T) {
// 	newNode := &Node[T]{val: val}
// 	if ll.head == nil {
// 		ll.head = newNode
// 		return
// 	}

// 	current := ll.head
// 	for current.next != nil {
// 		current = current.next
// 	}
// 	current.next = newNode
// }

// // Print displays all the nodes in the linked list
// func (ll *LinkedList[T]) Print() {
// 	current := ll.head
// 	for current != nil {
// 		fmt.Print(current.val, " ")
// 		current = current.next
// 	}
// 	fmt.Println()
// }

// // ConvertArray2LL converts a slice of any type into a linked list
// func ConvertArray2LL[T any](arr []T) *LinkedList[T] {
// 	ll := &LinkedList[T]{}
// 	for _, val := range arr {
// 		ll.AddToEnd(val)
// 	}
// 	return ll
// }
