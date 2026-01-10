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

func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return -1, errors.New("stack is empty")
	}
	return s.data[s.top-1], nil
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return -1, errors.New("stack is empty")
	}
	s.top--
	return s.data[s.top], nil
}

func main() {
	io := NewFastIO()
	defer io.Flush()

	n := io.ReadInt()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = io.ReadInt()
	}

	stack := NewStack(n)
	ans := make([]int, n)

	for i := 0; i < n; i++ {
		for !stack.IsEmpty() {
			topIdx, _ := stack.Peek()
			if a[topIdx] > a[i] {
				break
			}
			stack.Pop()
		}

		if stack.IsEmpty() {
			ans[i] = -1
		} else {
			topIdx, _ := stack.Peek()
			ans[i] = topIdx + 1
		}
		stack.Push(i)
	}

	for i := 0; i < n; i++ {
		if i > 0 {
			io.Printf(" ")
		}
		io.Printf("%d", ans[i])
	}
	io.Println()
}
