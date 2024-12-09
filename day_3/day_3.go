package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func readFileToString(filePath string) (string, error) {
	// Read the entire file into memory
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	// Convert the content to a string and return
	return string(content), nil
}

func computeMulResults(input string) int {
	// Define the regex pattern with capturing groups for numbers
	pattern := `mul\((\d+),(\d+)\)`
	re := regexp.MustCompile(pattern)

	// Find all matches with submatches
	matches := re.FindAllStringSubmatch(input, -1)

	// Var to store result
	result := 0

	// Iterate through the matches
	for _, match := range matches {
		if len(match) == 3 { // Ensure there are capturing groups
			num1, _ := strconv.Atoi(match[1]) // Convert first number to int
			num2, _ := strconv.Atoi(match[2]) // Convert second number to int

			// Compute the product
			product := num1 * num2

			// Store the result with the match as the key
			result += product
		}
	}

	return result
}

func main() {
	// Specify the input file
	filename := "day_3/input.txt"

	//part 1

	// Call the function to read the file
	content, err := readFileToString(filename)
	if err != nil {
		log.Fatalf("Failed to read the file: %v", err)
	}
	// Extract all patterns and compute result
	result := computeMulResults(content)
	fmt.Println(result)

}
