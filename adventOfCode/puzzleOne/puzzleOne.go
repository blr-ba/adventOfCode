// Advent of code
// Puzzle 1

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// DesiredOutput is what the sum should add up to
const DesiredOutput = 2020

// Read the numbers from the input file and store in an array
func readLinesFromFile(fileName string) []int {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var numList []int

	for scanner.Scan() {
		l, _ := strconv.Atoi(scanner.Text())
		numList = append(numList, l)
	}

	file.Close()

	return numList
}

// Given a list of numbers, find if the first number plus any of the following numbers add up to the total we want.
// If they do, return their multiplication.
func findTheAnswer(numList []int, desiredTotal int) int {
	a := numList[0]
	total := -1
	for _, b := range numList[1:] {
		if (a + b) == desiredTotal {
			total = a * b
			break
		}
	}
	return total
}

func main() {
	inputLines := readLinesFromFile("input.txt")

	//total := findTheAnswer(inputLines, DesiredOutput)
	//fmt.Println(total)

	var total int
	for i := range inputLines {
		total = findTheAnswer(inputLines[i:], DesiredOutput)
		if total != -1 {
			break
		}
	}

	fmt.Println(total)

}
