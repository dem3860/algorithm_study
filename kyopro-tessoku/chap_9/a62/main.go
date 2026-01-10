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

type Graph [][]int

func Dfs_by_stack(pos int, graph Graph) []bool {
	n := len(graph)
	visited := make([]bool, n)

	todo := NewStack(n)

	visited[pos] = true
	_ = todo.Push(pos)

	for !todo.IsEmpty() {
		v, _ := todo.Pop()

		for _, next := range graph[v] {
			if !visited[next] {
				visited[next] = true
				_ = todo.Push(next)
			}
		}
	}
	return visited
}

func main() {
	io := NewFastIO()
	defer io.Flush()

	n := io.ReadInt()
	m := io.ReadInt()

	graph := make(Graph, n)

	for i := 0; i < m; i++ {
		a := io.ReadInt() - 1
		b := io.ReadInt() - 1

		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	visited := Dfs_by_stack(0, graph)

	for i := 0; i < n; i++ {
		if visited[i] == false {
			io.Println("The graph is not connected.")
			return
		}
	}

	io.Println("The graph is connected.")

}
