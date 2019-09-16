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
