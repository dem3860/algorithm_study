package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"slices"
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

func (io *FastIO) PrintSliceInt(a []int) {
	for i := 0; i < len(a); i++ {
		if i > 0 {
			fmt.Fprint(io.writer, " ")
		}
		fmt.Fprint(io.writer, a[i])
	}
	fmt.Fprintln(io.writer)
}

type Item struct {
	loss    int // 現在の損失の合計
	lastIdx int // 最後に加算したdiffのindex
	count   int // 現在何枚入れたか
}
type PriorityQueue []Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].loss < pq[j].loss
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(Item))
}

func (pq PriorityQueue) Top() Item {
	return pq[0]
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

func main() {
	io := NewFastIO()
	defer io.Flush()

	n := io.ReadInt()
	k := io.ReadInt()
	x := io.ReadInt()
	a := make([]int, n)

	for i := 0; i < n; i++ {
		a[i] = io.ReadInt()
	}

	slices.SortFunc(a, func(a, b int) int {
		if a > b {
			return -1
		}
		if a < b {
			return 1
		}
		return 0
	})

	// 最も大きいa[0]からの差分。
	// 入れ替える時の指標として、m = a[0] - a[i]が最も小さくなるiを用いたいという発想
	diff := make([]int, 0, n-1)
	for i := 1; i < n; i++ {
		diff = append(diff, a[0]-a[i])
	}
	maxVal := int64(k * (a[0]))

	// 1つ目の答え（最大値）を出力
	io.Println(maxVal)

	// もし種類が1つしかない、または1つしか求めなくていい場合は終了
	if n == 1 || x == 1 {
		return
	}

	pq := &PriorityQueue{}
	heap.Init(pq)

	if k >= 1 && len(diff) > 0 {
		heap.Push(pq, Item{
			loss:    diff[0],
			lastIdx: 0,
			count:   1,
		})
	}

	for i := 0; i < x-1; i++ {
		if pq.Len() == 0 {
			break
		}

		item := heap.Pop(pq).(Item)

		io.Println(maxVal - int64(item.loss))

		// パターンA : Add(同じ変更をもう一度増やす)
		if item.count+1 <= k {
			heap.Push(pq, Item{
				loss:    item.loss + diff[item.lastIdx],
				lastIdx: item.lastIdx,
				count:   item.count + 1,
			})
		}

		// パターンB : Shift(最後の変更の種類を変える)
		if item.lastIdx+1 < len(diff) {
			newLoss := item.loss - diff[item.lastIdx] + diff[item.lastIdx+1]
			heap.Push(pq, Item{
				loss:    newLoss,
				lastIdx: item.lastIdx + 1,
				count:   item.count,
			})
		}
	}

}
