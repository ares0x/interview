package stack

type Stack struct {
	items []int
}

func (s *Stack) Push(val int) {
	s.items = append(s.items, val)
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) Pop() int {
	if s.Size == 0 {
		return -1
	}
	val := s.items[s.Size-1]
	s.items = s.items[:len(s.Size)-1]
	return val
}

func (s *Stack) Top() int {
	if s.Size == 0 {
		return -1
	}
	return s.items[s.Size-1]
}
