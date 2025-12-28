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

	a := make([]int, n)
	list_true := make([]int, n)
	list_false := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = io.ReadInt()
		if i == 0 {
			if a[i] == 1 {
				list_true[i] = 1
				list_false[i] = 0
			} else {
				list_true[i] = 0
				list_false[i] = 1
			}
		} else {
			if a[i] == 1 {
				list_true[i] = list_true[i-1] + 1
				list_false[i] = list_false[i-1]
			} else {
				list_true[i] = list_true[i-1]
				list_false[i] = list_false[i-1] + 1
			}
		}
	}

	q := io.ReadInt()
	l := make([]int, q)
	r := make([]int, q)

	for i := 0; i < q; i++ {
		l[i] = io.ReadInt()
		r[i] = io.ReadInt()
	}

	for i := 0; i < q; i++ {
		from := -1
		if l[i]-2 >= 0 {
			from = l[i] - 2
		}
		to := r[i] - 1

		if from == -1 {
			if list_true[to] > list_false[to] {
				io.Println("win")
			} else if list_true[to] < list_false[to] {
				io.Println("lose")
			} else {
				io.Println("draw")
			}
		} else {
			true_count := list_true[to] - list_true[from]
			false_count := list_false[to] - list_false[from]
			if true_count > false_count {
				io.Println("win")
			} else if true_count < false_count {
				io.Println("lose")
			} else {
				io.Println("draw")
			}
		}
	}
}
