package main

import (
	"math/big"
)

func uniquePaths(m int, n int) int {
	return getFactorial(m+n-2) / (getFactorial(m-1) * getFactorial(n-1))
}

func uniquePaths_leetcode(m, n int) int {
	return int(new(big.Int).Binomial(int64(m+n-2), int64(n-1)).Int64())
}
