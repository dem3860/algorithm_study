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

func (io *FastIO) PrintSliceInt(a []int) {
	for i := 0; i < len(a); i++ {
		if i > 0 {
			fmt.Fprint(io.writer, " ")
		}
		fmt.Fprint(io.writer, a[i])
	}
	fmt.Fprintln(io.writer)
}

func main() {
	io := NewFastIO()
	defer io.Flush()

	t := io.ReadInt()

	for i := 0; i < t; i++ {
		solve(io)
	}
}

func solve(io *FastIO) {
	n := io.ReadInt()
	w := io.ReadInt()

	mod := 2 * w

	// costs[k] は i % 2W == k となるマス i のコストの合計
	costs := make([]int64, mod)

	for i := 0; i < n; i++ {
		c := io.ReadInt64()
		idx := (i + 1) % mod
		costs[idx] += c
	}

	var currentSum int64 = 0
	for k := 0; k < w; k++ {
		currentSum += costs[k]
	}

	minSum := currentSum

	for k := 0; k < mod; k++ {
		currentSum -= costs[k]
		currentSum += costs[(k+w)%mod]

		if currentSum < minSum {
			minSum = currentSum
		}
	}

	io.Println(minSum)
}
