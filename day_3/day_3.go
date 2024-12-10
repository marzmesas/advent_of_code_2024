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

func processMulInstructions(input string) int {
	// Regex patterns for mul(num1,num2), do(), and don't()
	mulPattern := `mul\((\d+),(\d+)\)`
	doPattern := `do\(\)`
	dontPattern := `don't\(\)`

	// Compile regexes
	mulRegex := regexp.MustCompile(mulPattern)
	doRegex := regexp.MustCompile(doPattern)
	dontRegex := regexp.MustCompile(dontPattern)

	// Track whether mul operations are enabled
	mulEnabled := true

	// Map to store results of executed mul operations
	result := 0

	// Match all instructions (both mul and do/don't)
	allInstructions := regexp.MustCompile(fmt.Sprintf(`%s|%s|%s`, mulPattern, doPattern, dontPattern))
	matches := allInstructions.FindAllStringSubmatch(input, -1)

	// Iterate through all matched instructions in order
	for _, match := range matches {
		fullMatch := match[0] // Full matched string

		switch {
		case doRegex.MatchString(fullMatch): // Handle do()
			mulEnabled = true

		case dontRegex.MatchString(fullMatch): // Handle don't()
			mulEnabled = false

		case mulRegex.MatchString(fullMatch): // Handle mul(num1,num2)
			if mulEnabled {
				// Extract the numbers and compute the result
				num1, _ := strconv.Atoi(match[1])
				num2, _ := strconv.Atoi(match[2])
				result += num1 * num2
			}
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

	//part 2
	result2 := processMulInstructions(content)
	fmt.Printf("Part 2: %v\n", result2)
}
