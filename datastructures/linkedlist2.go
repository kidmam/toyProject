package main

// https://dkvist.com/posts/linked-lists-golang/
type Node struct {
	value int   // Our data
	next  *Node // Pointer (link) to the next Node in the list
}

type LinkedList struct {
	first  *Node // The first Node
	last   *Node // The last Node
	length int   // The number of nodes
}

func (l *LinkedList) Append(n *Node) {
	// We check if our Linked list is empty
	if l.length == 0 {
		l.first = n // If it is, we set the first Node to the new Node
		l.last = n  // We also set the last Node to the new Node
		l.length++  // Finally, we increment the length
		return
	}

	l.last.next = n // We link our last Node to out new last Node
	l.last = n      // Then, we set the last Node Node to the new Node
	l.length++      // And finally, we increment the length
}

func (l *LinkedList) Prepend(n *Node) {
	// As in Append, first we check if the linked list is empy
	if l.length == 0 {
		l.first = n // If it is, we set the first Node the new Node
		l.last = n  // We also set the last Node to the new Node
		l.length++  // Finally, we increment the length
		return
	}

	firstNode := l.first     // We save our old first Node in a temporal variable
	l.first = n              // Then, we set the first Node to the new Node
	l.first.next = firstNode // Finally we linke the new first Node to the old one
	l.length++               // And then, we increment the length
}

func (l *LinkedList) Lookup() []int {
	// If the list is empty we return nil
	if l.length == 0 {
		return nil
	}

	var nodes []int
	node := l.first

	// If the next Node is nil (doesn't exist) it means that
	// we're done.
	for node != nil {
		nodes = append(nodes, node.value)
		node = node.next
	}

	return nodes // Then, we can return the slice with all the values found
}

func (l *LinkedList) Delete(v int) {
	// If the list is empty we have nothing to delete.
	if l.length == 0 {
		return
	}

	// If instead, the Node to delete is the first one on our list
	// we can remove it directly.
	if l.first.value == v {
		l.first = l.first.next // We set the first Node to the second Node
		l.length--             // Then, we decrement the length of our list
		return
	}

	var prevNode *Node // A reference to the previous Node of the one we are looking for
	node := l.first
	for node.value != v {
		// If we have reached the end of the list it means that
		// our Node doesn't exist and that we have nothing to remove.
		if node == nil {
			return
		}

		prevNode = node
		node = node.next
	}

	prevNode.next = node.next // We remove the Node by removing the reference to it from the previous one
	l.length--                // Finally, we decrement the length
}
