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

func (io *FastIO) Println(a ...interface{}) {
	fmt.Fprintln(io.writer, a...)
}

func (io *FastIO) Flush() {
	io.writer.Flush()
}

func main() {
	io := NewFastIO()
	defer io.Flush()

	n := io.ReadInt()
	p := make([]int, n)
	a := make([]int, n)

	for i := 0; i < n; i++ {
		p[i] = io.ReadInt() - 1 // 0-index
		a[i] = io.ReadInt()
	}

	// dp[l][r]: l〜r が残っている状態での最大得点
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	// 区間の長さでループ
	for length := n - 1; length >= 0; length-- {
		for l := 0; l+length < n; l++ {
			r := l + length

			best := 0

			// 左側（l-1）を消す
			if l-1 >= 0 {
				score := 0
				if l <= p[l-1] && p[l-1] <= r {
					score = a[l-1]
				}
				best = max(best, dp[l-1][r]+score)
			}

			// 右側（r+1）を消す
			if r+1 < n {
				score := 0
				if l <= p[r+1] && p[r+1] <= r {
					score = a[r+1]
				}
				best = max(best, dp[l][r+1]+score)
			}

			dp[l][r] = best
		}
	}

	// 最後に 1 個だけ残る状態の最大値が答え
	answer := 0
	for i := 0; i < n; i++ {
		answer = max(answer, dp[i][i])
	}

	io.Println(answer)
}
