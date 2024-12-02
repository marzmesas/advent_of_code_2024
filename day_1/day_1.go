package day_1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func read_input_lists(filename string) ([]int, []int, error) {

	// Open the input file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		return nil, nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	// Initialize slices for the two columns
	var column1 []int
	var column2 []int

	// Read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Split the line into two parts
		parts := strings.Fields(line)
		if len(parts) != 2 {
			fmt.Printf("Skipping malformed line: %v\n", line)
			continue
		}

		// Convert parts to integers
		num1, err1 := strconv.Atoi(parts[0])
		num2, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			fmt.Printf("Skipping line with invalid numbers: %v\n", line)
			continue
		}

		// Append to respective columns
		column1 = append(column1, num1)
		column2 = append(column2, num2)
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return nil, nil, fmt.Errorf("error reading file: %w", err)
	}
	return column1, column2, nil
}
