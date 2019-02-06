package p1

import "reflect"

type item struct {
	value Node
	next  *item
}

func (i *item) isEmpty() bool {
	return reflect.DeepEqual(i, item{})
}

type Stack struct {
	top    *item
	length int
}

func (s *Stack) isEmpty() bool {
	return reflect.DeepEqual(s, Stack{})
}

func (s *Stack) len() int {
	return s.length
}

func (s *Stack) push(node Node) {
	s.top = &item{
		value: node,
		next:  s.top,
	}
	s.length++
}

func (s *Stack) pop() Node {
	if (s.len() > 0) {
		node := s.top.value
		s.top = s.top.next
		s.length--
		return node
	}

	return Node{}
}

func (s *Stack) retrieve() []Node {
	result := []Node{}
	for item := s.top; item!=nil; item = item.next {
		result = append(result, item.value)
	}
	return result
}
