package leet

type stack struct {
	Arr []int
	Tail int
}
func (s *stack) empty() bool {
	return s.Tail == 0
}

func (s *stack) push(val int) {
	if s.Tail == len(s.Arr) {
		a := make([]int, 2 * len(s.Arr)+1)
		for i := range s.Arr{
			  a[i] = s.Arr[i]
		}
		s.Arr = a
	}
	s.Arr[s.Tail] = val
	s.Tail++
}

func (s *stack) pop() (int, bool) {
	if s.empty() {
		return 0, false
	}
	if len(s.Arr) >= 2 && s.Tail < len(s.Arr) / 2{
		s.Arr = s.Arr[ :len(s.Arr) / 2]
	}
	s.Tail--
	return s.Arr[s.Tail], true
}
func (s *stack) peek() (int, bool) {
	if s.empty() {
		return 0, false
	}
	return s.Arr[s.Tail - 1], true
}
func (s *stack) clear() {
	s.Tail = 0
}
func longestValidParentheses0(s string) int {
	st := stack{Arr: []int{}}
	doStr := make([]int, len(s))
	for i:=range s {
		if s[i] == '(' {
			st.push(i)
		}  else {
			v, ok := st.pop()
			if ok {
				doStr[v] = 1
				doStr[i] = 1
			} else {
				continue
			}
		}
	}
	max, count := 0, 0
	for i := range doStr {
		  if doStr[i] == 1 {
			  count++
			  if count > max {
				  max = count
			  }
		  }  else {
			  count = 0
		  }
	}
	return max
}

func longestValidParentheses(s string) int {
	arr := make([]int, len(s))
	max := 0
	for i := 1; i < len(arr); i++ {
		if  s[i] == '(' || i - arr[i - 1] - 1 < 0 || s[i - arr[i - 1] - 1] != '(' {
			continue
		}
		arr[i] = arr[i - 1] + 2
		if i - arr[i]  >= 0 {
			arr[i] += arr[i - arr[i] ]
		}
		if arr[i] > max {
			max = arr[i]
		}

	}
	return max
}