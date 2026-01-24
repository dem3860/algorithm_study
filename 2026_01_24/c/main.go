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

func (io *FastIO) Print(a ...interface{}) {
	fmt.Fprint(io.writer, a...)
}

func (io *FastIO) Flush() {
	io.writer.Flush()
}

func main() {
	io := NewFastIO()
	defer io.Flush()

	N := io.ReadInt()
	M := io.ReadInt()
	degree := make([]int, N+1)

	for i := 0; i < M; i++ {
		u := io.ReadInt()
		v := io.ReadInt()
		degree[u]++
		degree[v]++
	}

	for i := 1; i <= N; i++ {
		k := int64(N - 1 - degree[i])

		if k < 3 {
			io.Print("0")
		} else {

			ans := k * (k - 1) * (k - 2) / 6
			io.Print(ans)
		}

		if i < N {
			io.Print(" ")
		}
	}
    io.Print("\n")
}