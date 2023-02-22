package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	if l == nil {
		return nil
	}
	res := l
	idx := 0
	for idx < pos {
		if res.Next == nil {
			return nil
		}
		res = res.Next
		idx++
	}
	return res
}
