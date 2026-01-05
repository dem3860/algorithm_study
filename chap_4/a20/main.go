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

	s := io.ReadString()
	t := io.ReadString()

	// 問題は英小文字のみだが、念の為
	rs := []rune(s)
	rt := []rune(t)

	len_s := len(rs)
	len_t := len(rt)

	// sのi文字目までとtのj文字目までの最長共通部分の文字数
	dp := make([][]int, len_s+1)
	for i := 0; i <= len_s; i++ {
		dp[i] = make([]int, len_t+1)
		for j := 0; j <= len_t; j++ {
			dp[i][j] = 0
		}
	}

	// 明示的に書いておく
	dp[0][0] = 0

	for i := 1; i <= len_s; i++ {
		for j := 1; j <= len_t; j++ {
			if rs[i-1] == rt[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	io.Println(dp[len_s][len_t])
}
