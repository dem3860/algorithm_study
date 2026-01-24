package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// --- FastIO Template ---
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

func (io *FastIO) Println(a ...interface{}) {
	fmt.Fprintln(io.writer, a...)
}

func (io *FastIO) Flush() {
	io.writer.Flush()
}

func main() {
	io := NewFastIO()
	defer io.Flush()

	Q := io.ReadInt()

	volume := 0
	isPlaying := false

	for i := 0; i < Q; i++ {
		a := io.ReadInt()

		if a == 1 {
			volume++
		} else if a == 2 {
			if volume > 0 {
				volume--
			}
		} else if a == 3 {
			isPlaying = !isPlaying
		}

		if volume >= 3 && isPlaying {
			io.Println("Yes")
		} else {
			io.Println("No")
		}
	}
}