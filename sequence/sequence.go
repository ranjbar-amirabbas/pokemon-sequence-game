package sequence

import (
	"bufio"
	"io"
	"strings"
	"unicode"
)

func Sequence(a rune, file io.ReadCloser) io.Reader {

	currentLetter := unicode.ToLower(a)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	result := []string{}

	index := map[rune][]string{}

	for scanner.Scan() {
		word := strings.ToLower(strings.Trim(scanner.Text(), " ,."))
		firstLetter := rune(word[0])
		index[firstLetter] = append(index[firstLetter], word)
	}
	for {
		targetList := index[currentLetter]
		if len(targetList) == 0 {
			break
		}
		target := targetList[0]

		result = append(result, target)
		index[currentLetter] = targetList[1:]
		targetRunes := []rune(target)
		currentLetter = targetRunes[len(targetRunes)-1]
	}

	return strings.NewReader(strings.Join(result, ", "))
}
