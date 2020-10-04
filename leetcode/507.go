// https://leetcode.com/problems/perfect-number/
package leetcode

import "math"

func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}
	ones, zeros := 0, 0
	isZeroContinuous := true
	for num > 0 {
		if num&1 == 0 {
			if !isZeroContinuous {
				return false
			}
			zeros++
		} else {
			isZeroContinuous = false
			ones++
		}
		num >>= 1
	}
	return ones-1 == zeros && isPrime(ones) && isPrime(int(math.Pow(2, float64(ones)))-1)
}

func isPrime(n int) bool {
	if n == 1 {
		return true
	}
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
