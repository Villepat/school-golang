package piscine

// Write a function ListPushBack that inserts a new element NodeL at the end of the list l while using the structure List.

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	n := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = n
	} else {
		elt := l.Head
		for elt.Next != nil {
			elt = elt.Next
		}
		elt.Next = n
	}
}
