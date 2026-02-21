package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// --- FastIO (そのまま) ---
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

// 初期化
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
		st.dat[pos] = st.dat[pos*2] + st.dat[pos*2+1]
	}
}

// 区間和 [l, r)
// 呼び出すときは Query(l, r, 1, st.siz+1, 1) のように呼ぶ
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

	N := io.ReadInt()
	Q := io.ReadInt()

	tree := &SegmentTree{}
	tree.Init(N)
	A := make([]int, N+1)

	for i := 1; i <= N; i++ {
		val := io.ReadInt()
		A[i] = val
		tree.Update(i, val)
	}

	for i := 0; i < Q; i++ {
		t := io.ReadInt()

		if t == 1 {
			x := io.ReadInt()

			A[x], A[x+1] = A[x+1], A[x]

			tree.Update(x, A[x])
			tree.Update(x+1, A[x+1])

		} else {
			l := io.ReadInt()
			r := io.ReadInt()

			ans := tree.Query(l, r+1, 1, tree.siz+1, 1)
			io.Println(ans)
		}
	}
}
