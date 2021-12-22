package handle_str

func repeatedStringMatch(a string, b string) int {
loop:
	for offset := range a {
		var x int
		for i := range b {
			j := (offset + i) % len(a)
			if b[i] != a[j] {
				continue loop
			}
		}
		x += (len(b) + offset) / len(a)
		if (len(b)+offset)%len(a) > 0 {
			x++
		}
		return x

	}
	return -1
}
