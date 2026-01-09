package util

// ---- 素数判定(ルートnまでで割って、nが素数かを判定する) -------
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// ------ エラトステネスのふるい -----------
func sieve(n int) []bool {
	isPrime := make([]bool, n+1)

	// 2以上は一旦すべて素数候補
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}

	for i := 2; i*i <= n; i++ {
		if !isPrime[i] {
			continue
		}
		for j := i * i; j <= n; j += i {
			isPrime[j] = false
		}
	}
	return isPrime
}

// ------ ユークリッドの互除法 ----------
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// ------ aのb乗をmで割ったあまりを返す関数 --------
func Power(a, b, m int64) int64 {
	var result int64 = 1
	p := a % m

	for b > 0 {
		if b&1 == 1 {
			result = result * p % m
		}
		p = p * p % m
		b >>= 1
	}
	return result
}

// -----
func SumOfDigitSums(n int64) int64 {
	var res int64 = 0

	for factor := int64(1); factor <= n; factor *= 10 {
		higher := n / (factor * 10)
		cur := (n / factor) % 10
		lower := n % factor

		res += higher * 45 * factor
		res += (cur * (cur - 1) / 2) * factor
		res += cur * (lower + 1)
	}

	return res
}

// ----- 角桁の和 --------
func digitSum(x int) int {
	sum := 0
	for x > 0 {
		sum += x % 10
		x /= 10
	}
	return sum
}
