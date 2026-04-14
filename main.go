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
			fmt.Println("Error in opeining file: ", err)
			return
		}
		var input string
		fmt.Scan(&input)

		result := sequence.Sequence(rune(input[0]), file)

		data, err := io.ReadAll(result)
		if err != nil {
			fmt.Println("Error on showing result", err)
			return
		}
		if len(data) == 0 {
			fmt.Println("no matching words")
		}

		fmt.Println(string(data))
	}
}
