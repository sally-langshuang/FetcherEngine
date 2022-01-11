package dp

import "fmt"

const (
	size = 10e9 + 7
)

func setPaths(locations []int, start int, finish int, fuel int, f *map[int]map[int]map[int]int) (ans int) {
	if start == 2 && finish == 2 && fuel == 4 {
		fmt.Println("a")
	}
	if fuel < 0 || fuel < absSub(locations[start], locations[finish]) {
		return 0
	}
	if _, ok := (*f)[fuel]; ok {
		if _, ok := (*f)[fuel][start]; ok {
			if _, ok := (*f)[fuel][start][finish]; ok {
				return (*f)[fuel][start][finish]
			}
		} else {
			(*f)[fuel][start] = map[int]int{}
		}
	} else {
		(*f)[fuel] = map[int]map[int]int{}
		(*f)[fuel][start] = map[int]int{}
	}
	if start == finish && fuel >= 0 {
		ans++
	}
	for next := range locations {
		if next != start {
			v := setPaths(locations, next, finish, fuel-absSub(locations[start], locations[next]), f)
			if v > 0 {
				ans += v
			}
		}
	}
	ans %= size
	(*f)[fuel][start][finish] = ans
	fmt.Printf("%d %d %d ==> %d\n", start, finish, fuel, ans)
	return ans
}

func countRoutes(locations []int, start int, finish int, fuel int) (ans int) {
	f := &map[int]map[int]map[int]int{}
	for i := 0; i <= fuel; i++ {
		setPaths(locations, start, finish, fuel, f)
	}
	return (*f)[fuel][start][finish]
}

func absSub(a, b int) int {
	if a >= b {
		return a - b
	}
	return b - a
}
