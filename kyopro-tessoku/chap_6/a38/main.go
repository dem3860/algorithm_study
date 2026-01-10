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

	d := io.ReadInt()
	n := io.ReadInt()

	l := make([]int, n)
	r := make([]int, n)
	h := make([]int, n)

	for i := 0; i < n; i++ {
		l_in := io.ReadInt()
		r_in := io.ReadInt()
		h_in := io.ReadInt()

		l[i] = l_in - 1
		r[i] = r_in - 1
		h[i] = h_in
	}

	// lim[i] : i日目の最大労働可能時間
	lim := make([]int, d)

	for i := 0; i < d; i++ {
		lim[i] = 24
	}

	for i := 0; i < n; i++ {
		for j := l[i]; j <= r[i]; j++ {
			lim[j] = min(lim[j], h[i])
		}
	}

	answer := 0
	for i := 0; i < d; i++ {
		answer += lim[i]
	}

	io.Println(answer)

}
