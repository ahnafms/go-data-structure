package main

import "fmt"

type LinkedList struct {
	head    *Node
	_length int
}

type Node struct {
	data int
	next *Node
}

func (list *LinkedList) ValidatePosition(position int) bool {
	if position > list._length || position < 0 {
		fmt.Println("position is out of range")
		return false
	}
	return true
}

func (list *LinkedList) Insert(data int) {
	var n Node
	ptr := &n
	ptr.data = data
	if list.head == nil {
		list.head = ptr
		list._length += 1
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = ptr
		list._length += 1
	}
}

func (list *LinkedList) Remove(position int) {
	current := list.head

	if !list.ValidatePosition(position) {
		return
	}

	currentIdx := 0
	for currentIdx != position-1 {
		current = current.next
		currentIdx++
	}

	deletedNode := current.next

	current.next = current.next.next
	fmt.Printf("Remove node position %d : %d memory address->%d\n", position, deletedNode.data, &deletedNode)
}

func (list *LinkedList) SearchNodeAt(position int) {
	current := list.head
	currentIdx := 0

	if !list.ValidatePosition(position) {
		return
	}

	for currentIdx != position {
		current = current.next
		currentIdx++
	}
	fmt.Printf("Search node position %d : %d\n", position, current.data)
}

func (list *LinkedList) Display() {
	current := list.head

	fmt.Printf("Display All Nodes: ")

	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
}

func main() {
	var newNode LinkedList
	newNode.Insert(10)
	newNode.Insert(20)
	newNode.Insert(30)
	newNode.Insert(40)

	newNode.SearchNodeAt(2)
	newNode.Remove(2)
	newNode.Display()
}
