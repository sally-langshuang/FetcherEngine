package handle_str

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNum686(t *testing.T) {
	datas := []struct {
		a, b string
		ans  int
	}{
		{"abcd", "cdabcdab", 3},
		{"a", "aa", 2},
		{"a", "a", 1},
		{"abc", "wxyz", -1},
	}
	for _, x := range datas {
		fmt.Println(repeatedStringMatch(x.a, x.b) == x.ans)
	}
}
func TestNum1154(t *testing.T) {
	dayOfYear("2004-03-01")
}
func TestNum748(t *testing.T) {
	datas := []struct {
		licensePlate, ans string
		words             []string
	}{
		//{"iMSlpe4", "simple", []string{"claim","consumer","student","camera","public","never","wonder","simple","thought","use"}},
		//{"OgEu755", "enough", []string{"enough","these","play","wide","wonder","box","arrive","money","tax","thus"}},
		//{"Ah71752", "husband", []string{"suggest","letter","of","husband","easy","education","drug","prevent","writer","old"}},
		//{"1s3 456", "pest", []string{"looks", "pest", "stew", "show"}},
		//{"1s3 PSt", "steps", []string{"step", "steps", "stripe", "stepple"}},
		{"AN87005", "important", []string{"participant", "individual", "start", "exist", "above", "already", "easy", "attack", "player", "important"}},
	}
	for _, d := range datas {
		fmt.Println(shortestCompletingWord(d.licensePlate, d.words) == d.ans)
	}
}
func TestNum44(t *testing.T) {
	datas := []struct {
		s, p string
		ans  bool
	}{
		{"abcabczzzde", "*abc???de*", true},
		{"zacabz", "*a?b*", false},
		{"acdcb", "a*c?b", false},
		{"adceb", "*a*b", true},
		{"cb", "?a", false},
		{"aa", "*", true},
		{"aa", "a", false},
		{"", "*", true},
		{"", "", true},
	}
	for _, x := range datas {
		fmt.Println(isMatch(x.s, x.p) == x.ans)
	}
}
func TestNum43(t *testing.T) {
	datas := []struct {
		num1, num2, ans string
	}{
		{"12", "3", "36"},
		{"9", "9", "81"},
		{"123", "456", "56088"},
		{"123", "0", "0"},
	}
	for _, x := range datas {
		fmt.Println(multiply(x.num1, x.num2) == x.ans)
	}
}
func TestNum1446(t *testing.T) {
	datas := []struct {
		s   string
		ans int
	}{
		{"j", 1},
		{"cc", 2},
		{"tourist", 1},
		{"hooraaaaaaaaaaay", 11},
		{"triplepillooooow", 5},
		{"abbcccddddeeeeedcba", 5},
		{"leetcode", 2},
	}
	for _, x := range datas {
		fmt.Println(reflect.DeepEqual(maxPower(x.s), x.ans))
	}
}
