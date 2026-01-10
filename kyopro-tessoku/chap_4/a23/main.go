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

func (io *FastIO) Printf(format string, a ...interface{}) {
	fmt.Fprintf(io.writer, format, a...)
}

func (io *FastIO) Flush() {
	io.writer.Flush()
}

const INF = 1000000000

func main() {
	io := NewFastIO()
	defer io.Flush()

	n := io.ReadInt()
	m := io.ReadInt()

	// a[i][k] = 1 なら i枚目クーポンで品物kが無料（0-index）
	a := make([][]int, m)
	for i := 0; i < m; i++ {
		a[i] = make([]int, n)
		for j := 0; j < n; j++ {
			a[i][j] = io.ReadInt()
		}
	}

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, 1<<n)
		for mask := 0; mask < (1 << n); mask++ {
			dp[i][mask] = INF
		}
	}

	// クーポン0枚で「何も無料じゃない」状態(0)は0枚
	dp[0][0] = 0

	// i枚目まで使うかどうか（i=1..m）
	for i := 1; i <= m; i++ {
		for mask := 0; mask < (1 << n); mask++ {
			// 使わない：そのまま
			dp[i][mask] = min(dp[i][mask], dp[i-1][mask])

			// 使う：mask と a[i-1] を合成して newMask を作る
			newMask := mask

			// mask の各ビットの状態 + クーポンで無料になる品物 を反映
			for k := 0; k < n; k++ {
				has := (mask>>k)&1 == 1 // すでに無料ならtrue
				can := a[i-1][k] == 1   // このクーポンで無料ならtrue
				if has || can {
					newMask |= 1 << k
				}
			}

			dp[i][newMask] = min(dp[i][newMask], dp[i-1][mask]+1)
		}
	}

	full := (1 << n) - 1
	if dp[m][full] == INF {
		io.Println(-1)
	} else {
		io.Println(dp[m][full])
	}
}
