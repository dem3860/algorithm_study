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

	h := io.ReadInt()
	w := io.ReadInt()

	// グリッド
	c := make([][]string, h)
	for i := 0; i < h; i++ {
		row := io.ReadString()
		c[i] = make([]string, w)
		for j := 0; j < w; j++ {
			c[i][j] = string(row[j])
		}
	}

	// dp[i][j] : (i,j)に何通り辿り着けるか
	dp := make([][]int64, h)
	for i := 0; i < h; i++ {
		dp[i] = make([]int64, w)
	}

	dp[0][0] = 1

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if c[i][j] == "#" {
				dp[i][j] = 0
				continue
			}

			if i > 0 {
				dp[i][j] += dp[i-1][j]
			}

			if j > 0 {
				dp[i][j] += dp[i][j-1]
			}
		}
	}

	io.Println(dp[h-1][w-1])
}
