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

// 配列と探索値を受け取り、そのindexを返す
func search(a []int, x int, n int) int {
	left := 1
	right := n

	for left <= right {
		mid := (left + right) / 2
		if a[mid] == x {
			return mid
		} else if a[mid] < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	io := NewFastIO()
	defer io.Flush()

	n := io.ReadInt()
	x := io.ReadInt()

	a := make([]int, n+1)

	for i := 1; i <= n; i++ {
		a[i] = io.ReadInt()
	}

	idx := search(a, x, n)
	io.Println(idx)

}
