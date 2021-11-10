package tree

import (
	"fmt"
	"strings"
)

type Palindrome struct {
	center, armLen int
}

func (p Palindrome) RightHand() int {
	return p.center + p.armLen - 1
}
func (p Palindrome) Mirror(i int) int {
	return 2*p.center - i
}
func (p Palindrome) Len() int {
	return p.armLen - 1
}
func (p Palindrome) S(s string) string {
	b := strings.Builder{}
	for i:= p.center - p.armLen + 1; i < p.center + p.armLen - 1;i++ {
		if s[i] != '#' {
			fmt.Fprint(&b, string(s[i]))
		}
	}
	return b.String()
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
func produceS(s string) string{
	if s == "" {
		return s
	}
	b := strings.Builder{}
	fmt.Fprint(&b, "#")
	for _, r := range s {
		fmt.Fprint(&b, string(r), "#")
	}
	return b.String()
}

func longestPalindrome(s string) string {
	s = produceS(s)
	arm := make([]int, len(s))
	rightMost := Palindrome{center: 0, armLen: 1}
	longest := Palindrome{center: 0, armLen: 1}
	for i, _ := range s {
		startArm := 1
		if rightMost.RightHand() > i {
			startArm = min(arm[rightMost.Mirror(i)], rightMost.RightHand() - i + 1)
		}
		for {
			if i + startArm < len(s) && i - startArm >= 0 && s[i + startArm] == s[i - startArm] {
				startArm++
			} else {
				break
			}
		}
		if startArm + i - 1 > rightMost.RightHand() {
			rightMost = Palindrome{center: i, armLen: startArm}
		}
		if (startArm - 1) > longest.Len() {
			longest = Palindrome{center: i, armLen: startArm}
		}
	}
	return longest.S(s)
}