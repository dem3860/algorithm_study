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

	a := make([]int, n+1)
	b := make([]int, n+1)
	c := make([]int, n+1)
	d := make([]int, n+1)

	for i := 1; i <= n; i++ {
		a[i] = io.ReadInt()
		b[i] = io.ReadInt()
		c[i] = io.ReadInt()
		d[i] = io.ReadInt()
	}

	paper_map := make([][]int, 1501)
	for i := 0; i <= 1500; i++ {
		paper_map[i] = make([]int, 1501)
	}

	for i := 1; i <= n; i++ {
		paper_map[a[i]][b[i]]++
		paper_map[c[i]][d[i]]++
		paper_map[a[i]][d[i]]--
		paper_map[c[i]][b[i]]--
	}

	// 横方向の累積和
	for i := 1; i <= 1500; i++ {
		for j := 1; j <= 1500; j++ {
			paper_map[i][j] += paper_map[i][j-1]
		}
	}

	// 縦方向の累積和
	for j := 1; j <= 1500; j++ {
		for i := 1; i <= 1500; i++ {
			paper_map[i][j] += paper_map[i-1][j]
		}
	}

	ans := 0
	for i := 1; i <= 1500; i++ {
		for j := 1; j <= 1500; j++ {
			if paper_map[i][j] > 0 {
				ans++
			}
		}
	}

	fmt.Println(ans)
}
