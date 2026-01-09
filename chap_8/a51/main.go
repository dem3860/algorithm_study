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
	data []string
	top  int
	cap  int
}

func NewStack(cap int) *Stack {
	return &Stack{
		data: make([]string, cap),
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

func (s *Stack) Push(x string) error {
	if s.IsFull() {
		return errors.New("stack is full")
	}

	s.data[s.top] = x
	s.top++
	return nil
}

func (s *Stack) Peek() (string, error) {
	if s.IsEmpty() {
		return "", errors.New("stack is empty")
	}
	return s.data[s.top-1], nil
}

func (s *Stack) Pop() (string, error) {
	if s.IsEmpty() {
		return "", errors.New("stack is empty")
	}
	s.top--
	return s.data[s.top], nil
}

func main() {
	io := NewFastIO()
	defer io.Flush()

	q := io.ReadInt()

	stack := NewStack(q)

	for i := 0; i < q; i++ {
		queryType := io.ReadInt()

		if queryType == 1 {
			x := io.ReadString()

			if err := stack.Push(x); err != nil {
				io.Println(err)
			}

		}

		if queryType == 2 {
			res, err := stack.Peek()
			if err != nil {
				io.Println(err)
				continue
			}
			io.Println(res)
		}

		if queryType == 3 {
			_, err := stack.Pop()
			if err != nil {
				io.Println(err)
			}
		}
	}
}
