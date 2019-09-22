package main

import "fmt"

type Node struct {
	value int
	pre   *Node
	next  *Node
}

type DoublyLinkedList struct {
	first  *Node
	last   *Node
	length int
}

func (dl *DoublyLinkedList) Append(n *Node) {
	if dl.length == 0 {
		dl.first = n
		dl.last = n
		dl.length++
		return
	}

	preLastNode := dl.last

	dl.last = n
	dl.last.pre = preLastNode
	preLastNode.next = n
	dl.length++
}

func (dl *DoublyLinkedList) Prepend(n *Node) {
	if dl.length == 0 {
		dl.first = n
		dl.last = n
		dl.length++
		return
	}

	preFirstNode := dl.first

	dl.first = n
	dl.first.next = preFirstNode
	preFirstNode.pre = n
	dl.length++
}

func (dl *DoublyLinkedList) Lookup() []int {
	if dl.length == 0 {
		return nil
	}

	var nodes []int
	node := dl.first
	for node != nil {
		nodes = append(nodes, node.value)
		node = node.next
	}

	return nodes
}

func (dl *DoublyLinkedList) ReverseLookup() []int {
	if dl.length == 0 {
		return nil
	}

	var nodes []int
	node := dl.last
	for node != nil {
		nodes = append(nodes, node.value)
		node = node.pre
	}

	return nodes
}

func (dl *DoublyLinkedList) Delete(v int) {
	if dl.length == 0 {
		return
	}

	if dl.first.value == v {
		dl.first = dl.first.next
		dl.first.pre = nil
		dl.length--
		return
	}

	if dl.last.value == v {
		prePreLastNode := dl.last.pre
		dl.last = prePreLastNode
		prePreLastNode.next = nil
		dl.length--
		return
	}

	var prevNode *Node
	var nextNode *Node
	node := dl.first

	for node.value != v {
		if node == nil {
			return
		}

		prevNode = node
		node = node.next
		nextNode = node.next
	}

	prevNode.next = node.next
	nextNode.pre = prevNode
	dl.length--
}

func main() {
	dl := &DoublyLinkedList{}

	secondNode := &Node{
		value: 2,
	}
	dl.Append(secondNode)

	thirdNode := &Node{
		value: 3,
	}
	dl.Append(thirdNode)

	firstNode := &Node{
		value: 1,
	}
	dl.Prepend(firstNode)

	fmt.Println(dl.Lookup())
	fmt.Println(dl.ReverseLookup())
	fmt.Println(dl.length)

	dl.Delete(2)

	fmt.Println(dl.Lookup())
	fmt.Println(dl.ReverseLookup())
	fmt.Println(dl.length)
}
