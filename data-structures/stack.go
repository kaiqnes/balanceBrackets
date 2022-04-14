package data_structures

type Stack struct {
	values []interface{}
}

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.values) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(val interface{}) {
	s.values = append(s.values, val) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
	} else {
		index := len(s.values) - 1    // Get the index of the top most element.
		element := (s.values)[index]  // Index into the slice and obtain the element.
		s.values = (s.values)[:index] // Remove it from the stack by slicing it off.
		return element, true
	}
}
