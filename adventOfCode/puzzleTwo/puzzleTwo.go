// Advent of code
// Puzzle 2

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type passwordSpec struct {
	min      int
	max      int
	letter   string
	password string
}

// Read the password specification from the file, return them in an array
func readLinesFromFile(fileName string) []passwordSpec {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var passwordSpecList []passwordSpec

	for scanner.Scan() {
		//l, _ := strconv.Atoi(scanner.Text())

		ps := passwordSpec{}
		fmt.Sscanf(scanner.Text(), "%d-%d %s %s", &ps.min, &ps.max, &ps.letter, &ps.password)
		// Hack to strip the trailing colon
		ps.letter = strings.TrimSuffix(ps.letter, ":")
		passwordSpecList = append(passwordSpecList, ps)
	}

	file.Close()

	return passwordSpecList
}

// A password is valid if the number of times it contains the letter is >= min and <= max
func isValidPassword(password passwordSpec) bool {
	letterCount := strings.Count(password.password, password.letter)
	return (letterCount >= password.min) && (letterCount <= password.max)
}

// Alternate validation
// Password is valid if the letter is either in the 'min' position or the 'max' position, but not both
func isValidPasswordAlt(password passwordSpec) bool {
	passwordRune := []rune(password.password)
	firstLetter := string(passwordRune[password.min-1])
	secondLetter := string(passwordRune[password.max-1])

	return (firstLetter == password.letter) != (secondLetter == password.letter)
}

func main() {
	inputPasswordSpec := readLinesFromFile("input.txt")

	var validPasswords = 0

	for i := range inputPasswordSpec {
		fmt.Println(inputPasswordSpec[i])
		if isValidPasswordAlt(inputPasswordSpec[i]) {
			validPasswords++
		}
	}

	fmt.Println(validPasswords)

}
