package main

// https://ieftimov.com/post/golang-datastructures-stacks-queues/
import (
	"fmt"
)

type Action struct {
	name string
	next *Action
}

type ActionHistory struct {
	top  *Action
	size int
}

func (history *ActionHistory) Apply(newAction *Action) {
	if history.top != nil {
		oldTop := history.top
		newAction.next = oldTop
	}
	history.top = newAction
	history.size++
}

func (history *ActionHistory) Undo() *Action {
	topAction := history.top
	if topAction != nil {
		history.top = topAction.next
	} else if topAction.next == nil {
		history.top = nil
	}
	history.size--
	return topAction
}

func main() {
	h := &ActionHistory{
		size: 0,
	}

	h.Apply(&Action{name: "Bold"})
	h.Apply(&Action{name: "Cut"})
	h.Apply(&Action{name: "Underline"})

	fmt.Println(h)

	fmt.Println(h.Undo())
	fmt.Println(h)

	fmt.Println(h.Undo())
	fmt.Println(h)

	fmt.Println(h.Undo())
	fmt.Println(h)
}
