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
	v     int
	cost  int64
	depth int
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

type Edge struct {
	to   int
	cost int64
}

type Graph [][]Edge

var (
	okVertices []bool
	l          int
	s          int64
	t          int64
)

func Dfs_by_stack(pos int, graph Graph) {
	n := len(graph)

	todo := NewStack(n)

	// ノードposから始まる
	_ = todo.Push(State{v:pos,cost:0,depth: 0 })

	for !todo.IsEmpty() {
		cur, _ := todo.Pop()

		if cur.cost > t {
			continue
		}

		if cur.depth == l {
			if cur.cost >= s && cur.cost <= t {
				okVertices[cur.v] = true
			}
			continue
		}

		for _, next := range graph[cur.v] {
				nextState := State{
					v: next.to,
					cost: cur.cost + next.cost,
					depth: cur.depth + 1,
				}
				_ = todo.Push(nextState)
		}
	}
}

func main() {
	io := NewFastIO()
	defer io.Flush()

	n := io.ReadInt()
	m := io.ReadInt()
	l = io.ReadInt()
	s = io.ReadInt64()
	t = io.ReadInt64()

	graph := make(Graph, n)

	for i := 0; i < m; i++ {
		u := io.ReadInt() - 1
		v := io.ReadInt() - 1
		c := io.ReadInt64()

		graph[u] = append(graph[u], Edge{to: v, cost: c})
	}

	okVertices = make([]bool, n)

	Dfs_by_stack(0, graph)

	ans := make([]int,0)

	for i := 0;i < n;i++ {
		if okVertices[i] {
			ans = append(ans, i+1)
		}
	}

	if len(ans) == 0 {
		io.Println()
	}else {
		io.PrintSliceInt(ans)
	}

}
