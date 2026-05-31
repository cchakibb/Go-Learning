package stack

import "testing"

func TestPush (t *testing.T) {
	s := Stack{}
	s.Push(3)
	if s.items[0] != 3 {
		t.Errorf("Just pushed 3, got %d", s.items[0])
		return
	}
	s.Push(4)
	if len(s.items) != 2 {
		t.Errorf("Just pushed 2 items. Items lenght should be 2, got %d", len(s.items))
		return
	}
}

func TestPop (t *testing.T) {
	s := Stack{}


	
	s.Push(1)
	s.Push(3)
	s.Push(5)

	val, err := s.Pop()
	if err != nil {
		t.Error(err)
		return
	}
	if val != 5 {
		t.Errorf("Last element had value 5, got %d", val)
		return
	}
	if len(s.items) != 2 {
		t.Errorf("Had 3 elements at the beginning. We poped the last element, so we should have 2 elements left, but got %d", len(s.items))
		return
	}
	
}

func TestLen (t *testing.T) {
	s := Stack{}
	if len(s.items) != 0 {
		t.Errorf("Stack len is 0, but got %d", len(s.items))
	}
	s.Push(1)
	if len(s.items) != 1 {
		t.Errorf("Stack len is 1, but got %d", len(s.items))
	}
	s.Push(2)
	s.Push(3)
	// s.items length = 3
	if len(s.items) != 3 {
		t.Errorf("Stack len is 3, but got %d", len(s.items))
	}
	s.Pop()
	if len(s.items) != 2 {
		t.Errorf("Stack len is 2, but got %d", len(s.items))
	}
}