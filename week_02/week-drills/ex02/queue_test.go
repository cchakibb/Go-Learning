package queue

import "testing"

func TestEnqueue(t *testing.T) {
	q := Queue{}
	q.Enqueue(1)
	q.Enqueue(3)
	if q.items[1] != 3 {
		t.Errorf("q.items[1] should be 3, but got %d", q.items[1])
	}

}

func TestDequeue(t *testing.T) {
	q := Queue{}
	q.Enqueue(1)
	q.Enqueue(5)
	q.Enqueue(7)
	res, err := q.Dequeue()
	if err != nil {
		t.Error(err)
	}
	if res != 1 {
		t.Errorf("q.items[0] = 1, but got %d", res)
	}

}

func TestLen(t *testing.T) {
	q := Queue{}
	q.Enqueue(1)
	q.Enqueue(3)
	q.Enqueue(12)
	if len(q.items) != 3 {
		t.Errorf("Len should be 3, but got %d", len(q.items))
	}
}