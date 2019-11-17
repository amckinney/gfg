package main

import "fmt"

// StringList represents a thread-safe linked list of strings.
type StringList struct {
	value string
	next  *StringList
}

// Push adds this element to the front of the list.
func (s *StringList) Push(val string) *StringList {
	return &StringList{
		value: val,
		next:  s,
	}
}

// End returns the end node of this StringList.
func (s *StringList) End() *StringList {
	for ; !s.next.IsEmpty(); s = s.next {
	}
	return s
}

// IsEmpty returns true if the StringList is empty.
func (s *StringList) IsEmpty() bool {
	return s == nil || (s.value == "" && s.next == nil)
}

// Reverse reverses this linked list in-place.
func (s *StringList) Reverse() *StringList {
	current := s
	var prev, next *StringList
	for current != nil {
		next = current.next
		current.next = prev

		prev = current
		current = next
	}
	return prev
}

// AsSlice returns the *StringList as a []string.
func (s *StringList) AsSlice() []string {
	var vals []string
	for ; s != nil; s = s.next {
		vals = append(vals, s.value)
	}
	return vals
}

func main() {
	var (
		list     = new(StringList)
		reversed = new(StringList)
	)
	for _, s := range []string{"r", "a", "c", "e", "c", "a", "r"} {
		list = list.Push(s)
		reversed = reversed.Push(s)
	}

	fmt.Printf("Palindrome: %v\n", isPalindrome(list, reversed))
}

func isPalindrome(list, reversed *StringList) bool {
	for {
		if list == nil && reversed == nil {
			return true
		}
		if (list == nil && reversed != nil) || (list != nil && reversed == nil) {
			return false
		}
		head, tail := list.value, reversed.value
		if head != tail {
			break
		}
		list, reversed = list.next, reversed.next
	}
	return false
}
