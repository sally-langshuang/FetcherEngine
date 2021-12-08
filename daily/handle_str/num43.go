package handle_str

import (
	"fmt"
	"strconv"
	"strings"
)

func add(i, val int, ans []int) {
	for x := i; val > 0; x-- {
		y := val + ans[x]
		if y > 9 {
			ans[x] = y % 10
			val = y / 10
			continue
		}
		ans[x] = y
		val = 0
	}
}
func multiply(num1 string, num2 string) string {
	ans := make([]int, len(num1) + len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j--{
			x := int(num1[i] - '0') * int(num2[j] - '0')
			add(i + j + 1, x, ans)
		}
	}
	res := &strings.Builder{}
	notZero := false
	for i:= range ans {
		if notZero == false && i != len(ans) - 1 && ans[i] == 0{
			continue
		}
		if ans[i] != 0 {
			notZero = true
		}
		fmt.Fprintf(res, strconv.Itoa(ans[i]))
	}
	return res.String()
}