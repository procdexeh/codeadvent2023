package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type viableGame struct {
	binned bool
}

func find_digit(str string) int {
	for i, char := range str {
		if char >= '0' && char <= '9' {
			digit, err := strconv.Atoi(string(str[i:]))
			if err != nil {
				log.Fatal(err)
			}
			return digit
		}
	}
	return 0
}

func parse_subgame(game []string, cubeTracker *viableGame) *viableGame {
	for _, subset := range game {
		subgame := strings.Split(subset, ",")
		for _, cube := range subgame {
			trimmedCube := strings.ReplaceAll(cube, " ", "")
			cubeValue := 0
			// get value of the cube
			var colorString string
			for i, char := range trimmedCube {
				if unicode.IsDigit(char) {
					// search for where digits end
					for j := i; j < len(trimmedCube); j++ {
						if !unicode.IsDigit(rune(trimmedCube[(j + 1)])) {
							digit, err := strconv.Atoi(string(trimmedCube[0 : j+1]))
							if err != nil {
								log.Fatal(err)
							}
							colorString = string(trimmedCube[(j + 1):])
							if digit > cubeValue {
								cubeValue = digit
							}
							break
						}
					}
					break
				}
			}
			// println("value: ", cubeValue, "color: ", colorString)
			switch colorString {
			case "red":
				if cubeValue >= 12 {
					cubeTracker.binned = true
				}
				break
			case "green":
				if cubeValue >= 13 {
					cubeTracker.binned = true
				}
				break
			case "blue":
				if cubeValue >= 14 {
					cubeTracker.binned = true
				}
				break
			}
		}
	}
	return cubeTracker
}

func k() {
	// open file
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file word by word using scanner
	scanner := bufio.NewScanner(f)

	totalViableGames := 0
	for scanner.Scan() {
		// do something with a word
		gameTracker := viableGame{binned: false}
		word := scanner.Text()
		splitWord := strings.Split(word, ":")
		gameId := find_digit(splitWord[0])
		game := splitWord[1]
		subsets := strings.Split(game, ";")
		gameTracker = *parse_subgame(subsets, &gameTracker)
		fmt.Println("gameID: ", gameId, "binned: ", gameTracker.binned)
		if !gameTracker.binned {
			totalViableGames += gameId
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("viable game sum: ", totalViableGames)
}
func main() {
	k()
}
