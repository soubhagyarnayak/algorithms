package datastructure

func FindSubstring(s, substring string) int {
	for i := 0; i <= len(s)-len(substring); i++ {
		if s[i] == substring[0] {
			j := 1
			for ; j < len(substring); j++ {
				if s[i+j] != substring[j] {
					break
				}
			}
			if j == len(substring) {
				return i
			}
		}
	}
	return -1
}
