package lin

import "errors"

type Stack struct {
	slice []string
	cap   int
}

func NewStack(size int) *Stack {
	return &Stack{
		slice: make([]string, 0, size),
		cap:   size,
	}
}

func (s *Stack) Push(val string) error {
	if len(s.slice) >= s.cap {
		return errors.New("栈满了，压栈失败")
	}
	s.slice = append(s.slice, val)
	return nil
}

func (s *Stack) Pop() (string, error) {
	if len(s.slice) <= 0 {
		return "", errors.New("栈没有任何元素，出栈失败")
	}
	res := s.slice[len(s.slice)-1]
	s.slice = s.slice[:len(s.slice)-1]
	return res, nil
}

func (s *Stack) Peek() (string, error) {
	if len(s.slice) <= 0 {
		return "", errors.New("栈没有任何元素，出栈失败")
	}
	res := s.slice[len(s.slice)-1]
	return res, nil
}
