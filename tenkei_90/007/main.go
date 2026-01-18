package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

func (io *FastIO) PrintSliceInt(a []int) {
	for i := 0; i < len(a); i++ {
		if i > 0 {
			fmt.Fprint(io.writer, " ")
		}
		fmt.Fprint(io.writer, a[i])
	}
	fmt.Fprintln(io.writer)
}

func lowerBound(a []int, x int) (int, bool) {
	i, found := slices.BinarySearch(a, x)
	return i, found
}

// 便利な関数
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	io := NewFastIO()
	defer io.Flush()

	n := io.ReadInt()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = io.ReadInt()
	}

	slices.Sort(a)

	q := io.ReadInt()
	for i := 0; i < q; i++ {
		b := io.ReadInt()

		idx, _ := lowerBound(a, b)
		ans := 2000000000

		if idx < n {
			diff := abs(a[idx] - b)
			ans = min(ans, diff)
		}

		// 候補2: 左側の値 (A[idx-1])
		if idx > 0 {
			diff := abs(a[idx-1] - b)
			ans = min(ans, diff)
		}

		io.Println(ans)
	}

}
