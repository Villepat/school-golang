package piscine

func ListLast(l *List) interface{} {
	if l.Head == nil { // if there is no value in head, the link list is empty an we return nil
		return nil
	} else {
		current := l.Head         // starting point for while loop
		for current.Next != nil { // inspect elements until the second last one, which will not have nil in it
			current = current.Next
		}
		return current.Data // above we declared current to be .next, so this is the last nodes data
	}
}
