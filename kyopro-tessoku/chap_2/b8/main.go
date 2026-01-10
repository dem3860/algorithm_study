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

	n := io.ReadInt()

	x := make([]int, n+1)
	y := make([]int, n+1)

	for i := 1; i <= n; i++ {
		x[i] = io.ReadInt()
		y[i] = io.ReadInt()
	}

	q := io.ReadInt()

	a := make([]int, q+1)
	b := make([]int, q+1)
	c := make([]int, q+1)
	d := make([]int, q+1)

	for i := 1; i <= q; i++ {
		a[i] = io.ReadInt()
		b[i] = io.ReadInt()
		c[i] = io.ReadInt()
		d[i] = io.ReadInt()
	}

	// 各座標の点の個数を数えるための2次元スライス
	z := make([][]int, 1501)
	for i := 0; i <= 1500; i++ {
		z[i] = make([]int, 1501)
	}

	for i := 0; i <= n; i++ {
		z[x[i]][y[i]] += 1
	}

	//累積和用スライス
	t := make([][]int, 1501)
	for i := 0; i <= 1500; i++ {
		t[i] = make([]int, 1501)
		for j := 0; j <= 1500; j++ {
			t[i][j] = 0
		}
	}

	//横方向の累積和
	for i := 1; i <= 1500; i++ {
		for j := 1; j <= 1500; j++ {
			t[i][j] = t[i][j-1] + z[i][j]
		}
	}

	//　縦方向の累積和
	for j := 1; j <= 1500; j++ {
		for i := 1; i <= 1500; i++ {
			t[i][j] = t[i-1][j] + t[i][j]
		}
	}

	for i := 1; i <= q; i++ {
		ans := t[c[i]][d[i]] - t[a[i]-1][d[i]] - t[c[i]][b[i]-1] + t[a[i]-1][b[i]-1]
		io.Println(ans)
	}
}
