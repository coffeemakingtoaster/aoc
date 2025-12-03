package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func evalRow(row string) int {
	b1, _ := strconv.Atoi(string(rune(row[len(row)-1])))
	b2, _ := strconv.Atoi(string(rune(row[len(row)-2])))
	buffer := []int{b2, b1}
	fmt.Printf("Start value: %v\n", buffer)
	i := len(row) - 2 - 1
	for i >= 0 {
		n, _ := strconv.Atoi(string(rune(row[i])))
		fmt.Printf("Val: %d\n", n)
		if n >= buffer[0] {
			buffer[1] = max(buffer[0], buffer[1])
			buffer[0] = n
		}
		i--
	}
	return buffer[0]*10 + buffer[1]
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func solve(input []string) []string {
	res := 0
	for _, l := range input {
		res += evalRow(l)
	}
	return []string{fmt.Sprintf("%d", res)}
}

func main() {
	content, _ := readLines("./in.txt")
	solution := solve(content)
	for _, line := range solution {
		fmt.Println(line)
	}
}
