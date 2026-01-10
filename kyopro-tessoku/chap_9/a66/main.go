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

func (io *FastIO) PrintSliceInt(a []int) {
	for i := 0; i < len(a); i++ {
		if i > 0 {
			fmt.Fprint(io.writer, " ")
		}
		fmt.Fprint(io.writer, a[i])
	}
	fmt.Fprintln(io.writer)
}

func (io *FastIO) Flush() {
	io.writer.Flush()
}

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

func main() {
	io := NewFastIO()
	defer io.Flush()

	n := io.ReadInt()
	q := io.ReadInt()

	uf := NewUnionFind(n)

	for i := 0; i < q; i++ {
		t := io.ReadInt()

		if t == 1 {
			u := io.ReadInt() - 1
			v := io.ReadInt() - 1
			uf.Unite(u, v)
		}

		if t == 2 {
			u := io.ReadInt() - 1
			v := io.ReadInt() - 1
			res := uf.IsSame(u, v)
			if res == true {
				io.Println("Yes")
			} else {
				io.Println("No")
			}
		}
	}

}
