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

func GetScore(a int, b int, k int, A []int, B []int, n int) int {
	cnt := 0

	for i := 0; i < n; i++ {
		if a <= A[i] && A[i] <= a+k && b <= B[i] && B[i] <= b+k {
			cnt++
		}
	}
	return cnt
}

func main() {
	io := NewFastIO()
	defer io.Flush()

	n := io.ReadInt()
	k := io.ReadInt()

	A := make([]int, n)
	B := make([]int, n)

	for i := 0; i < n; i++ {
		A[i] = io.ReadInt()
		B[i] = io.ReadInt()
	}

	max_ans := 0
	for i := 1; i <= 100; i++ {
		for j := 1; j <= 100; j++ {
			max_ans = max(max_ans, GetScore(i, j, k, A, B, n))
		}
	}

	io.Println(max_ans)
}
