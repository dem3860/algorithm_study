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
	return &FastIO{scanner: sc, writer: bufio.NewWriter(os.Stdout)}
}

func (io *FastIO) ReadInt() int {
	if !io.scanner.Scan() {
		panic("入力が足りません")
	}
	v, _ := strconv.Atoi(io.scanner.Text())
	return v
}

func (io *FastIO) Println(a ...interface{}) {
	fmt.Fprintln(io.writer, a...)
}

func (io *FastIO) Flush() {
	io.writer.Flush()
}

// ===== セグメント木（RSQ） =====

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
	st.dat[pos] += x
	for pos >= 2 {
		pos >>= 1
		st.dat[pos] = st.dat[pos*2] + st.dat[pos*2+1]
	}
}

// 区間和 [l, r)
func (st *SegmentTree) Query(l, r, a, b, u int) int {
	if r <= a || b <= l {
		return 0
	}
	if l <= a && b <= r {
		return st.dat[u]
	}
	m := (a + b) / 2
	return st.Query(l, r, a, m, u*2) +
		st.Query(l, r, m, b, u*2+1)
}

func main() {
	io := NewFastIO()
	defer io.Flush()

	n := io.ReadInt()
	tree := &SegmentTree{}
	tree.Init(n)

	var ans int64 = 0

	for i := 0; i < n; i++ {
		a := io.ReadInt() // 1〜N

		// 自分より大きい値の個数
		cnt := tree.Query(a+1, n+1, 1, tree.siz+1, 1)
		ans += int64(cnt)

		// a が出現したことを記録
		tree.Update(a, 1)
	}

	io.Println(ans)
}
