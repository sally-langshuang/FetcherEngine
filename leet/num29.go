package leet

import (
	"math"
	"unsafe"
)

func div(dividend int, divisor int) (d int, m int){
	//defer func(){fmt.Printf("%d / %d = %d mod %d\n", dividend, divisor, d, m)}()
	if dividend == 0 {
		return 0, 0
	}
	if dividend  < divisor {
		d = 0
		m = divisor
		return
	}
	if dividend - divisor < divisor {
		d = 1
		m = dividend - divisor
		return
	}
	left := dividend >> 1
	right := dividend - left
	leftDiv, leftMod := div(left, divisor)
	rightDiv, rightMod := div(right, divisor)
	x, m := div(leftMod+rightMod, divisor)
	d = leftDiv + rightDiv + x
	return
}
//假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−231,  231 − 1]。本题中，如果除法结果溢出，则返回 231 − 1。
func divide(dividend int, divisor int) int {
	negative := (dividend ^ divisor) >> (unsafe.Sizeof(dividend) * 8 - 1)
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor <0 {
		divisor = -divisor
	}
	x := 0
	i := 1
	div := divisor
	for dividend >= divisor {
		for dividend  >= (div << 1){
			div += div
			i += i
			continue
		}
		dividend -= div
		x += i
		for dividend < div {
			div -= div >> 1
			i -= i  >> 1
			if div < divisor || dividend < divisor {
				break
			}
		}
	}
	if negative == -1 {
		x = -x
	}
	if x > math.MaxInt32 {
		x = math.MaxInt32
	}
	if x < math.MinInt32 {
		x = math.MinInt32
	}
	return x
}
