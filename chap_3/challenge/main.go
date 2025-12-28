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

func lowerBound(a []int, x int) (int, bool) {
	i, found := slices.BinarySearch(a, x)
	return i, found
}

func main() {
	io := NewFastIO()
	defer io.Flush()

	n := io.ReadInt()

	a := make([]int, n)
	a_value := make([]int, n)

	for i := 0; i < n; i++ {
		val := io.ReadInt()
		a[i] = val
		a_value[i] = val
	}

	slices.Sort(a_value)

	uniq := make([]int, 0, n)
	for _, v := range a_value {
		if len(uniq) == 0 || uniq[len(uniq)-1] != v {
			uniq = append(uniq, v)
		}
	}

	b := make([]int, n)
	for i := 0; i < n; i++ {
		idx, _ := lowerBound(uniq, a[i])
		b[i] = idx
	}

	for i := 0; i < n; i++ {
		if i > 0 {
			io.Printf(" ")
		}
		io.Printf("%d", b[i])
	}
	io.Printf("\n")
}
