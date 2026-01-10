package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

type Graph [][]int

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

	for i := 0; i < n; i++ {
		sort.Ints(graph[i])

		io.Printf("%d: {", i+1)
		for j := 0; j < len(graph[i]); j++ {
			if j > 0 {
				io.Printf(", ")
			}
			io.Printf("%d", graph[i][j]+1)
		}
		io.Println("}")
	}

}
