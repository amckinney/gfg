package simple

// Queue represents a nil-safe queue of string elements.
type Queue struct {
	elems []string
}

// NewQueue creates a new *Queue with the
// given values. Note that these values
// are enqueued in the order that they
// are provided.
func NewQueue(values ...string) *Queue {
	return &Queue{
		elems: values,
	}
}

// Len returns the length of this queue.
func (q *Queue) Len() int {
	if q == nil {
		return 0
	}
	return len(q.elems)
}

// Dequeue removes and returns the element
// at the front of the queue.
func (q *Queue) Dequeue() string {
	if q.Len() == 0 {
		return ""
	}
	front := q.elems[0]
	q.elems = q.elems[1:]
	return front
}

// Enqueue adds an element to the end of the queue.
// The same Queue is returned so that this call can
// be chained.
//
//  q := NewStack()
//  q = q.Enqueue("foo").Enqueue("bar").Enqueue("baz")
//
//  foo, bar, baz := q.Dequeue(), q.Dequeue(), q.Dequeue()
//  ...
func (q *Queue) Enqueue(val string) *Queue {
	if q == nil {
		return NewQueue(val)
	}
	q.elems = append(q.elems, val)
	return q
}
