package main

import (
	"bufio"
	"container/heap"
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

type Node struct {
	d int
	v int
}

type PriorityQueue []Node

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].d < pq[j].d
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(Node))
}

func (pq PriorityQueue) Top() Node {
	return pq[0]
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

type Edge struct {
	to int
	w  int
}

const INF = 2000000000

type Graph [][]Edge

func main() {
	io := NewFastIO()
	defer io.Flush()

	n := io.ReadInt()
	m := io.ReadInt()

	graph := make(Graph, n)

	for i := 0; i < m; i++ {
		a := io.ReadInt() - 1
		b := io.ReadInt() - 1
		c := io.ReadInt()

		graph[a] = append(graph[a], Edge{b, c})
		graph[b] = append(graph[b], Edge{a, c})
	}

	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = INF
	}

	dist[0] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, Node{v: 0, d: 0})

	for pq.Len() > 0 {
		node := heap.Pop(pq).(Node)
		v := node.v
		d := node.d

		if d > dist[v] {
			continue
		}

		for _, e := range graph[v] {
			// 改善できるのであれば
			if dist[e.to] > dist[v]+e.w {
				dist[e.to] = dist[v] + e.w
				heap.Push(pq, Node{v: e.to, d: dist[e.to]})
			}
		}
	}

	for i := 0; i < n; i++ {
		if dist[i] == INF {
			io.Println("-1")
		} else {
			io.Println(dist[i])
		}
	}
}
