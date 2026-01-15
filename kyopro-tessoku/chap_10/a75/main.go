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

type Work struct {
	t int
	d int
}

func main() {
	io := NewFastIO()
	defer io.Flush()

	n := io.ReadInt()
	work := make([]Work, n+1)
	for i := 1; i <= n; i++ {
		work[i].t = io.ReadInt()
		work[i].d = io.ReadInt()
	}

	slices.SortFunc(work, func(a, b Work) int {
		if a.d < b.d {
			return -1
		}
		if a.d > b.d {
			return 1
		}
		return 0
	})

	// dp[i][j] : 手順iの時点で時刻jの時、最大の解けている問題数
	dp := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 1441)
		for j := 0; j <= 1440; j++ {
			dp[i][j] = -1
		}
	}

	dp[0][0] = 0

	for i := 1; i <= n; i++ {
		for j := 0; j <= 1440; j++ {
			// 1. i問目を解かないとき
			dp[i][j] = max(dp[i][j], dp[i-1][j])

			// 2. i問目を回答するとき
			if j-work[i].t >= 0 && work[i].d >= j {
				dp[i][j] = max(dp[i][j], dp[i-1][j-work[i].t]+1)
			}

		}
	}

	ans := 0

	for i := 0; i <= 1440; i++ {
		ans = max(ans, dp[n][i])
	}

	io.Println(ans)

}
