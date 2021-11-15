package leet

import (
	"strings"
)

//relation := map[string]int{
//"I": 1,
//"IV": 4,
//"V": 5,
//"IX": 9,
//"X": 10,
//"XL": 40,
//"L": 50,
//"XC": 90,
//"C": 100,
//"CD": 400,
//"D": 500,
//"CM": 900,
//"M": 1000,
//}

//"I": 1,
//"V": 5,
//"X": 10,
//"L": 50,
//"C": 100,
//"D": 500,
//"M": 1000,
//1994
func intToRoman(num int) string {
	ans :=""
	a := []string{"I", "V", "X", "L", "C", "D", "M"}
	for i:=0; num > 0 ; i++ {
		n := num % 10
		if n <= 3 {
			ans = strings.Repeat(a[i*2], n) + ans
		} else if n == 4{
			ans = a[i*2] + a[i*2 + 1]  + ans
		} else if n <= 8 {
			ans = a[i*2 + 1] + strings.Repeat(a[i*2], n - 5) + ans
		} else{
			ans = a[i*2] + a[i*2 + 2] + ans
		}
		num /= 10
	}
	return ans

}
func romanToInt(s string) int {
//LVIII 58
	var ans int
	m := map[uint8]int{'I':1, 'V':5, 'X':10, 'L':50, 'C':100, 'D':500, 'M':1000}
	for i:= 0; i < len(s); i++ {
		if i + 1 < len(s) && m[s[i+1]] > m[s[i]]{
			ans -= m[s[i]]
		} else {
			ans += m[s[i]]
		}
	}
	return ans
}