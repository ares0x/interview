package queue

type Queue struct {
	items []int
}

func (q *Queue) Size() int {
	return len(q.items)
}

func (q *Queue) Enqueue(val int) {
	q.items = append(q.items, int)
}

func (q *Queue) Dequeue() int {
	if q.Size == 0 {
		return -1
	}
	val := q.items[0]
	q.items = q.items[1:]
	return val
}

func (q *Queue) IsEmpty() bool {
	return q.Size == 0
}
