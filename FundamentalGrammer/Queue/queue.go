package queue

// An FIFO Queue
type Queue []interface{}

// Push an element into the queue
func (q *Queue) Push(value interface{}) {
	*q = append(*q, value)
}

// Pop element from queue
func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

// Returns element is empty or not
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
