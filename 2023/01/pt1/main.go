package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

type inputLine struct {
	string string
	value  int
}

func extractDigits(s string) ([]int, error) {
	var digits []int
	for _, char := range s {
		if unicode.IsDigit(char) {
			digit, err := strconv.Atoi(string(char))
			if err != nil {
				return nil, err
			}

			digits = append(digits, digit)
		}
	}

	return digits, nil
}

func concatFirstLast(digits []int) int {
	if len(digits) == 0 {
		return 0
	}

	first, last := digits[0], digits[len(digits)-1]
	return first*10 + last // hackerman int concatenation
}

func processLine(s string) (inputLine, error) {
	digits, err := extractDigits(s)
	if err != nil {
		return inputLine{}, err
	}

	value := concatFirstLast(digits)
	return inputLine{string: s, value: value}, nil
}

func main() {
	file, err := os.Open("../input")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var totalSum int

	for scanner.Scan() {
		line, err := processLine(scanner.Text())
		if err != nil {
			log.Fatalf("Error processing line: %v", err)
		}

		totalSum += line.value
		// fmt.Printf("Line: %s, Value: %d\n", line.string, line.value)
		// fmt.Printf("Sum so far: %d\n\n", totalSum)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading file:", err)
	}

	fmt.Println("Sum of calibration values:", totalSum)
}
