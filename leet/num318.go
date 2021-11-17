package leet

func valid(a, b string)  bool{
	ma := make(map[rune]bool)
	for _, i := range a {
		ma[i] = true
	}
	for _, i :=  range b {
		if _, ok := ma[i]; ok {
			return false
		}
	}
	return true
}

func maxProduct0(words []string) int {
	ans := 0
	for i:=0; i < len(words) - 1; i++ {
		for j:= i+1; j < len(words); j++ {
			if i == j || !valid(words[i], words[j]){
				continue
			}
			if len(words[i]) * len(words[j]) > ans {
				ans = len(words[i]) * len(words[j])
			}
		}
	}
	return ans
}

func stringToMap(s string) map[byte]int {
	m :=  make(map[byte]int, 0)
	for i:=0 ; i < len(s); i++ {
		m[s[i]] += 1
	}
	return m
}

func valid2(a, b map[byte]int)  bool{
	for c, _ := range a {
		if _, ok := b[c]; ok {
			return false
		}
	}
	return true
}
func maxProduct2(words []string) int {
	ans := 0
	wordMaps := make([]map[byte]int, len(words))
	for i := range words {
		wordMaps[i] = stringToMap(words[i])
	}
	for i:=0; i < len(wordMaps) - 1; i++ {
		for j:= i+1; j < len(wordMaps); j++ {
			if i == j || !valid2(wordMaps[i], wordMaps[j]){
				continue
			}
			if len(words[i]) * len(words[j]) > ans {
				ans = len(words[i]) * len(words[j])
			}
		}
	}
	return ans
}

//func maxProduct(words []string) (ans int) {
//	masks := make([]int, len(words))
//	for i, word := range words {
//		for _, ch := range word {
//			masks[i] |= 1 << (ch - 'a')
//		}
//	}
//
//	for i, x := range masks {
//		for j, y := range masks[:i] {
//			if x&y == 0 && len(words[i])*len(words[j]) > ans {
//				ans = len(words[i]) * len(words[j])
//			}
//		}
//	}
//	return
//}
//o(n^2)
func maxProduct3(words []string) (ans int) {
	masks := make([]int, len(words))
	for i := range words {
		for j := range words[i] {
			masks[i] |= 1<< (words[i][j] - 'a')
		}
	}
	for i:= 0; i < len(masks) - 1; i++ {
		for j:= i+1; j < len(masks); j++{
			if l:= len(words[i]) * len(words[j]); masks[i] & masks[j] == 0 && l > ans {
				ans = l
			}
		}
	}
	return
}
//func maxProduct(words []string) (ans int) {
//   masks := map[int]int{}
//   for _, word := range words {
//       mask := 0
//       for _, ch := range word {
//           mask |= 1 << (ch - 'a')
//       }
//       if len(word) > masks[mask] {
//           masks[mask] = len(word)
//       }
//   }
//
//   for x, lenX := range masks {
//       for y, lenY := range masks {
//           if x&y == 0 && lenX*lenY > ans {
//               ans = lenX * lenY
//           }
//       }
//   }
//   return
//}


func maxProduct(words []string) (ans int) {
	masks := map[int]int{}
	for i := range words {
		m := 0
		for j := range words[i] {
			m |= 1<< (words[i][j] - 'a')
		}
		if masks[m] < len(words[i]) {
			masks[m] = len(words[i])
		}
	}
	for x, lenX := range masks {
		for y, lenY := range masks {
			if l:= lenX * lenY; x & y== 0 && l > ans {
				ans = l
			}
		}
	}
	return
}
