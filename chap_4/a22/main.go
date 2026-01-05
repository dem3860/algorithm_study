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
	a := make([]int, n-1)
	b := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		a_val := io.ReadInt()
		a[i] = a_val - 1
	}

	for i := 0; i < n-1; i++ {
		b_val := io.ReadInt()
		b[i] = b_val - 1
	}

	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = -10000000
	}

	dp[0] = 0

	for i := 0; i < n-1; i++ {
		dp[a[i]] = max(dp[a[i]], dp[i]+100)
		dp[b[i]] = max(dp[b[i]], dp[i]+150)
	}

	io.Println(dp[n-1])
}
