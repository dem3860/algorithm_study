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

// func main() {
// 	io := NewFastIO()
// 	defer io.Flush()

// 	h := io.ReadInt()
// 	w := io.ReadInt()

// 	x := make([][]int, h+1)

// 	res_per_row := make([][]int, h+1)
// 	for i := 1; i <= h; i++ {
// 		x[i] = make([]int, w + 1)
// 		res_per_row[i] = make([]int, w+1)
// 		for j := 1; j <= w; j++ {
// 			x[i][j] = io.ReadInt()
// 			res_per_row[i][j] = res_per_row[i][j-1] + x[i][j]
// 		}
// 	}

// 	q := io.ReadInt()

// 	a := make([]int, q)
// 	b := make([]int, q)
// 	c := make([]int, q)
// 	d := make([]int, q)

// 	for i := 0; i < q; i++ {
// 		a[i] = io.ReadInt()
// 		b[i] = io.ReadInt()
// 		c[i] = io.ReadInt()
// 		d[i] = io.ReadInt()
// 	}

// 	for i := 0; i < q; i++ {
// 		ans := 0
// 		for row := a[i]; row <= c[i]; row++ {
// 			ans += res_per_row[row][d[i]] - res_per_row[row][b[i]-1]
// 		}
// 		io.Println(ans)
// 	}

// }

func main() {
	io := NewFastIO()
	defer io.Flush()

	h := io.ReadInt()
	w := io.ReadInt()

	x := make([][]int, h+1)

	z := make([][]int, h+1)

	for i := 1; i <= h; i++ {
		x[i] = make([]int, w+1)
		for j := 1; j <= w; j++ {
			x[i][j] = io.ReadInt()
		}
	}

	q := io.ReadInt()

	a := make([]int, q)
	b := make([]int, q)
	c := make([]int, q)
	d := make([]int, q)

	for i := 0; i < q; i++ {
		a[i] = io.ReadInt()
		b[i] = io.ReadInt()
		c[i] = io.ReadInt()
		d[i] = io.ReadInt()
	}

	// zの初期化
	for i := 1; i <= h; i++ {
		z[i] = make([]int, w+1)
		for j := 1; j <= w; j++ {
			z[i][j] = 0
		}
	}

	// 横方向の累積和
	for i := 1; i <= h; i++ {
		for j := 1; j <= w; j++ {
			z[i][j] = z[i][j-1] + x[i][j]
		}
	}

	// 縦方向の累積和
	for j := 1; j <= w; j++ {
		for i := 1; i <= h; i++ {
			z[i][j] = z[i-1][j] + z[i][j]
		}
	}

	for i := 1; i <= q; i++ {
		io.Println(z[c[i]][d[i]] - z[a[i]-1][d[i]] - z[c[i]][b[i]-1] + z[a[i]-1][b[i]-1])
	}

}
