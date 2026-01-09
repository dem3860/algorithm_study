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

func (io *FastIO) Flush() {
	io.writer.Flush()
}

type SegmentTree struct {
	siz int
	dat []int
}

func (st *SegmentTree) Init(n int) {
	st.siz = 1
	for st.siz < n {
		st.siz <<= 1
	}
	st.dat = make([]int, st.siz*2)
}

func (st *SegmentTree) Update(pos, x int) {
	pos = pos + st.siz - 1
	st.dat[pos] = x
	for pos >= 2 {
		pos >>= 1
		st.dat[pos] = max(st.dat[pos*2], st.dat[pos*2+1])
	}
}

func (st *SegmentTree) Query(l, r, a, b, u int) int {
	if r <= a || b <= l {
		return -1_000_000_000
	}
	if l <= a && b <= r {
		return st.dat[u]
	}
	m := (a + b) / 2
	left := st.Query(l, r, a, m, u*2)
	right := st.Query(l, r, m, b, u*2+1)
	return max(left, right)
}

func main() {
	io := NewFastIO()
	defer io.Flush()

	n := io.ReadInt()
	q := io.ReadInt()

	tree := &SegmentTree{}

	tree.Init(n)

	for i := 0; i < q; i++ {
		t := io.ReadInt()

		if t == 1 {
			pos := io.ReadInt()
			x := io.ReadInt()
			tree.Update(pos, x)
		}

		if t == 2 {
			l := io.ReadInt()
			r := io.ReadInt()
			res := tree.Query(l, r, 1, tree.siz+1, 1)
			io.Println(res)
		}
	}
}
