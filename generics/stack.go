package generics

type StackOfInts = Stack
type StackOfStrings = Stack

type Stack struct {
	values []interface{}
}

// refactoring with interface{}
// type StackOfInts struct {
// 	values []int
// }

func (s *Stack) Push(value interface{}) {
	s.values = append(s.values, value)
}

func (s *Stack) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *Stack) Pop() (interface{}, bool) {
	if s.IsEmpty() {
		return 0, false
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
