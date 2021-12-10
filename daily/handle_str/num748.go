package handle_str

func shortestCompletingWord(licensePlate string, words []string) string {
	pattern := [26]int{}
	for _, i := range licensePlate {
		if i >= 'A' && i <= 'Z' {
			i += 'a' - 'A'
		}
		if i >= 'a' && i <= 'z' {
			pattern[i-'a'] += 1
		}
	}
	anss := []string{}
	loop:
	for _, word:= range words{
		match := [26]int{}
		for _, w := range word {
			if w >= 'A' && w <= 'Z' {
				w += 'a' - 'A'
			}
			match[w - 'a'] += 1
		}
		for i := 0; i < 26; i++{
			if pattern[i] != 0 && pattern[i] > match[i] {
				continue loop
			}
		}
		if len(anss) == 0 {
			anss = append(anss, word)
		} else if len(word) < len(anss[0]){
			anss[0] = word
		}

	}
	return anss[0]
}
