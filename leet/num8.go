package leet

func myAtoi(s string) (res int) {
	const (
		stripLStep =  1 << iota
		sginStep
		readStep
	)

	step := stripLStep
	sign := 1
	for _, i := range s {
		switch step {
		case stripLStep:
			if i == ' ' {
				continue
			} else {
				step = step << 1
			}
			fallthrough
		case sginStep:
			if i == '-' {
				sign = -1
			}
			step = step << 1
			if i == '-' || i == '+' {
				continue
			}
			if i < '0' || i > '9' {
				return res
			}
			fallthrough
		case readStep:
			if i >= '0' && i <='9' {
				res = res * 10 + int(i - '0')
			} else {
				goto stop
			}
			if res * sign < int(int32(-1 << 31)) {
				res = int(int32(-1 << 31))
				return res
			} else if res * sign > int(int32(1 << 31 -1)){
				res = int(int32(1 << 31 -1))
				return res
			}
		}
	}
	stop:
	return res * sign
}