package permutation

func IsPermutation(src, txt string) bool {
	srcCounts := make(map[rune]int)
	txtCounts := make(map[rune]int)

	for _, r := range src {
		srcCounts[r]++
	}

	for _, r := range txt {
		txtCounts[r]++
	}

	if len(srcCounts) != len(txtCounts) {
		return false
	}

	for k, v := range srcCounts {
		if v != txtCounts[k] {
			return false
		}
	}
	return true
}
