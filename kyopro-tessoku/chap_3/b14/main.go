package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
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

func (io *FastIO) ReadInt64() int64 {
	if !io.scanner.Scan() {
		panic("入力が足りません")
	}
	v, _ := strconv.ParseInt(io.scanner.Text(), 10, 64)
	return v
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

func enumerateSubsetSums(x []int) []int {
	m := len(x)
	res := make([]int, 0, 1<<m)

	for mask := 0; mask < (1 << m); mask++ {
		sum := 0
		for i := 0; i < m; i++ {
			if (mask>>i)&1 == 1 {
				sum += x[i]
			}
		}
		res = append(res, sum)
	}

	return res
}

func lowerBound(a []int, x int) (int, bool) {
	i, found := slices.BinarySearch(a, x)
	return i, found
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

	fmt.Println("a", a)

	l1 := make([]int, n/2)
	l2 := make([]int, n-n/2)

	for i := 0; i < n/2; i++ {
		l1[i] = a[i]
	}

	for i := n / 2; i < n; i++ {
		l2[i-n/2] = a[i]
	}

	fmt.Println("l1", l1)
	fmt.Println("l2", l2)

	fmt.Println("k", k)

	sums1 := enumerateSubsetSums(l1)
	sums2 := enumerateSubsetSums(l2)

	sort.Slice(sums1, func(i, j int) bool {
		return sums1[i] < sums1[j]
	})

	sort.Slice(sums2, func(i, j int) bool {
		return sums2[i] < sums2[j]
	})

	for i := 0; i < len(sums1); i++ {
		x := k - sums1[i]
		_, found := lowerBound(sums2, x)
		if found {
			io.Println("Yes")
			return
		}
	}

	io.Println("No")

}
