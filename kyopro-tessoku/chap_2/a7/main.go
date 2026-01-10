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

// 自分の回答
// func main() {
// 	io := NewFastIO()
// 	defer io.Flush()

// 	d := io.ReadInt()
// 	n := io.ReadInt()

// 	l := make([]int, n)
// 	r := make([]int, n)

// 	day := make([]int, d)

// 	for i := 0; i < n; i++ {
// 		l[i] = io.ReadInt()
// 		r[i] = io.ReadInt()
// 		for j := l[i] - 1; j < r[i]; j++ {
// 			day[j]++
// 		}
// 	}

// 	for i := 0; i < d; i++ {
// 		io.Println(day[i])
// 	}

// }

func main() {
	io := NewFastIO()
	defer io.Flush()

	d := io.ReadInt()
	n := io.ReadInt()

	l := make([]int, n+1)
	r := make([]int, n+1)

	for i := 1; i <= n; i++ {
		l[i] = io.ReadInt()
		r[i] = io.ReadInt()
	}

	b := make([]int, d+2)

	for i := 1; i <= n; i++ {
		b[l[i]] += 1
		b[r[i]+1] -= 1
	}

	answer := make([]int, d+1)

	for i := 1; i <= d; i++ {
		answer[i] = answer[i-1] + b[i]
		io.Println(answer[i])
	}

}
