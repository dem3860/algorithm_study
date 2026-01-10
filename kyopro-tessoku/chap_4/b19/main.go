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
	W := int64(io.ReadInt())

	weights := make([]int64, n)
	values := make([]int, n)

	for i := 0; i < n; i++ {
		weights[i] = io.ReadInt64()
		values[i] = io.ReadInt()
	}

	const INF int64 = 1e18
	maxValue := n * 1000

	// dp[i][j]: i番目までの品物で価値jを達成するための最小の重さ
	dp := make([][]int64, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int64, maxValue+1)
		for j := 0; j <= maxValue; j++ {
			dp[i][j] = INF
		}
	}

	// 初期条件
	dp[0][0] = 0

	for i := 1; i <= n; i++ {
		for j := 0; j <= maxValue; j++ {
			// 選ばない
			dp[i][j] = dp[i-1][j]

			// 選ぶ
			if j-values[i-1] >= 0 {
				dp[i][j] = min(
					dp[i][j],
					dp[i-1][j-values[i-1]]+weights[i-1],
				)
			}
		}
	}

	// 最大価値を探す
	answer := 0
	for v := maxValue; v >= 0; v-- {
		if dp[n][v] <= W {
			answer = v
			break
		}
	}

	io.Println(answer)
}
