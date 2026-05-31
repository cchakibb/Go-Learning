package queue

import "errors"

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(n int) { // Push / Append
	q.items = append(q.items, n)
}

func (q *Queue) Dequeue() (int, error) { // remove 1st element and returns it
	if len(q.items) == 0 {
		return 0, errors.New("Queue is empty")
	}
	firstVal := q.items[0]
	q.items = q.items[1:]

	return firstVal, nil

}

func (q *Queue) Len() int {

	return len(q.items)
}