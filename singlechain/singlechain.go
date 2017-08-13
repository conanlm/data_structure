package singlechain

type Node struct {
	Data int
	Next *Node
}

func GetFirst(h *Node) *Node {
	if h.Next == nil {
		return nil
	}
	return h.Next
}

func GetLast(h *Node) *Node {
	if h.Next == nil {
		return nil
	}
	i := h
	for i.Next != nil {
		i = i.Next
		if i.Next == nil {
			return i
		}
	}
	return nil
}

func GetLength(h *Node) int {
	var i int = 0
	n := h
	for n.Next != nil {
		i++
		n = n.Next
	}
	return i
}
func Insert(h, d *Node, p int) bool {
	if h.Next == nil {
		h.Next = d
		return true
	}
	i := 0
	n := h
	for n.Next != nil {
		i++
		if i == p {
			if n.Next.Next == nil {
				n.Next = d
				return true
			} else {
				d.Next = n.Next
				n.Next = d.Next
				return true
			}
		}
		n = n.Next
		if n.Next == nil {
			n.Next = d
			return true
		}
	}
	return false
}

func GetLoc(h *Node, p int) *Node {
	if p < 0 || p > GetLength(h) {
		return nil
	}
	var i int = 0
	n := h
	for n.Next != nil {
		i++
		n = n.Next
		if i == p {
			return n
		}
	}
	return nil
}
