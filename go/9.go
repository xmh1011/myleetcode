package main

func isPalindrome(x int) bool {
	var length, i int
	var num = [10]int{0}
	length = 0
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	} else if x < 10 {
		return true
	} else {
		for i = 0; x >= 10; i++ {
			num[i] = x % 10
			x = x / 10
			length++
			if x < 10 {
				num[i+1] = x
			}
		}
		length++
		for i = 0; i <= length/2; i++ {
			if num[i] != num[length-i-1] {
				return false
			}
		}
		return true
	}
}
