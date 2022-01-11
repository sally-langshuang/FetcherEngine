package handle_str

var week = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
var monthDays = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30}

func dayOfTheWeek(day int, month int, year int) string {
	day += (year-1971)*365 + (year-1969)/4
	for _, ds := range monthDays[:month-1] {
		day += ds
	}
	if month >= 3 && (year%400 == 0 || year%4 == 0 && year%100 != 0) {
		day++
	}

	day += 3
	day %= 7
	return week[day]
}
