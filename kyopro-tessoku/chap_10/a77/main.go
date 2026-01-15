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

// スコアの最大値がx以上かどうかを判定する関数
func check(a []int, n int, l int, k int, x int) bool {
	count := 0
	last := 0
	for i := 0; i < n; i++ {
		if a[i]-last >= x && l-a[i] >= x {
			count += 1
			last = a[i]
		}
	}
	if count >= k {
		return true
	}
	return false
}

func main() {
	io := NewFastIO()
	defer io.Flush()

	n := io.ReadInt()
	l := io.ReadInt()

	k := io.ReadInt()

	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = io.ReadInt()
	}

	left := 0
	right := 1000000000

	for left < right {
		mid := (left + right) / 2

		ans := check(a, n, l, k, mid)
		if ans == false {
			right = mid - 1
		}
		if ans == true {
			left = mid
		}
	}

	io.Println(left)
}
