package piscine

func ListPushFront(l *List, data interface{}) {
	n := NodeL{Data: data}
	if l.Head == nil {
		l.Head = &n
		l.Tail = l.Head
	} else {
		newNode := &NodeL{Data: data}
		newNode.Next = l.Head
		l.Head = newNode
	}
}
