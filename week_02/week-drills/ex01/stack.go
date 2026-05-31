package stack

import "errors"

type Stack struct {
	items []int
}

func (s *Stack) Push(n int) {
	s.items = append(s.items, n)
}

func (s *Stack) Pop() (int, error) {

	if len(s.items) == 0 {
		return 0, errors.New("Slice is empty, can not pop")
	}

	lastIndex := len(s.items) - 1
	last := s.items[lastIndex] // value
	s.items = s.items[:lastIndex] // remove the last index

	return last, nil
}

func (s *Stack) Len() int {
	return len(s.items)
}

