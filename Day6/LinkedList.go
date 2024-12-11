package main

import (
	"fmt"
)

type Node struct {
	data string
	next *Node
}

type MyLinkedList struct {
	head *Node
}

func (list *MyLinkedList) insertAtLast(data string) {
	newNode := &Node{data: data, next: nil}

	if list.head == nil {
		list.head = newNode
		return
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}

	current.next = newNode
}

func (list *MyLinkedList) insertAtFront(data string) {
	if list.head == nil {
		newNode := &Node{data: data, next: nil}
		list.head = newNode
		return
	}

	newNode := &Node{data: data, next: list.head}
	list.head = newNode
}

func (list *MyLinkedList) insertAfterValue(afterValue, data string) {
	newNode := &Node{data: data, next: nil}

	current := list.head

	for current != nil {
		if current.data == afterValue {
			newNode.next = current.next
			current.next = newNode
			return
		}
		current = current.next
	}

	fmt.Println("error to insert after this value")
}

func (list *MyLinkedList) inserBeforeValue(beforeValue, data string) {
	if list.head == nil {
		return
	}

	if list.head.data == beforeValue {
		newNode := &Node{data: data, next: list.head}
		list.head = newNode
		return
	}

	current := list.head
	for current.next != nil {
		if current.next.data == beforeValue {
			newNode := &Node{data: data, next: current.next}
			current.next = newNode
			return
		}
		current = current.next
	}
}

func (list *MyLinkedList) getLast() *Node {
	current := list.head

	for current.next != nil {
		current = current.next
	}

	return current
}

func (list *MyLinkedList) Print() {
	current := list.head
	for current != nil {
		fmt.Printf("%s -> ", current.data)
		current = current.next
	}
	fmt.Println()
}

func (list *MyLinkedList) findNodeAt(index int) *Node {
	var count int = 0
	var current *Node = list.head

	for current != nil {
		count++
		current = current.next
	}

	if index <= 0 || index > count {
		return nil
	}

	current = list.head
	for count = 1; count < index; count++ {
		current = current.next
	}

	return current
}
