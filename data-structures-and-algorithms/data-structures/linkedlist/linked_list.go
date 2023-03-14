package linked_list

type ListNode struct {
	Next *ListNode
	Val  interface{}
}

func (n *ListNode) Next() *ListNode {
	return n.Next
}

type LinkedList struct {
	head *ListNode
	len  int
}

func (l *LinkedList) Init() {
	l.head = new(ListNode)
	l.len = 0
}

func New() *LinkedList {
	l := new(LinkedList)
	l.Init()
	return l
}

func (l *LinkedList) Clear() {
	l.Init()
}

func (l *LinkedList) Front() *ListNode {
	return l.head.Next()
}

func (l *LinkedList) Back() *ListNode {
	if l.len == 0 {
		return nil
	}
	tmp := l.head.Next()
	for {
		if tmp.Next() == nil {
			return tmp
		}
		tmp = tmp.Next()
	}
}
