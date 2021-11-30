package leet
import (
	. "github.com/golang-collections/collections/stack"
)
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

func longestValidParentheses3(s string) int {
	arr := make([]int, len(s))
	max := 0
	for i := 1; i < len(arr); i++ {
		//s[i - arr[i - 1] - 1] 为 （ 需要匹配的字符，如果是）才会连续+2， 如果arr[i - arr[i]]不为零，需要追加上
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

func longestValidParentheses4(s string) int {
	st := New()
	st.Push(-1)
	maxL, l := 0, 0
	for i:= range s {
		if s[i] == '(' {
			st.Push(i)
			continue
		}
		// )  )
		peer := st.Pop().(int)
		if peer == -1 || s[peer] == ')' {
			st.Push(i)
			continue
		}
		// )  (
		l = i - st.Peek().(int)
		if l > maxL {
			maxL = l
		}
	}
	return maxL
}
//best
func longestValidParentheses(s string) int {
	left, right,  maxL := 0, 0, 0
	for i := range s {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			if left + right > maxL {
				maxL = left + right
			}
		} else if right > left {
			left = 0
			right = 0
		}
	}
	left, right = 0, 0
	for i := range s {
		if s[len(s) - 1 - i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			if left + right > maxL {
				maxL = left + right
			}
		} else if right < left {
			left = 0
			right = 0
		}
	}
	return maxL
}