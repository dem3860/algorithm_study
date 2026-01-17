package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// FastIOテンプレート
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
	return &FastIO{scanner: sc, writer: wr}
}

func (io *FastIO) ReadInt() int {
	if !io.scanner.Scan() {
		panic("入力が足りません")
	}
	i, _ := strconv.Atoi(io.scanner.Text())
	return i
}

func (io *FastIO) ReadString() string {
	if !io.scanner.Scan() {
		panic("入力が足りません")
	}
	return io.scanner.Text()
}

func (io *FastIO) Println(a ...interface{}) {
	fmt.Fprintln(io.writer, a...)
}

func (io *FastIO) Flush() {
	io.writer.Flush()
}

func canMake(word string, allowed map[rune]bool) bool {
	for _, char := range word {
		if !allowed[char] {
			return false
		}
	}
	return true
}

func main() {
	io := NewFastIO()
	defer io.Flush()

	_ = io.ReadInt()
	_ = io.ReadInt()
	S := io.ReadString()
	takahashiMap := make(map[rune]bool)
	for _, r := range S {
		takahashiMap[r] = true
	}

	T := io.ReadString()
	aokiMap := make(map[rune]bool)
	for _, r := range T {
		aokiMap[r] = true
	}

	Q := io.ReadInt()
	for i := 0; i < Q; i++ {
		w := io.ReadString()

		isTakahashi := canMake(w, takahashiMap)
		isAoki := canMake(w, aokiMap)

		if isTakahashi && !isAoki {
			io.Println("Takahashi")
		} else if !isTakahashi && isAoki {
			io.Println("Aoki")
		} else {
			io.Println("Unknown")
		}
	}
}
