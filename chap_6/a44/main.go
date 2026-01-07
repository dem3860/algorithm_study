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

	n := io.ReadInt()
	q := io.ReadInt()

	a := make([]int, n)

	for i := 0; i < n; i++ {
		a[i] = i + 1
	}

	// queryType = 2で反転されているかを示す状態変数
	state := 1

	for i := 0; i < q; i++ {
		queryType := io.ReadInt()

		if queryType == 1 {
			x := io.ReadInt()
			y := io.ReadInt()

			if state == 1 {
				a[x-1] = y
			}
			if state == 2 {
				a[n-x] = y
			}
		}

		if queryType == 2 {
			if state == 1 {
				state = 2
			} else {
				state = 1
			}
		}

		if queryType == 3 {
			x := io.ReadInt()
			if state == 1 {
				io.Println(a[x-1])
			}
			if state == 2 {
				io.Println(a[n-x])
			}
		}

	}
}
