package handle_str

import (
	"fmt"
	"strings"
)

//"adceb", "*a*b",

func printBoolArr(arr [][]bool)  {
	b := &strings.Builder{}
	for i:= range arr {
		for j := range arr[i] {
			if arr[i][j] {
				fmt.Fprintf(b, "1")
			} else {
				fmt.Fprintf(b, "0")
			}
		}
		fmt.Fprintf(b, "\n")
	}
	fmt.Println(b.String())
}
func isMatch(s string, p string) bool {
	f := make([][]bool, 2)
	for i := range f {
		f[i] = make([]bool, len(s) + 1)
	}
	f[0][0] = true
	for i := 1; i <= len(p); i++{
		cur := i % 2
		pre := cur ^ 1
		for j := 0; j <= len(s); j++{
			if p[i - 1] != '*'{
				if j > 0 {
					f[cur][j] = f[pre][j - 1] && (p[i - 1] == s[j - 1] || p[i - 1] == '?')
				} else {
					f[cur][j] = false
				}
			} else {
				f[cur][j] = f[pre][j]
				if j > 0 {
					f[cur][j] = f[cur][j] || f[cur][j - 1]
				}
			}
		}
	}
	return f[len(p) % 2][len(s)]
}
