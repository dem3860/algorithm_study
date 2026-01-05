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

func reverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	io := NewFastIO()
	defer io.Flush()

	n := io.ReadInt()
	a := make([]int, n)
	b := make([]int, n)
	for i := 1; i < n; i++ {
		a_val := io.ReadInt()
		a[i] = a_val
	}

	for i := 2; i < n; i++ {
		b_val := io.ReadInt()
		b[i] = b_val
	}

	// iからi+1へ行くときa_i+1,iからi+2へ行くとき、b_i+2を用いる
	dp := make([]int, n)

	for i := 0; i < n; i++ {
		dp[i] = 2000000000
	}

	dp[0] = 0

	for i := 0; i < n-1; i++ {
		dp[i+1] = min(dp[i+1], dp[i]+a[i+1])
		if i+2 < n {
			dp[i+2] = min(dp[i+2], dp[i]+b[i+2])
		}
	}

	io.Println(dp[n-1])
}
