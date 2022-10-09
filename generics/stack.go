package generics

// no need for an interface
// type StackOfInts = Stack
// type StackOfStrings = Stack

// type Stack struct {
// 	values []interface{}
// }

type Stack[T any] struct {
	values []T
}

// refactoring with interface{}
// type StackOfInts struct {
// 	values []int
// }

// func (s *Stack) Push(value interface{}) {
// 	s.values = append(s.values, value)
// }
func (s *Stack[T]) Push(value T) {
	s.values = append(s.values, value)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	index := len(s.values) - 1
	el := s.values[index]
	s.values = s.values[:index]
	return el, true
}

// refactoring with interface{}
// type StackOfStrings struct {
// 	values []string
// .. rest pretty much the same as above
// }

// then going to have to add type assertion,
// when I run in to type safety compiler error
