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

func abs(x int) int {
	y := x >> 31
	return (x ^ y) - y
}

func main() {
	io := NewFastIO()
	defer io.Flush()

	n := io.ReadInt()
	h := make([]int, n)
	for i := 0; i < n; i++ {
		h[i] = io.ReadInt()
	}

	dp := make([]int, n)
	dp[0] = 0
	dp[1] = abs(h[1] - h[0])

	for i := 2; i < n; i++ {
		// 1つ前から来る場合
		cost1 := dp[i-1] + abs(h[i]-h[i-1])
		// 2つ前から来る場合
		cost2 := dp[i-2] + abs(h[i]-h[i-2])
		dp[i] = min(cost1, cost2)
	}

	io.Println(dp[n-1])
}
