package main

import (
	"fmt"
	"io"
	"monstercat-pokemon-sequence/sequence"
	"os"
)

func main() {
	for true {
		file, err := os.Open("sequence/pokemon.txt")
		if err != nil {
			fmt.Println("Error in opening file: ", err)
			return
		}
		var input string
		_, err = fmt.Scan(&input)
		if err != nil || len(input) == 0 {
			break
		}
		result := sequence.Sequence(rune(input[0]), file)
		file.Close()
		data, err := io.ReadAll(result)
		if err != nil {
			fmt.Println("Error on showing result", err)
			return
		}
		if len(data) == 0 {
			fmt.Println("no matching words")
		} else {
			fmt.Println(string(data))
		}

	}
}
