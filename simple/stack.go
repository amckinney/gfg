package simple

// Stack represents a nil-safe stack of string elements.
type Stack struct {
	elems []string
}

// NewStack creates a new *Stack with the
// given values. Note that these values
// are pushed onto the stack in the order
// that they are provided.
func NewStack(values ...string) *Stack {
	var elems []string
	for i := len(values) - 1; i >= 0; i-- {
		elems = append(elems, values[i])
	}
	return &Stack{
		elems: elems,
	}
}

// Len returns the length of this stack.
func (s *Stack) Len() int {
	if s == nil {
		return 0
	}
	return len(s.elems)
}

// Pop removes and returns the element
// on top of the stack.
func (s *Stack) Pop() string {
	if s.Len() == 0 {
		return ""
	}
	top := s.elems[0]
	s.elems = s.elems[1:]
	return top
}

// Peek returns the value on top of the stack
// without popping it off.
func (s *Stack) Peek() string {
	if s.Len() == 0 {
		return ""
	}
	return s.elems[0]
}

// Push pushes the given value on top of the stack.
// The same Stack is returned so that this call can
// be chained.
//
//  s := NewStack()
//  s = s.Push("foo").Push("bar").Push("baz")
//
//  baz, bar, foo := s.Pop(), s.Pop(), s.Pop()
//  ...
func (s *Stack) Push(val string) *Stack {
	if s == nil {
		return NewStack(val)
	}
	s.elems = append([]string{val}, s.elems...)
	return s
}
