package leet

func getNextID1(words []string) []int {
	// []string{"wd","gd","wd", "wd"} ==> []int{3, 0, 4, 0}
	nextNum := make([]int, len(words))
next:
	for i := range words {
		for j:= i + 1; j < len(words);j++ {
			if words[j] == words[i] {
				nextNum[i] = j + 1
				continue next
			}
		}
	}
	return nextNum
}

func getFirstWordNum1(s string, words []string) []int {
	firstWordID := make([]int, len(s) - len(words[0]) + 1)
c:
	for i := 0; i <= len(s) - len(words[0]); i++ {
		w := s[i: i+len(words[0])]
		for j := range words {
			if words[j] == w {
				firstWordID[i] = j + 1
				continue c
			}
		}
	}
	return firstWordID
}

func findAFreeWord1(nextNum, firstWordID[]int, i int, currentMarkState *[]int) bool{
	num := firstWordID[i]
	if num == 0 {
		return false
	}
	for num > 0 && (*currentMarkState)[num - 1] == 1 && nextNum[num - 1] != 0 {
		num = nextNum[num - 1]
	}
	if (*currentMarkState)[num - 1] == 1 {
		return false
	}
	(*currentMarkState)[num - 1] = 1
	return true
}
func findSubstring1(s string, words []string) []int {
	ans := make([]int, 0)
	//"wdwdgdwdwdgdwdwd"  []string{"wd","gd","wd", "wd"}
	if len(words) == 0 || len(s) < len(words) * len(words[0]) {
		return []int{}
	}
	nextNum := getNextID1(words)//[3 0 4 0]
	firstWordID := getFirstWordNum1(s, words)//[1 0 1 0 2 0 1 0 1 0 2 0 1 0]
n:
	for i := 0; i <= len(s) - len(words) * len(words[0]); i++ {
		if firstWordID[i] == 0 {
			continue n
		}
		marks := make([]int, len(words))
		for j:=0; j < len(words) && i + j*len(words[0]) < len(firstWordID) ; j += 1 {
			if ok := findAFreeWord1(nextNum, firstWordID, i + j*len(words[0]), &marks); !ok {
				continue n
			}
		}
		toAdd := true
		for x:= range marks {
			if marks[x] == 0 {
				toAdd = false
			}
		}
		if toAdd {
			ans = append(ans, i)
		}
	}
	return ans
}

