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
	return &FastIO{sc, wr}
}

func (io *FastIO) ReadInt() int {
	io.scanner.Scan()
	v, _ := strconv.Atoi(io.scanner.Text())
	return v
}
func (io *FastIO) ReadInt64() int64 {
	io.scanner.Scan()
	v, _ := strconv.ParseInt(io.scanner.Text(), 10, 64)
	return v
}
func (io *FastIO) Println(a ...interface{}) {
	fmt.Fprintln(io.writer, a...)
}
func (io *FastIO) Flush() {
	io.writer.Flush()
}

// lower_bound for int64
func lowerBoundInt64(a []int64, x int64) int {
	l, r := 0, len(a)
	for l < r {
		m := (l + r) / 2
		if a[m] < x {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

const Mod = 1000000007

func main() {
	io := NewFastIO()
	defer io.Flush()

	n := io.ReadInt()
	w := io.ReadInt64()
	l := io.ReadInt64()
	r := io.ReadInt64()

	x := make([]int64, n+2)
	for i := 1; i <= n; i++ {
		x[i] = io.ReadInt64()
	}
	x[0] = 0
	x[n+1] = w

	dp := make([]int64, n+2)
	sum := make([]int64, n+2)

	dp[0] = 1
	sum[0] = 1

	for i := 1; i <= n+1; i++ {
		left := x[i] - r
		right := x[i] - l

		posL := lowerBoundInt64(x, left)
		posR := lowerBoundInt64(x, right+1)

		if posL < posR {
			if posL == 0 {
				dp[i] = sum[posR-1]
			} else {
				dp[i] = (sum[posR-1] - sum[posL-1] + Mod) % Mod
			}
		}

		sum[i] = (sum[i-1] + dp[i]) % Mod
	}

	io.Println(dp[n+1] % Mod)
}
