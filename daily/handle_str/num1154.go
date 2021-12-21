package handle_str

import "strconv"

func dayOfYear(date string) (ans int) {
	y, _ := strconv.Atoi(date[:4])
	m, _ := strconv.Atoi(date[5:7])
	d, _ := strconv.Atoi(date[8:10])
	for i := 1; i < m; i++ {
		if i == 2 { //29 /28
			if y%3200 == 0 || (y%100 == 0 && y%400 != 0) || y%4 != 0 {
				ans += 28
			} else {
				ans += 29
			}
		} else if (i+i/8)%2 == 1 {
			ans += 31
		} else {
			ans += 30
		}
	}
	ans += d
	return ans
}
