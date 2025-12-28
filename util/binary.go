package util

import "slices"

// ----- 二分探索関数 -----
// 配列と探索値を受け取り、そのindexを返す
func search(a []int, x int, n int) int {
	left := 1
	right := n

	for left <= right {
		mid := (left + right) / 2
		if a[mid] == x {
			return mid
		} else if a[mid] < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// ------- lower bound関数 -------
// c++のlower_boundに相当する関数
// aという配列の中からx以上の最小の値のindexとそれが配列内に存在するかどうかのboolを返す
func lowerBound(a []int, x int) (int, bool) {
	i, found := slices.BinarySearch(a, x)
	return i, found
}

// ------- 全列挙 -------
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

