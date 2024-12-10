package main

import (
	"bufio"
	"fmt"
	"os"
)

func readMatrix(filePath string) ([][]rune, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error closing file: %v\n", err)
		}
	}(file)

	var matrix [][]rune // Slice of slices to hold the matrix

	// Use a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text() // Read a line
		var row []rune         // Slice to hold a row
		for _, char := range line {
			row = append(row, char)
		}
		matrix = append(matrix, row) // Add the row to the matrix
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return matrix, nil
}

func main() {
	matrix, _ := readMatrix("day_4/input.txt")
	fmt.Println(string(matrix[0]))

}
