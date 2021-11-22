package leet
var relations = []string{"",
	"", "abc", "def",
	"ghi", "jkl", "mno",
	"pqrs", "tuv", "wxyz",
}
func f17(i int, digits string, prefix []string) (ans []string){
	if i == len(digits) {
		return prefix
	}
	chars := relations[digits[i] - '0']
	if len(prefix) == 0 {
		prefix = []string{""}
	}
	ans = make([]string, 0)

	for x := range chars{
		for y := range prefix {
			ans = append(ans, prefix[y] +string(chars[x]))
		}
	}
	i++
	return f17(i, digits, ans)
}
func letterCombinations(digits string) []string {
	ans := make([]string, 0)
	return f17(0, digits, ans)
}
