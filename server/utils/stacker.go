package utils

type Stack[T StackTypes] []T

type StackTypes interface {
	any
}

func (s *Stack[T]) Push(val T) {
	*s = append(*s, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}
	index := len(*s) - 1
	val := (*s)[index]
	*s = (*s)[:index]
	return val, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}
	index := len(*s) - 1
	val := (*s)[index]
	return val, true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}
