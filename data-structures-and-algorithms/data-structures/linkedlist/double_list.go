package main

type DoublyNode struct {
	next, prev *Node
	value      int
	list       *DoublyList
}

type DoublyList struct {
	head *Node
	llen int
}

func (n *DoublyNode) Next() *DoublyNode {
	if p := n.next; p.list != nil && p != n.list.head {
		return p
	}
	return nil
}

func (n *DoublyNode) Prev() *DoublyNode {
	if p := n.prev; p.list != nil && p != n.list.head {
		return n
	}
	return nil
}

func (l *DoublyList) Init() *DoublyList {
	l.head.next = head
	l.head.prev = head
	l.llen = 0
	return l
}

func New() *DoublyList { return new(DoublyList).Init() }

func (l *DoublyList) Len() int { return l.llen }

func (l *DoublyList) Front() *DoublyNode {
	if l == nil {
		return nil
	}
	return l.head.next
}

func (l *DoublyList) Back() *DoublyNode {
	if l == nil {
		return nil
	}
	return l.head.prev
}

func (l *DoublyList) insert(n, at *DoublyNode) *DoublyNode {
	n.prev = at
	n.next = at.next
	n.prev.next = n
	n.next.prev = n
	n.list = l
	l.llen++
	return n
}

func (l *DoublyList) insertvalue(v int, at *DoublyNode) *DoublyNode {
	return l.insert(&DoublyNode{value: v}, at)
}

func (l *DoublyList) remove(n *DoublyNode) {
	n.prev.next = n.next
	n.next.prev = n.prev
	n.prev = nil
	n.next = nil
	n.list = nil
	l.llen--
}

func (l *DoublyList) move(n, at *DoublyNode) {
	if n == at {
		return
	}
	n.prev.next = n.next
	n.next.prev = n.prev

	n.prev = at
	n.next = at.next
	n.prev.next = n
	n.next.prev = n
}

func (l *DoublyList) lazyInit() {
	if l.head.next == nil {
		l.Init()
	}
}

func (l *DoublyList) Remove(n *DoublyNode) int {
	if n.list == l {
		return l.remove(n)
	}
	return n.value
}

func (l *DoublyList) PushFront(v int) *DoublyNode {
	l.lazyInit()
	return l.insertValue(v, l.head)
}

func (l *DoublyList) PushBack(v int) *DoublyNode {
	l.lazyInit()
	return l.insertValue(v, l.head.prev)
}

func (l *DoublyList) InsertBefore(v int, at *DoublyNode) *DoublyNode {
	if at.list != l {
		return nil
	}
	return l.insertValue(v, at.prev)
}

func (l *DoublyList) InsertAfter(v int, at *DoublyNode) *DoublyNode {
	if at.list != l {
		return nil
	}
	return l.insertValue(v, at)
}

func (l *DoublyList) MoveToFront(n *DoublyNode) {
	if n.list != l || l.head.next == n {
		return
	}
	l.move(n, l.head)
}

func (l *DoublyList) MoveToBack(n *DoublyNode) {
	if n.list != l || l.head.prev == n {
		return
	}
	l.move(n, l.head.prev)
}

func (l *DoublyList) MoveBefore(n, at *DoublyNode) {
	if n.list != l || at.list != l || n == at {
		return
	}
	l.move(n, at.prev)
}

func (l *DoublyList) MoveAfter(n, at *DoublyNode) {
	if n.list != l || at.list != l || n == at {
		return
	}
	l.move(n, at)
}
