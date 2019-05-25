package urlify

import "fmt"

func URLify(url string, curLength int) string {
	input := []rune(url)

	for i := curLength - 1; i >= 0; i-- {
		if input[i] != ' ' {
			continue
		}

		for j := curLength - 1; j > i; j-- {
			input[j+2] = input[j]
		}
		input[i+2] = '0'
		input[i+1] = '2'
		input[i+0] = '%'

		curLength = curLength + 2
		fmt.Println(string(input))
	}

	return string(input)
}
