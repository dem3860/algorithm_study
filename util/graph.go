package util

import "container/heap"

// -------- 深さ優先探索(再帰実装) ---------
type Graph [][]int

func Dfs(pos int, graph Graph, visited []bool) {
	visited[pos] = true
	for i := 0; i < len(graph[pos]); i++ {
		next := graph[pos][i]
		if visited[next] == false {
			Dfs(next, graph, visited)
		}
	}
}

// --------- 深さ優先探索(stackを用いる) ----------
func Dfs_by_stack(pos int, graph Graph) {
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
}

// ------------ 幅優先探索(queueを用いる) -------------
func Bfs_by_stack(pos int, graph Graph) {
	n := len(graph)
	visited := make([]bool, n)

	todo := NewQueue(n)

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
}

// ---------- ダイクストラ ---------
type Node struct {
	d int
	v int
}

type PriorityQueue []Node

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].d < pq[j].d }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(Node))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

const INF = 2000000000

type Edge struct {
	to int
	w  int
}

// Graph 型を定義しておくと関数で受け取る時に便利です
type Graph_Dijkstra [][]Edge

func Dijkstra(graph Graph_Dijkstra, start int) []int {
	n := len(graph)
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = INF
	}

	dist[start] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, Node{v: start, d: 0})

	for pq.Len() > 0 {
		node := heap.Pop(pq).(Node)
		v := node.v
		d := node.d

		// 既により短い経路が見つかっている場合はスキップ
		if d > dist[v] {
			continue
		}

		for _, e := range graph[v] {
			if dist[e.to] > dist[v]+e.w {
				dist[e.to] = dist[v] + e.w
				heap.Push(pq, Node{v: e.to, d: dist[e.to]})
			}
		}
	}

	return dist
}

// ---------- Union-Find ----------
type UnionFind struct {
	par  []int
	size []int
}

func NewUnionFind(n int) *UnionFind {
	par := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		par[i] = -1
		size[i] = 1
	}
	return &UnionFind{par, size}
}

// 経路圧縮
func (uf *UnionFind) Root(x int) int {
	if uf.par[x] == -1 {
		return x
	}
	uf.par[x] = uf.Root(uf.par[x])
	return uf.par[x]
}

func (uf *UnionFind) Unite(x, y int) bool {
	rx := uf.Root(x)
	ry := uf.Root(y)

	if rx == ry {
		return false
	}

	if uf.size[rx] < uf.size[ry] {
		rx, ry = ry, rx
	}

	uf.par[ry] = rx
	uf.size[rx] += uf.size[ry]
	return true
}

func (uf *UnionFind) IsSame(x, y int) bool {
	return uf.Root(x) == uf.Root(y)
}

func (uf *UnionFind) Size(x int) int {
	return uf.size[uf.Root(x)]
}
