package main

import (
	"bufio"
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
	return &FastIO{sc, wr}
}

func (io *FastIO) ReadInt() int {
	io.scanner.Scan()
	v, _ := strconv.Atoi(io.scanner.Text())
	return v
}

func (io *FastIO) Println(a ...interface{}) {
	fmt.Fprintln(io.writer, a...)
}

func (io *FastIO) Flush() {
	io.writer.Flush()
}

type Graph [][]int

func Dfs(pos int, graph Graph, visited []bool, parent []int) {
	visited[pos] = true
	for _, next := range graph[pos] {
		if !visited[next] {
			parent[next] = pos
			Dfs(next, graph, visited, parent)
		}
	}
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
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = -1
	}

	Dfs(0, graph, visited, parent)

	path := []int{}
	cur := n - 1
	for cur != -1 {
		path = append(path, cur)
		cur = parent[cur]
	}

	for i := len(path) - 1; i >= 0; i-- {
		io.Println(path[i] + 1)
	}
}
