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

func check(a []int, x int, n int, k int) bool {
	sum := 0
	for i := 0; i < n; i++ {
		sum += x / a[i]
	}
	return sum >= k
}

func main() {
	io := NewFastIO()
	defer io.Flush()

	n := io.ReadInt()

	k := io.ReadInt()

	a := make([]int, n+1)

	for i := 1; i <= n; i++ {
		a[i] = io.ReadInt()
	}

	fmt.Println("a", a)

	//　累積和用スライス
	s := make([]int, n+1)
	s[0] = a[0]
	for i := 1; i <= n; i++ {
		s[i] = s[i-1] + a[i]
	}
	fmt.Println("s", s)

	// しゃくとり法
	r := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if i == 1 {
			r[i] = 0
		} else {
			r[i] = r[i-1]
		}
		for r[i]+1 <= n && s[r[i]+1]-s[i-1] <= k {
			r[i]++
		}
	}

	answer := 0
	for i := 0; i < n; i++ {
		if r[i] >= i {
			answer += (r[i] - i + 1)
		}
	}
	io.Println(answer)
}
