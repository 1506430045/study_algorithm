package stack

type Stack struct {
	data     []interface{}
	length   int
	capacity int
}

func InitStack() *Stack {
	return &Stack{
		data:     make([]interface{}, 8),
		length:   0,
		capacity: 8,
	}
}

func (s *Stack) Push(val interface{}) {
	if s.length+1 > s.capacity { //需要扩容
		s.capacity = s.capacity * 2
		t := s.data
		s.data = make([]interface{}, s.capacity)
		copy(s.data, t)
	}
	s.data[s.length] = val
	s.length++
}

func (s *Stack) Pop() interface{} {
	if s.Empty() {
		return nil
	}
	data := s.data[s.length-1]
	s.data[s.length-1] = nil
	s.length--
	return data
}

func (s *Stack) Top() interface{} {
	if s.Empty() {
		return nil
	}
	return s.data[s.length-1]
}

func (s *Stack) Empty() bool {
	return s.length == 0
}