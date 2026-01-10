package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// 1. ここでUtilを呼び出す（初期化）
	io := NewFastIO()
	defer io.Flush() // 最後に必ず出力する

	// --- ここから問題を解く ---

	// intの読み込み (C++: cin >> n;)
	a := io.ReadInt()
	b := io.ReadInt()
	c := io.ReadInt()

	// 計算
	ans := a + b

	// 出力 (C++: cout << ans << endl;)
	io.Println(ans)
	io.Println(b + c)
}

// ==================================================
// 以下、Utilコード (問題の下の方に貼り付けっぱなしでOK)
// ==================================================

type FastIO struct {
	scanner *bufio.Scanner
	writer  *bufio.Writer
}

func NewFastIO() *FastIO {
	// 入力のセットアップ
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	// バッファを少し増やしておく（長い文字列対策）
	const bufSize = 1024 * 1024
	sc.Buffer(make([]byte, bufSize), bufSize)

	// 出力のセットアップ
	wr := bufio.NewWriter(os.Stdout)

	return &FastIO{
		scanner: sc,
		writer:  wr,
	}
}

// 読み込み用: int
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

// 書き込み用: Println
func (io *FastIO) Println(a ...interface{}) {
	fmt.Fprintln(io.writer, a...)
}

// 書き込み用: Printf (フォーマット指定したい時)
func (io *FastIO) Printf(format string, a ...interface{}) {
	fmt.Fprintf(io.writer, format, a...)
}

// 必須: バッファのフラッシュ
func (io *FastIO) Flush() {
	io.writer.Flush()
}
