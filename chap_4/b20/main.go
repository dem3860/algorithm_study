package main

import (
	"bufio"
	"fmt"
	"os"
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

func (io *FastIO) ReadString() string {
	if !io.scanner.Scan() {
		panic("入力が足りません")
	}
	return io.scanner.Text()
}

func (io *FastIO) Println(a ...interface{}) {
	fmt.Fprintln(io.writer, a...)
}

func (io *FastIO) Flush() {
	io.writer.Flush()
}

func main() {
	io := NewFastIO()
	defer io.Flush()

	s := io.ReadString()
	t := io.ReadString()

	rs := []rune(s)
	rt := []rune(t)

	n := len(rs)
	m := len(rt)

	// dp[i][j]: rs[:i] を rt[:j] に変える最小操作回数
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}

	// 初期条件
	for i := 0; i <= n; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= m; j++ {
		dp[0][j] = j
	}

	// 遷移
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if rs[i-1] == rt[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(
					dp[i-1][j]+1,   // 削除
					dp[i][j-1]+1,   // 挿入
					dp[i-1][j-1]+1, // 変更
				)
			}
		}
	}

	io.Println(dp[n][m])
}
