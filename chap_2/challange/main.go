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

func main() {
	io := NewFastIO()
	defer io.Flush()

	n := io.ReadInt()

	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		a[i] = io.ReadInt()
	}

	d := io.ReadInt()

	l := make([]int, d+1)
	r := make([]int, d+1)
	for i := 1; i <= d; i++ {
		l[i] = io.ReadInt()
		r[i] = io.ReadInt()
	}

	// 前側から累積maxを保持するスライス
	p := make([]int, n+1)
	p[1] = a[1]
	for i := 2; i <= n; i++ {
		p[i] = max(p[i-1], a[i])
	}

	// 後ろ側から累積maxを保持するスライス
	b := make([]int, n+1)
	b[n] = a[n]
	for i := n - 1; i >= 1; i-- {
		b[i] = max(b[i+1], a[i])
	}

	for i := 1; i <= d; i++ {
		var left_max int
		var right_max int

		left_max = p[l[i]-1]
		right_max = b[r[i]+1]

		if left_max > right_max {
			io.Println(left_max)
		} else {
			io.Println(right_max)
		}
	}

}
