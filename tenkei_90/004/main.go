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

	h := io.ReadInt()
	w := io.ReadInt()

	a := make([][]int, h)
	for i := 0; i < h; i++ {
		a[i] = make([]int, w)
		for j := 0; j < w; j++ {
			a[i][j] = io.ReadInt()
		}
	}

	row := make([]int, h)
	column := make([]int, w)

	for i := 0; i < h; i++ {
		sum := 0
		for j := 0; j < w; j++ {
			sum += a[i][j]
		}
		row[i] = sum
	}

	for j := 0; j < w; j++ {
		sum := 0
		for i := 0; i < h; i++ {
			sum += a[i][j]
		}
		column[j] = sum
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if j > 0 {
				fmt.Fprint(io.writer, " ")
			}
			fmt.Fprint(io.writer, row[i]+column[j]-a[i][j])
		}
		fmt.Fprintln(io.writer)
	}

}
