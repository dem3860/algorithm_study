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

func main() {
	io := NewFastIO()
	defer io.Flush()

	n := io.ReadInt()
	q := io.ReadInt()

	a := make([]int, n)
	list := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = io.ReadInt()
		if i == 0 {
			list[i] = a[i]
		} else {
			list[i] = list[i-1] + a[i]
		}
	}

	l := make([]int, q)
	r := make([]int, q)
	for i := 0; i < q; i++ {
		l[i] = io.ReadInt()
		r[i] = io.ReadInt()
	}

	fmt.Println(list)
	for i := 0; i < q; i++ {
		var from int = -1
		if l[i]-2 >= 0 {
			from = l[i] - 2
		}
		var to int = r[i] - 1
		io.Println("from:", from, " to:", to)
		if from == -1 {
			io.Println(list[to])
		} else {
			io.Println(list[to] - list[from])
		}
	}
}
