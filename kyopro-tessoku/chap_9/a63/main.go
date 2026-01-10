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

type Graph [][]int

// posから各頂点への最短距離を返す
func Bfs_by_stack(pos int, graph Graph) []int {
	n := len(graph)
	dist := make([]int, n)

	for i := 0; i < n; i++ {
		dist[i] = -1
	}

	que := NewQueue(n)

	dist[pos] = 0
	_ = que.Push(pos)

	for !que.IsEmpty() {
		v, _ := que.Pop()

		for _, next := range graph[v] {
			if dist[next] == -1 {
				dist[next] = dist[v] + 1
				_ = que.Push(next)
			}
		}
	}
	return dist
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

	visited := make([]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = false
	}

	dist := Bfs_by_stack(0, graph)

	for i := 0; i < n; i++ {
		io.Println(dist[i])
	}

}
