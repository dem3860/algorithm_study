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
	v    int // 現在の頂点
	d    int // ここまでの距離
	tree int //　ここまでの木の本数
}

type PriorityQueue []Node

func (pq PriorityQueue) Len() int {
	return len(pq)
}

// 距離が同じなら木の多い方がいい
func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].d != pq[j].d {
		return pq[i].d < pq[j].d
	}
	// 木は多い方がいい
	return pq[i].tree > pq[j].tree
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
	to   int
	w    int
	tree int
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
		d := io.ReadInt()

		// 結ぶ先、距離、木の有無を登録
		graph[a] = append(graph[a], Edge{b, c, d})
		graph[b] = append(graph[b], Edge{a, c, d})
	}

	dist := make([]int, n)
	tree := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = INF
		tree[i] = -INF
	}

	dist[0] = 0
	tree[0] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, Node{v: 0, d: 0, tree: 0})

	for pq.Len() > 0 {
		node := heap.Pop(pq).(Node)
		v := node.v
		d := node.d
		t := node.tree

		if d > dist[v] {
			continue
		}

		// 距離が同じで木の本数が少ないものが得られたら既存の方がいいのでスキップする
		if d == dist[v] && t < tree[v] {
			continue
		}

		for _, e := range graph[v] {
			// 改善できるのであれば
			if dist[e.to] > dist[v]+e.w || (dist[v]+e.w == dist[e.to] && tree[v]+e.tree > tree[e.to]) {
				dist[e.to] = dist[v] + e.w
				tree[e.to] = tree[v] + e.tree
				heap.Push(pq, Node{v: e.to, d: dist[e.to], tree: tree[e.to]})
			}
		}
	}

	io.Println(dist[n-1], tree[n-1])

}
