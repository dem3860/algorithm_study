package util

import "errors"

// ------- Stack (sampleとしてintで実装した) --------------
type Stack struct {
	data []int
	top  int
	cap  int
}

func NewStack(cap int) *Stack {
	return &Stack{
		data: make([]int, cap),
		top:  0,
		cap:  cap,
	}
}

func (s Stack) IsEmpty() bool {
	return s.top == 0
}

func (s Stack) IsFull() bool {
	return s.top == s.cap
}

func (s *Stack) Push(x int) error {
	if s.IsFull() {
		return errors.New("stack is full")
	}

	s.data[s.top] = x
	s.top++
	return nil
}

func (s *Stack) Top() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	return s.data[s.top-1], nil
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	s.top--
	return s.data[s.top], nil
}

// ------- キュー -----------

type Queue struct {
	data []int
	head int
	tail int
	size int
	cap  int
}

func NewQueue(cap int) *Queue {
	return &Queue{
		data: make([]int, cap),
		head: 0,
		tail: 0,
		size: 0,
		cap:  cap,
	}
}

func (q Queue) IsEmpty() bool {
	return q.size == 0
}

func (q Queue) IsFull() bool {
	return q.size == q.cap
}

func (q *Queue) Push(x int) error {
	if q.IsFull() {
		return errors.New("queue is full")
	}
	q.data[q.tail] = x
	q.tail = (q.tail + 1) % q.cap
	q.size++
	return nil
}

func (q *Queue) Pop() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	v := q.data[q.head]
	q.head = (q.head + 1) % q.cap
	q.size--
	return v, nil
}

// 取り出すことなく先頭を参照するためのメソッド
func (q Queue) Top() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	return q.data[q.head], nil
}

func (q Queue) Size() int {
	return q.size
}
