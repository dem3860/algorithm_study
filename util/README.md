## util

```
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
```

## usage

main 関数内で最初に以下のように書く。
io := NewFastIO()
defer io.Flush()
その後、標準入力から整数を読み込むときは
n := io.ReadInt()
文字列を読み込むときは
s := io.ReadString()
標準出力に出力するときは
io.Println("Hello, World!")
や
io.Printf("Number: %d\n", n)
のように書く。

## lower_bound と upper_bound

```
package main

import (
	"fmt"
	"slices"
)

func main() {
	a := []int{1, 3, 3, 5, 7}

	// lower_bound
	i, found := slices.BinarySearch(a, 4)
	fmt.Println(i, found)

	x := 5

	// upper_bound
	i, found = slices.BinarySearch(a, x)
	for i < len(a) && a[i] == x {
		i++
	}
	fmt.Println(i, found)

}
```
