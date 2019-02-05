package p1

type item struct {
	value Node
	next *item
}

type Stack struct  {
	top *item
	length int
}

func (s *Stack) len() int {
	return s.length;
}

func (s *Stack) push(node Node){
	s.top = &item{
		value:node,
		next:s.top,
	}
	s.length++;
}


func (s *Stack) pop() *Node {
	if(s.len() > 0) {
		node := s.top.value;
		s.top = s.top.next;
		s.length--;
		return &node;
	}

	return nil;
}