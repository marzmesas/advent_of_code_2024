package main

import (
	"fmt"
	"log"
	"os"
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

func main() {
	// Specify the input file
	filename := "day_3/input.txt"

	// Call the function to read the file
	content, err := readFileToString(filename)
	if err != nil {
		log.Fatalf("Failed to read the file: %v", err)
	}

	// Print the content
	fmt.Println("File content as a string:")
	fmt.Println(content)
}
