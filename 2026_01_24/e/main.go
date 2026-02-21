package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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

// --- Monster Structure ---
type Monster struct {
	ID    int
	X, Y  int
	Angle float64
}

// 2つのモンスターが完全に同じ方向（同一直線上かつ同じ向き）か判定
// 外積が0 かつ 内積が正
func isSameDirection(a, b Monster) bool {
	// Cross Product: x1*y2 - x2*y1
	cp := int64(a.X)*int64(b.Y) - int64(b.X)*int64(a.Y)

	// Dot Product: x1*x2 + y1*y2
	dp := int64(a.X)*int64(b.X) + int64(a.Y)*int64(b.Y)

	// 外積0(平行) かつ 内積正(同じ向き)
	return cp == 0 && dp > 0
}

func main() {
	io := NewFastIO()
	defer io.Flush()

	N := io.ReadInt()
	Q := io.ReadInt()

	monsters := make([]Monster, N)
	for i := 0; i < N; i++ {
		monsters[i].ID = i + 1 // 1-based ID
		monsters[i].X = io.ReadInt()
		monsters[i].Y = io.ReadInt()
		// Atan2で角度を計算 (-PI ~ +PI)
		monsters[i].Angle = math.Atan2(float64(monsters[i].Y), float64(monsters[i].X))
	}

	// 1. 時計回り（角度の降順）にソート
	sort.Slice(monsters, func(i, j int) bool {
		return monsters[i].Angle > monsters[j].Angle
	})

	// ソート後の位置を記録するためのマップ (ID -> ソート後Index)
	posMap := make([]int, N+1)
	for i := 0; i < N; i++ {
		posMap[monsters[i].ID] = i
	}

	// 2. 配列を2倍にする (円環対策)
	// 実際に配列をコピーして、ブロック判定をシンプルにする
	doubledN := 2 * N
	doubledMonsters := make([]Monster, doubledN)
	copy(doubledMonsters, monsters)
	copy(doubledMonsters[N:], monsters)

	// 3. ブロック情報の事前計算
	// BlockStart[i] : インデックスiのモンスターと同じ方向の塊の「開始インデックス」
	// BlockEnd[i]   : インデックスiのモンスターと同じ方向の塊の「終了インデックス」

	blockStart := make([]int, doubledN)
	blockEnd := make([]int, doubledN)

	for i := 0; i < doubledN; {
		j := i
		// jを増やしていき、向きが違うやつが出るまで進める
		// ※doubledMonsters[i]とdoubledMonsters[j]を比較
		for j < doubledN && isSameDirection(doubledMonsters[i], doubledMonsters[j]) {
			j++
		}
		// [i, j-1] が同じ向きのブロック
		for k := i; k < j; k++ {
			blockStart[k] = i
			blockEnd[k] = j - 1
		}
		i = j
	}

	// 4. クエリ処理
	for i := 0; i < Q; i++ {
		idA := io.ReadInt()
		idB := io.ReadInt()

		idxA := posMap[idA]
		idxB := posMap[idB]

		// 【修正ポイント】
		// まず、AとBが「同じブロック（同じ方向）」にいるか確認する
		// これをしないと、同じブロック内で idxA > idxB だった場合に
		// 「1周回る」と誤判定されてしまう。
		if blockStart[idxA] == blockStart[idxB] { // 同じブロックIDを持っているか
			// 同じ方向なら回転しない。そのブロックの全モンスター数が答え。
			// ブロックのサイズ = End - Start + 1
			ans := blockEnd[idxA] - blockStart[idxA] + 1
			io.Println(ans)
			continue
		}

		// ここに来たら「方向が違う」ので、必ず回転が発生する。
		// 基本の移動先
		targetIdxB := idxB

		// 境界をまたぐ場合 (例: idxA=N-1(南), idxB=0(北) みたいなケースはidxA < idxBじゃないのでここ通らない)
		// 境界またぎ: A(角度小) -> B(角度大) ※降順ソートなので indexは大 -> 小 になるはずだが...
		// 降順ソートなので:
		// index 0 (180度), index 1 (90度), index 2 (0度) ...
		// 0 -> 1 (180 -> 90: 時計回り OK)
		// 2 -> 0 (0 -> 180: 時計回り OK, しかし 2 > 0 なのでまたぎ発生)
		if idxA > idxB {
			targetIdxB += N
		}

		// 計算: Aのブロックの先頭 〜 Bのブロックの末尾 まで
		realStart := blockStart[idxA]
		realEnd := blockEnd[targetIdxB]

		ans := realEnd - realStart + 1
		io.Println(ans)
	}
}
