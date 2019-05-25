package unique

func IsUnique(str string) bool {
	runes := []rune(str)
	for i, r := range runes {
		for j := i + 1; j < len(runes); j++ {
			if r == runes[j] {
				return false
			}
		}
	}
	return true
}
