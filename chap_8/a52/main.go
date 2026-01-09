package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type FastIO struct {
	scanner *bufio.Scanner
	writer  *bufio.Writer
}

func NewFastIO() *FastIO {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	const bufSize = 1024 * 1024
	sc.Buffer(make([]byte, bufSize), bufSize)

	wr := bufio.NewWriter(os.Stdout)

	return &FastIO{
		scanner: sc,
		writer:  wr,
	}
}

func (io *FastIO) ReadInt() int {
	if !io.scanner.Scan() {
		panic("入力が足りません")
	}
	i, _ := strconv.Atoi(io.scanner.Text())
	return i
}

func (io *FastIO) ReadInt64() int64 {
	if !io.scanner.Scan() {
		panic("入力が足りません")
	}
	v, _ := strconv.ParseInt(io.scanner.Text(), 10, 64)
	return v
}

// 読み込み用: string
func (io *FastIO) ReadString() string {
	if !io.scanner.Scan() {
		panic("入力が足りません")
	}
	return io.scanner.Text()
}

func (io *FastIO) Println(a ...interface{}) {
	fmt.Fprintln(io.writer, a...)
}

func (io *FastIO) Printf(format string, a ...interface{}) {
	fmt.Fprintf(io.writer, format, a...)
}

func (io *FastIO) Flush() {
	io.writer.Flush()
}

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

func (s *Stack) IsEmpty() bool {
	return s.top == 0
}

func (s *Stack) IsFull() bool {
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
	data []string
	head int
	tail int
	size int
	cap  int
}

func NewQueue(cap int) *Queue {
	return &Queue{
		data: make([]string, cap),
		head: 0,
		tail: 0,
		size: 0,
		cap:  cap,
	}
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) IsFull() bool {
	return q.size == q.cap
}

func (q *Queue) Push(x string) error {
	if q.IsFull() {
		return errors.New("queue is full")
	}
	q.data[q.tail] = x
	q.tail = (q.tail + 1) % q.cap
	q.size++
	return nil
}

func (q *Queue) Pop() (string, error) {
	if q.IsEmpty() {
		return "", errors.New("queue is empty")
	}
	v := q.data[q.head]
	q.head = (q.head + 1) % q.cap
	q.size--
	return v, nil
}

// 取り出すことなく先頭を参照するためのメソッド
func (q *Queue) Top() (string, error) {
	if q.IsEmpty() {
		return "", errors.New("queue is empty")
	}
	return q.data[q.head], nil
}

func (q *Queue) Size() int {
	return q.size
}

func main() {
	io := NewFastIO()
	defer io.Flush()

	q := io.ReadInt()

	queue := NewQueue(q)

	for i := 0; i < q; i++ {
		queryType := io.ReadInt()

		if queryType == 1 {
			x := io.ReadString()
			err := queue.Push(x)
			if err != nil {
				io.Println(err)
			}
		}

		if queryType == 2 {
			res, err := queue.Top()
			if err != nil {
				io.Println(err)
				continue
			}
			io.Println(res)
		}

		if queryType == 3 {
			_, err := queue.Pop()
			if err != nil {
				io.Println(err)
			}
		}
	}

}
