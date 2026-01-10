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
	k := io.ReadInt()

	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = io.ReadInt()
	}

	found := false

	// 2^n 通りの全探索
	// (1 << n) は 2のn乗
	for i := 0; i < (1 << n); i++ {
		sum := 0

		// j は 0 から n-1 まで（配列の添字に合わせる）
		for j := 0; j < n; j++ {

			// 【重要】ビット判定の定石テクニック
			// (1 << j) で「j番目のビットだけが1」の値を作る（マスク）
			// i と AND(&) を取り、0でなければ「j番目のビットが立っている」と判定
			if (i & (1 << j)) != 0 {
				sum += a[j]
			}
		}

		if sum == k {
			found = true
			break // 1つでも見つかれば終了して良い（高速化）
		}
	}

	if found {
		io.Println("Yes")
	} else {
		io.Println("No")
	}
}
