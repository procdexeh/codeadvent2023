package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func part_one() {
	// open file
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file word by word using scanner
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	total := 0
	for scanner.Scan() {
		// do something with a word
		word := scanner.Text()

		var firstDigit, lastDigit int
		for _, char := range word {
			if char >= '0' && char <= '9' {
				digit, err := strconv.Atoi(string(char))
				if err == nil {
					if firstDigit == 0 {
						firstDigit = digit
					}
					lastDigit = digit
				}
			}
		}

		// Combine the first and last digits to form a 2-digit number
		twoDigitNumber := firstDigit*10 + lastDigit

		// Print the result
		fmt.Printf("First Digit: %d\n", firstDigit)
		fmt.Printf("Last Digit: %d\n", lastDigit)
		fmt.Printf("Combined 2-digit Number: %d\n", twoDigitNumber)

		// add the 2 digit to the total
		total += twoDigitNumber
	}
	fmt.Printf("Total: %d\n", total)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func word_to_number(word string) int {
	switch word {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	default:
		return 0

	}
}

func part_two() {
	// open file
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file word by word using scanner
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	numbers := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	total := 0
	for scanner.Scan() {
		// do something with a word
		word := scanner.Text()
		currentNumbers := make([]int, 0)
		// var firstDigit, lastDigit int
		for index, char := range word {
			if unicode.IsDigit(char) {
				digit, err := strconv.Atoi(string(char))
				if err == nil {
					currentNumbers = append(currentNumbers, digit)
				}
			}
			if unicode.IsLetter(char) {
				// go to next number.
				currentNumberWord := make([]rune, 0)
				for i := index; i < len(word); i++ {
					if unicode.IsDigit(rune(word[i])) {
						break
					}
					currentNumberWord = append(currentNumberWord, rune(word[i]))
					for _, number := range numbers {
						if number == string(currentNumberWord) {
							currentNumbers = append(currentNumbers, word_to_number(string(currentNumberWord)))
						}
					}
				}
			}
		}
		fmt.Printf("Current Numbers: %v\n", currentNumbers)
		var firstDigit, lastDigit int
		if len(currentNumbers) == 1 {
			firstDigit = currentNumbers[0]
			lastDigit = currentNumbers[0]
		} else {
			firstDigit = currentNumbers[0]
			lastDigit = currentNumbers[len(currentNumbers)-1]
		}
		//Combine the first and last digits to form a 2-digit number
		twoDigitNumber := firstDigit*10 + lastDigit

		// // Print the result
		fmt.Printf("First Digit: %d\n", firstDigit)
		fmt.Printf("Last Digit: %d\n", lastDigit)
		fmt.Printf("Combined 2-digit Number: %d\n", twoDigitNumber)

		// // add the 2 digit to the total
		total += twoDigitNumber
	}
	fmt.Printf("Total: %d\n", total)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func main() {
	part_two()
}
