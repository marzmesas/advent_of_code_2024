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

//part 2

func checkIncreasingPt2(row []int, modif_element int) bool {
	for i := 0; i < len(row)-1; i++ {
		j := i + 1
		if (absInt(row[i]-row[j]) > 3 && modif_element == 1) || (absInt(row[i]-row[j]) < 1 && modif_element == 1) {
			return false
		} else if (absInt(row[i]-row[j]) > 3 && modif_element == 0) || (absInt(row[i]-row[j]) < 1 && modif_element == 0) {
			modif_element += 1
		} else if row[j] < row[i] && modif_element == 1 {
			return false
		} else if row[j] < row[i] && modif_element == 0 {
			modif_element += 1
		}

	}
	return true
}

func checkDecreasingPt2(row []int, modif_element int) bool {

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

func checkSafetyPt2(row []int) bool {
	i := 0
	j := 1
	var result bool
	modif_element := 0
	if absInt(row[i]-row[j]) > 3 || absInt(row[i]-row[j]) < 1 {
		return false
	}
	if row[i] < row[j] {
		slice := row[j:]
		result = checkIncreasingPt2(slice, modif_element)
	} else if row[i] > row[j] {
		slice := row[j:]
		result = checkDecreasingPt2(slice, modif_element)
	} else {
		result = false
	}
	return result
}

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

	//part 2
	var safe_pt2 = 0
	for _, row := range matrix {
		val := checkSafetyPt2(row)
		if val {
			safeCounter += 1
		}
	}
	fmt.Printf("Total number of safe reports, part 2: %v\n", safe_pt2)

}
