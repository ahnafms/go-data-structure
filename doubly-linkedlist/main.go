package main

import "fmt"

type LinkedList struct {
	head *Node
	tail *Node

	_length int
}

type Node struct {
	data int
	next *Node
	prev *Node
}

func (list *LinkedList) ValidatePosition(position int) bool {
	if position > list._length || position < 0 {
		fmt.Println("position is out of range")
		return false
	}
	return true
}

func (list *LinkedList) AddNode(data int) {
	var n Node
	ptr := &n
	ptr.data = data

	if list.head == nil {
		list.head = ptr
		list.tail = ptr
	} else {
		current := list.head

		for current.next != nil {
			current = current.next
		}

		list.tail = ptr

		current.next = ptr
		ptr.prev = current
	}

	list._length = +1
}

func (list *LinkedList) InsertBegin(data int) {
	var n Node
	ptr := &n
	ptr.data = data

	list.head.prev = ptr
	ptr.next = list.head
	list.head = ptr

	list._length++
}

func (list *LinkedList) InsertLast(data int) {
	var n Node
	ptr := &n
	ptr.data = data

	list.tail.next = ptr
	ptr.prev = list.tail
	list.tail = ptr

	list._length++
}

func (list *LinkedList) InsertAtPosition(position int, data int) {
	if !list.ValidatePosition(position) {
		return
	}

	var n Node
	ptr := &n
	n.data = data

	current := list.head
	currentIdx := 0

	for currentIdx < position-1 {
		current = current.next
		currentIdx++
	}

	ptr.next = current.next
	ptr.prev = current

	current.next.prev = ptr
	current.next = ptr

	list._length++
}

func (list *LinkedList) RemoveBegin() {
	current := list.head.next
	current.prev = nil
	list.head = current

	list._length--
}

func (list *LinkedList) RemoveLast() {
	current := list.tail.prev
	current.next = nil
	list.tail = current

	list._length--
}

func (list *LinkedList) RemoveAtPosition(position int) {
	current := list.head

	currentIdx := 0
	if currentIdx != position-1 {
		current = current.next
	}

	current.next = current.next.next
	current.next.prev = current

	list._length--
}

func (list *LinkedList) Display() {
	if list == nil {
		fmt.Println("List is empty")
		return
	}

	fmt.Println("Display")

	current := list.head

	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println("\n")
}

func (list *LinkedList) DisplayReverse() {
	current := list.tail

	fmt.Println("Display reverse")

	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.prev
	}
}

func main() {
	var newNode LinkedList
	newNode.AddNode(10)
	newNode.AddNode(20)
	newNode.AddNode(40)
	newNode.AddNode(50)
	newNode.AddNode(60)
	// newNode.InsertAtPosition(1, 70)

	// newNode.RemoveBegin()
	// newNode.RemoveLast()
	newNode.RemoveAtPosition(2)

	newNode.Display()
	newNode.DisplayReverse()
}
