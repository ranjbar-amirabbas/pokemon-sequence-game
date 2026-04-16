package sequence

import (
	"bufio"
	"io"
	"strings"
	"unicode"
)

/*
Sequence does not call Close() on the provided reader.
The caller is responsible for opening and closing the stream,
as Sequence should not assume ownership of the resource it receives.
*/
func Sequence(a rune, file io.ReadCloser) io.Reader {

	/*
		Normalize to lowercase for case-insensitive matching
		The source data is already lowercase, so ToLower avoids
		unnecessary transformations compared to ToUpper
	*/
	currentLetter := unicode.ToLower(a)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	result := []string{}

	/*
		We build an index grouped by first letter so that each lookup is O(1)
		avoiding repeated scans of the stream.
		The stream is read once; don't need it again
	*/
	index := map[rune][]string{}

	for scanner.Scan() {
		// Normalize to lowercase for case-insensitive matching
		// Trim because words in the source are separated by commas(,)
		word := strings.ToLower(strings.Trim(scanner.Text(), " ,."))
		firstLetter := rune(word[0])
		index[firstLetter] = append(index[firstLetter], word)
	}

	// Each word's last letter becomes the key for the next lookup.
	// Words are removed from the index as they are used to prevent repetition.
	for {
		targetList := index[currentLetter]
		if len(targetList) == 0 {
			break
		}
		target := targetList[0]

		result = append(result, target)

		// Prevent repetition.
		index[currentLetter] = targetList[1:]

		// Advance to the last letter of the current word to find the next match.
		targetRunes := []rune(target)
		currentLetter = targetRunes[len(targetRunes)-1]
	}

	return strings.NewReader(strings.Join(result, ", "))
}
