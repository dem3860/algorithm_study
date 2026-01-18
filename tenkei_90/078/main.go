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

func (io *FastIO) PrintSliceInt(a []int) {
	for i := 0; i < len(a); i++ {
		if i > 0 {
			fmt.Fprint(io.writer, " ")
		}
		fmt.Fprint(io.writer, a[i])
	}
	fmt.Fprintln(io.writer)
}

type State struct {
	str   string // 作成される文字列
	left  int    // (の数
	right int    // )の数
}

type Stack struct {
	data []State
	top  int
	cap  int
}

func NewStack(cap int) *Stack {
	return &Stack{
		data: make([]State, cap),
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

func (s *Stack) Push(x State) error {
	if s.IsFull() {
		return errors.New("stack is full")
	}

	s.data[s.top] = x
	s.top++
	return nil
}

func (s *Stack) Top() (State, error) {
	if s.IsEmpty() {
		return State{}, errors.New("stack is empty")
	}
	return s.data[s.top-1], nil
}

func (s *Stack) Pop() (State, error) {
	if s.IsEmpty() {
		return State{}, errors.New("stack is empty")
	}
	s.top--
	return s.data[s.top], nil
}

func main() {
	io := NewFastIO()
	defer io.Flush()

	n := io.ReadInt()

	if n%2 != 0 {
		return
	}

	// 左右の括弧を n / 2ずつ使用する必要がある
	limit := n / 2

	stack := NewStack(1000000)

	_ = stack.Push(State{str: "", left: 0, right: 0})

	for !stack.IsEmpty() {
		cur, _ := stack.Pop()
		if len(cur.str) == n {
			io.Println(cur.str)
			continue
		}

		if cur.right < cur.left {
			nextState := State{
				str:   cur.str + ")",
				left:  cur.left,
				right: cur.right + 1,
			}
			_ = stack.Push(nextState)
		}

		if cur.left < limit {
			nextState := State{
				str:   cur.str + "(",
				left:  cur.left + 1,
				right: cur.right,
			}
			_ = stack.Push(nextState)
		}
	}
}
