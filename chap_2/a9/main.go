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

	h := io.ReadInt()
	w := io.ReadInt()
	n := io.ReadInt()

	// 1-indexed (1始まり) で処理するため +1 サイズで確保
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

	// マップ作成 (H+1 x W+1)
	// 二次元配列の初期化はループで行うのが基本
	kingdom_map := make([][]int, h+2) // 範囲外アクセス防止のため少し大きめに(h+2)とっておくと安心です
	for i := 0; i <= h+1; i++ {
		kingdom_map[i] = make([]int, w+2)
	}

	// いもす法 (加算処理)
	for i := 1; i <= n; i++ {
		kingdom_map[a[i]][b[i]]++
		kingdom_map[c[i]+1][d[i]+1]++
		kingdom_map[a[i]][d[i]+1]--
		kingdom_map[c[i]+1][b[i]]--
	}

	// 横方向の累積和
	for i := 1; i <= h; i++ {
		for j := 1; j <= w; j++ {
			kingdom_map[i][j] += kingdom_map[i][j-1]
		}
	}

	// 縦方向の累積和
	for j := 1; j <= w; j++ {
		for i := 1; i <= h; i++ {
			kingdom_map[i][j] += kingdom_map[i-1][j]
		}
	}

	for i := 1; i <= h; i++ {
		for j := 1; j <= w; j++ {
			if j > 1 {
				io.Printf(" ")
			}
			io.Printf("%d", kingdom_map[i][j])
		}
		io.Printf("\n")
	}
}
