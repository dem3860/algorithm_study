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

func solveQuery(a []int, x int, y int) int64 {
	n := len(a)

	// aの中でx未満の個数
	l, _ := slices.BinarySearch(a, x)

	left := x
	right := x + y + n

	for left < right {
		mid := (left + right) / 2

		r, _ := slices.BinarySearch(a, mid+1)
		// これでxからmidまでにaの値が何個あるか
		count := r - l
		// いくつ欠落しているか
		missing := (mid - x + 1) - count
		if missing >= y {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return int64(left)
}

func main() {
	io := NewFastIO()
	defer io.Flush()

	n := io.ReadInt()
	q := io.ReadInt()
	a := make([]int, n)

	for i := 0; i < n; i++ {
		a[i] = io.ReadInt()
	}

	slices.SortFunc(a, func(a, b int) int {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	})

	for i := 0; i < q; i++ {
		x := io.ReadInt()
		y := io.ReadInt()

		res := solveQuery(a, x, y)
		io.Println(res)

	}

}
