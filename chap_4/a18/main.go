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

func main() {
	io := NewFastIO()
	defer io.Flush()

	n := io.ReadInt()
	s := io.ReadInt()

	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = io.ReadInt()
	}

	dp := make([][]bool, n+1)
	// dpを初期化
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, s+1)
	}

	// 0枚で0は再現可能
	dp[0][0] = true
	// それ以外は不可能
	for i := 1; i <= s; i++ {
		dp[0][i] = false
	}

	for i := 1; i <= n; i++ {
		for j := 0; j <= s; j++ {
			// i枚目を使わないパターン
			dp[i][j] = dp[i-1][j]
			// i枚目を使うパターン
			if j-a[i-1] >= 0 && dp[i-1][j-a[i-1]] {
				dp[i][j] = true
			}
		}
	}

	if dp[n][s] {
		io.Println("Yes")
	} else {
		io.Println("No")
	}

}
