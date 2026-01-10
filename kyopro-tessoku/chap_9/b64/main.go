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
	return &FastIO{
		scanner: sc,
		writer:  bufio.NewWriter(os.Stdout),
	}
}

func (io *FastIO) ReadInt() int {
	io.scanner.Scan()
	v, _ := strconv.Atoi(io.scanner.Text())
	return v
}
func (io *FastIO) Printf(f string, a ...interface{}) {
	fmt.Fprintf(io.writer, f, a...)
}
func (io *FastIO) Println(a ...interface{}) {
	fmt.Fprintln(io.writer, a...)
}
func (io *FastIO) Flush() {
	io.writer.Flush()
}

/* ---------- ダイクストラ用 ---------- */

type Node struct {
	v int
	d int
}

type PriorityQueue []Node

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].d < pq[j].d }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x any)        { *pq = append(*pq, x.(Node)) }
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

type Graph [][]Edge

const INF = 1_000_000_000

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

	/* ---------- ダイクストラ ---------- */

	dist := make([]int, n)
	prev := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = INF
		prev[i] = -1
	}
	dist[0] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, Node{v: 0, d: 0})

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(Node)
		v, d := cur.v, cur.d
		if d > dist[v] {
			continue
		}
		for _, e := range graph[v] {
			if dist[e.to] > dist[v]+e.w {
				dist[e.to] = dist[v] + e.w
				prev[e.to] = v
				heap.Push(pq, Node{v: e.to, d: dist[e.to]})
			}
		}
	}

	/* ---------- 経路復元 ---------- */

	path := []int{}
	cur := n - 1
	for cur != -1 {
		path = append(path, cur)
		cur = prev[cur]
	}

	// 逆順にする
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	/* ---------- 出力 ---------- */

	for i := 0; i < len(path); i++ {
		if i > 0 {
			io.Printf(" ")
		}
		io.Printf("%d", path[i]+1)
	}
	io.Println()
}
