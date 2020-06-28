package main

type stack struct {
	brackets []string
}

func newStack() stack {
	return stack{
		brackets: make([]string, 0),
	}
}

func (s *stack) push(b string) {
	s.brackets = append([]string{b}, s.brackets...)
}

func (s *stack) pop() (string, bool) {
	if s.isEmpty() {
		return "", false
	}
	head := s.brackets[0]
	s.brackets = s.brackets[1:]
	return head, true
}

func (s stack) isEmpty() bool {
	return len(s.brackets) == 0
}
