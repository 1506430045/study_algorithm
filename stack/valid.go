package stack

func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) == 0 && len(popped) == 0 {
		return true
	}
	if len(pushed) == 0 {
		return false
	}
	if len(popped) == 0 {
		return false
	}
	s := InitStack()
	j := 1
	s.Push(pushed[0])
	for i := 0; i < len(popped); i++ {
		for s.Empty() || (!s.Empty() && s.Top().(int) != popped[i]) {
			if j >= len(pushed) {
				break
			}
			s.Push(pushed[j])
			j++
		}
		if !s.Empty() && s.Top().(int) == popped[i] {
			s.Pop()
		}
	}
	return s.Empty()
}
