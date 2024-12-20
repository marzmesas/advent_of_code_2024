package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func readInputLists(filename string) ([]int, []int, error) {

	// Open the input file
	file, err := os.Open("day_1/input.txt")
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

func countOccurrences(slice []int, target int) int {
	count := 0
	for _, value := range slice {
		if value == target {
			count++
		}
	}
	return count
}

func main() {

	column1, column2, err := readInputLists("input.txt")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// part 1

	sort.Ints(column1)
	sort.Ints(column2)
	var output = 0
	var val = 0
	for i := 0; i < len(column1); i++ {
		val = column1[i] - column2[i]
		output += absInt(val)
	}

	fmt.Printf("Column 1: %v\n", len(column1))
	fmt.Printf("Column 2: %v\n", len(column2))
	fmt.Printf("Output: %v\n", output)

	// part 2
	var output2 = 0
	for i := 0; i < len(column1); i++ {
		val2 := countOccurrences(column2, column1[i]) //Revisit, right now big o is o(n)^2
		output2 += column1[i] * val2
	}
	fmt.Printf("Output 2: %v\n", output2)

}
