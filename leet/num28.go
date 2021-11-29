package leet

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if haystack == "" {
		return -1
	}
	str := needle + "#" + haystack
	f := make([]int, len(haystack) + len(needle) + 1)

	for i := 1;  i < len(str); i++ {
		l := f[i - 1]
		for ; l > 0; {
			if str[l] == str[i] {
				f[i] += l
				break
			} else {
				l = f[l - 1]

			}
		}
		if str[l] == str[i] {
			f[i] += 1
		}
		//fmt.Println(i, f[i], str[0:f[i]], str[i - f[i] + 1:i + 1], str[:i], str[i:])
		if f[i] == len(needle) {
			return i - 2 * len(needle)
		}
	}
	//fmt.Println(f)
	return -1
}

func strStr1(haystack string, needle string) int {
	if needle == "" {
		return 0
	} else if haystack == "" {
		return -1
	}
	for i := range haystack {
		if haystack[i] != needle[0] {
			continue
		}
		for j:=1; j < len(needle); j ++ {
			if i+j >= len(haystack) || needle[j] != haystack[i+j] {
				goto next
			}
		}
		return i
		next:
	}
	return -1
}
