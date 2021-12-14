package dp

import "fmt"

func max(a, b int) int{
	if a >= b {
		return a
	}
	return b
}

func dp(courses [][]int, i, offset int, f*[]map[int]int) int {
	if i >= len(courses) {
		return 0
	}
	if v, ok := (*f)[i][offset]; ok {
		return v
	}
	if len(courses) == i + 1 {
		if courses[i][1] >= offset + courses[i][0] {
			(*f)[i][offset] = 1
			return 1
		} else {
			(*f)[i][offset] = 0
			return 0
		}
	}
	yes := 1 + dp(courses, i + 1, offset + courses[i][0], f)
	no := dp(courses, i + 1, offset, f)
	if courses[i][1] >= offset + courses[i][0] {
		(*f)[i][offset] = max(yes, no)
		return (*f)[i][offset]
	}
	(*f)[i][offset] = no
	return (*f)[i][offset]
}

func isAsc(a, b[]int) bool {
	if a[1] < b[1] {
		return true
	}
	if a[1] == b[1] && a[0] <= b[0] {
		return true
	}
	return false
}
func sort(courses [][]int)  {
	if len(courses) <= 1 {
		return
	}
	i, j:= 0, len(courses) - 1
	for i < j {
		for isAsc(courses[i], courses[j]) {
			j--
			if j == i {
				goto out
			}
		}
		courses[i], courses[j] = courses[j], courses[i]
		for isAsc(courses[i], courses[j]){
			i++
			if j == i {
				goto out
			}
		}
		courses[i], courses[j] = courses[j], courses[i]
	}
	out:
	sort(courses[:i])
	sort(courses[i + 1:])
	return

}
func scheduleCourse(courses [][]int) int {
	sort(courses)
	f := make([]map[int]int, len(courses))
	for i:= range  f {
		f[i] = map[int]int{}
	}
	res :=  dp(courses, 0, 0, &f)
	fmt.Println(f)
	return  res
}