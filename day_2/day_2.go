package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readMatrix(filePath string) ([][]int, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var matrix [][]int // Slice of slices to hold the matrix

	// Use a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()           // Read a line
		numStrs := strings.Fields(line)  // Split the line into fields
		var row []int                    // Slice to hold a row
		for _, numStr := range numStrs { // Convert each field to an integer
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return nil, fmt.Errorf("invalid number in file: %w", err)
			}
			row = append(row, num)
		}
		matrix = append(matrix, row) // Add the row to the matrix
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return matrix, nil
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// part 1
func checkIncreasing(row []int) bool {
	for i := 0; i < len(row)-1; i++ {
		j := i + 1
		if absInt(row[i]-row[j]) > 3 || absInt(row[i]-row[j]) < 1 {
			return false
		}
		if row[j] < row[i] {
			return false
		}

	}
	return true
}

func checkDecreasing(row []int) bool {

	for i := 0; i < len(row)-1; i++ {
		j := i + 1
		if absInt(row[i]-row[j]) > 3 || absInt(row[i]-row[j]) < 1 {
			return false
		}
		if row[j] > row[i] {
			return false
		}

	}
	return true
}

func checkSafety(row []int) bool {
	i := 0
	j := 1
	var result bool
	if absInt(row[i]-row[j]) > 3 || absInt(row[i]-row[j]) < 1 {
		return false
	}
	if row[i] < row[j] {
		slice := row[j:]
		result = checkIncreasing(slice)
	} else if row[i] > row[j] {
		slice := row[j:]
		result = checkDecreasing(slice)
	} else {
		result = false
	}
	return result
}

// part 2
// Check if the row is safe in the given mode (increasing or decreasing)
func isSafe(row []int, mode string) bool {
	for i := 0; i < len(row)-1; i++ {
		diff := row[i+1] - row[i]
		if absInt(diff) < 1 || absInt(diff) > 3 {
			return false
		}
		if (mode == "increasing" && diff < 0) || (mode == "decreasing" && diff > 0) {
			return false
		}
	}
	return true
}

// Check if a row can be made safe by removing one element
func canBeMadeSafe(row []int) bool {
	for i := 0; i < len(row); i++ {
		// Create a new slice excluding the current element (ensure no accidental modification)
		newRow := make([]int, len(row)-1)
		copy(newRow[:i], row[:i])
		copy(newRow[i:], row[i+1:])

		if isSafe(newRow, "increasing") || isSafe(newRow, "decreasing") {
			fmt.Printf("Row %v becomes safe by removing index %d: %v\n", row, i, newRow)
			return true
		}
	}
	return false
}

// Main safety check function
func checkSafetyTwoPointer(row []int) bool {
	fmt.Printf("Checking row: %v\n", row)
	// First check if the row is safe as-is
	if isSafe(row, "increasing") || isSafe(row, "decreasing") {
		fmt.Printf("Row %v is safe as-is\n", row)
		return true
	}
	// If not safe, check if it can be made safe by removing one element
	if canBeMadeSafe(row) {
		fmt.Printf("Row %v can be made safe by removing one element\n", row)
		return true
	}
	// Otherwise, the row is unsafe
	fmt.Printf("Row %v is unsafe\n", row)
	return false
}

// Main function
func main() {

	// Example usage
	filePath := "day_2/input_day2.txt" // Replace with your file path
	matrix, err := readMatrix(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// part 1
	var safeCounter = 0
	for _, row := range matrix {
		val := checkSafety(row)
		if val {
			safeCounter += 1
		}
	}
	fmt.Printf("Total number of safe reports: %v\n", safeCounter)

	safeCount := 0
	for _, row := range matrix {
		if checkSafetyTwoPointer(row) {
			safeCount++
		}
	}
	fmt.Printf("Total number of safe reports, part2: %v\n", safeCount)
}
