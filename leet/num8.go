package leet


func myAtoi(s string) (res int) {
	const (
		stripLStep =  1 << iota
		sginStep
		readStep
	)

	step := stripLStep
	sign := 1
	for i:= range s {
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
			} else {
				step = step << 1
			}
			fallthrough
		case readStep:
			if i > '0' || i < '9' {
				res = res*10 + i - '0'
			} else {
				return res * sign
			}
		}
	}
	return res * sign
}