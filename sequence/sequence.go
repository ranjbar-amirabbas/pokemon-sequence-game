package sequence

import (
	"bufio"
	"io"
	"os"
	"strings"
	"unicode"
)

func Sequence(a rune, file io.ReadCloser) io.Reader {
	//guarantees the file closes when the function exits
	defer file.Close()

	f, ok := file.(*os.File)
	if !ok {
		return strings.NewReader("Error on reading the file")
	}

	currentLetter := unicode.ToLower(a)
	var result []string
	used := map[string]bool{}

	for {

		word := findWord(currentLetter, f, used)
		if word == "" {
			break
		}
		result = append(result, word)
		currentLetter = rune(word[len(word)-1])

	}

	return strings.NewReader(strings.Join(result, ", "))
}

func findWord(letter rune, file *os.File, used map[string]bool) string {

	if _, err := file.Seek(0, io.SeekStart); err != nil {
		return ""
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := strings.ToLower(strings.Trim(scanner.Text(), " ,."))
		if used[word] {
			continue
		}

		if rune(word[0]) == letter {
			used[word] = true
			return word
		}
	}
	return ""

}
