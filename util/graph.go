package util

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
