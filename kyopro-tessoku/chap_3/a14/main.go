package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

func check(a []int, x int, n int, k int) bool {
	sum := 0
	for i := 0; i < n; i++ {
		sum += x / a[i]
	}
	return sum >= k
}

func lowerBound(a []int, x int) (int, bool) {
	i, found := slices.BinarySearch(a, x)
	return i, found
}

func main() {
	io := NewFastIO()
	defer io.Flush()

	n := io.ReadInt()
	k := io.ReadInt()

	a := make([]int, n)
	b := make([]int, n)
	c := make([]int, n)
	d := make([]int, n)

	for i := 0; i < n; i++ {
		a[i] = io.ReadInt()
	}
	for i := 0; i < n; i++ {
		b[i] = io.ReadInt()
	}
	for i := 0; i < n; i++ {
		c[i] = io.ReadInt()
	}
	for i := 0; i < n; i++ {
		d[i] = io.ReadInt()
	}

	p := make([]int, n*n)
	q := make([]int, n*n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			p[i*n+j] = a[i] + b[j]
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			q[i*n+j] = c[i] + d[j]
		}
	}

	sort.Slice(p, func(i, j int) bool {
		return p[i] < p[j]
	})

	sort.Slice(q, func(i, j int) bool {
		return q[i] < q[j]
	})

	for i := 0; i < n*n; i++ {
		x := k - p[i]
		_, found := lowerBound(q, x)
		if found {
			io.Println("Yes")
			return
		}
	}

	io.Println("No")

}
