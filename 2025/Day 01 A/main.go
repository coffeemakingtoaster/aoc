package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

func turn(dial, turn int) int {
	dial = dial + turn
	return dial % 100
}

func solve(input []string) []string {
	res := 0
	dial := 50
	for _, l := range input {
		n, _ := strconv.Atoi(l[1:])
		if l[0] == 'L' {
			n *= -1
		}
		dial = turn(dial, n)
		if dial == 0 {
			res++
		}
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
