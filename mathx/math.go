package mathx

// GCD 返回 a, b 的最大公因数
func GCD(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

// LCM 返回 a, b 的最小公倍数
func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

// IsPrime 判断素数
func IsPrime(a int) bool {
	if a < 2 {
		return false
	}
	for i := 2; i*i < a; i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}

// QMI 快速幂求 a ^ k % p
func QMI(a, k, p int) int {
	res := 1
	for k > 0 {
		if k&1 == 1 {
			res = res * a % p // 二进制后，当前第 0 位为 1
		}
		a = a * a % p
		k >>= 1
	}
	return res
}
