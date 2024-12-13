package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Parse an integer from a string
func parseInt(s string) int {
	var num int
	_, err := fmt.Sscanf(s, "%d", &num)
	if err != nil {
		return 0
	}
	return num
}

// Parse a comma-separated list of integers
func parseList(s string) []int {
	parts := strings.Split(s, ",")
	var result []int
	for _, part := range parts {
		result = append(result, parseInt(part))
	}
	return result
}

func main() {
	// Open the input file (you can replace this with reading from stdin or any other method)
	file, err := os.Open("day_5/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	scanner := bufio.NewScanner(file)

	// Parse rules
	rules := map[int][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break // End of rules section
		}
		parts := strings.Split(line, "|")
		if len(parts) != 2 {
			fmt.Println("Invalid rule:", line)
			continue
		}
		from := parseInt(parts[0])
		to := parseInt(parts[1])
		rules[from] = append(rules[from], to)
	}

	// Parse updates
	var updates [][]int
	for scanner.Scan() {
		line := scanner.Text()
		pages := parseList(line)
		updates = append(updates, pages)
	}
}
