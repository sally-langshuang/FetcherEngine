package datastruct

import "fmt"

// finite length queue
type FiniteQueueArr struct {
	Arr []int
	head, tail int
}

func (q FiniteQueueArr) Empty() bool {
	return q.tail == q.head
}

func (q FiniteQueueArr) Full() bool {
	return q.tail % len(q.Arr) == q.head
}

func (q FiniteQueueArr) Push(val int) error {
	if q.Full() {
		return fmt.Errorf("cannot push to full queue")
	}
	q.Arr[q.tail] = val
	q.tail += 1
	return nil
}

func (q FiniteQueueArr) Pop() (val int, e error) {
	if q.Empty() {
		return -1, fmt.Errorf("cannot pop from empty queue")
	}
	val = q.Arr[q.head]
	q.head -= 1
	return
}