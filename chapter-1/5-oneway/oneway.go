package oneway

func IsOneWay(original, after string) bool {
	origLength, afterLength := len(original), len(after)

	switch {
	case origLength < afterLength: // addition
		j := 0
		for i := 0; i < afterLength; i++ {
			if original[j] == after[i] {
				j++
				continue
			}

			if j != i {
				return false
			}
		}
	case origLength == afterLength: // edit
		hadEdition := false
		for i := 0; i < origLength; i++ {
			if original[i] == after[i] {
				continue
			}

			if hadEdition {
				return false
			}
			hadEdition = true
		}
	case origLength > afterLength: // removal
		j := 0
		for i := 0; i < origLength; i++ {
			if j >= len(after) {
				return j == i
			}
			if original[i] == after[j] {
				j++
				continue
			}

			if j != i {
				return false
			}
		}
	}
	return true
}
