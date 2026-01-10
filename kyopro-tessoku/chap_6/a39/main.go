package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

type Movie struct {
	l int
	r int
}

func main() {
	io := NewFastIO()
	defer io.Flush()

	n := io.ReadInt()

	movies := make([]Movie, n)
	for i := 0; i < n; i++ {
		movies[i] = Movie{
			l: io.ReadInt(),
			r: io.ReadInt(),
		}
	}

	slices.SortFunc(movies, func(a, b Movie) int {
		if a.r < b.r {
			return -1
		}
		if a.r > b.r {
			return 1
		}
		return 0
	})

	currentTime := 0
	answer := 0

	for i := 0; i < n; i++ {
		if currentTime <= movies[i].l {
			currentTime = movies[i].r
			answer++
		}
	}

	io.Println(answer)

}
