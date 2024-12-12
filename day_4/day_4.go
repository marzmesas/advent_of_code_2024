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

func countWordOccurrences(matrix [][]rune, word string) int {
	wordLen := len(word)
	directions := [][2]int{
		{-1, 0},  // Up
		{1, 0},   // Down
		{0, -1},  // Left
		{0, 1},   // Right
		{-1, -1}, // Up-Left
		{-1, 1},  // Up-Right
		{1, -1},  // Down-Left
		{1, 1},   // Down-Right
	}

	// Helper function to check if the word exists starting at (row, col) in the given direction
	isWordAt := func(row, col, dirX, dirY int) bool {
		for i := 0; i < wordLen; i++ {
			newRow := row + i*dirX
			newCol := col + i*dirY

			// Check bounds
			if newRow < 0 || newRow >= len(matrix) || newCol < 0 || newCol >= len(matrix[0]) {
				return false
			}

			// Check character match
			if matrix[newRow][newCol] != rune(word[i]) {
				return false
			}
		}
		return true
	}

	count := 0

	// Traverse the entire matrix
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			// Check all 8 directions
			for _, dir := range directions {
				if isWordAt(row, col, dir[0], dir[1]) {
					count++
				}
			}
		}
	}

	return count
}

func countXMASShapes(matrix [][]rune) int {
	count := 0
	rows := len(matrix)
	cols := len(matrix[0])

	// Helper function to check boundaries
	isValid := func(row, col int) bool {
		return row >= 0 && row < rows && col >= 0 && col < cols
	}

	// Helper function to check a diagonal pattern
	checkDiagonal := func(row, col, dRow1, dCol1, dRow2, dCol2 int) bool {
		// Check first diagonal (M->A->S or S->A->M)
		firstM := isValid(row+dRow1, col+dCol1) && matrix[row+dRow1][col+dCol1] == 'M'
		firstS := isValid(row+dRow1, col+dCol1) && matrix[row+dRow1][col+dCol1] == 'S'
		secondM := isValid(row+dRow2, col+dCol2) && matrix[row+dRow2][col+dCol2] == 'M'
		secondS := isValid(row+dRow2, col+dCol2) && matrix[row+dRow2][col+dCol2] == 'S'

		return (firstM && secondS) || (firstS && secondM)
	}

	// Traverse the matrix
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			// Check if 'A' is at the center of an X shape
			if matrix[row][col] == 'A' {
				// Check both diagonals for valid patterns
				if checkDiagonal(row, col, -1, -1, 1, 1) && // Top-Left to Bottom-Right
					checkDiagonal(row, col, -1, 1, 1, -1) { // Top-Right to Bottom-Left
					count++
				}
			}
		}
	}

	return count
}

func main() {
	matrix, _ := readMatrix("day_4/input.txt")
	word := "XMAS"

	// part 1

	count := countWordOccurrences(matrix, word)
	fmt.Printf("The word %q appears %d times in the matrix.\n", word, count)

	// part 2
	count2 := countXMASShapes(matrix)
	fmt.Printf("The X-MAS shapes found: %d\n", count2)
}
