package main

import (
	"bufio"
	"fmt"
	"os"
)

func saveGetRoll(field []string, x, y int) int {
	if x < 0 || x >= len(field[0]) {
		return 0
	}
	if y < 0 || y >= len(field) {
		return 0
	}
	if field[y][x] == '@' {
		return 1
	}
	return 0
}

func evalPosition(field []string, x, y int) int {
	cnt := 0
	for _, scan_y := range []int{y - 1, y, y + 1} {
		for _, scan_x := range []int{x - 1, x, x + 1} {
			if scan_x == x && scan_y == y {
				continue
			}
			cnt += saveGetRoll(field, scan_x, scan_y)
		}
	}
	if cnt < 4 {
		return 1
	}
	return 0
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

func updateField(field []string, x, y int) []string {
	a := []rune(field[y])
	a[x] = '.'
	field[y] = string(a)
	return field
}

func solveIter(input []string) (int, []string) {
	res := 0
	for y := range len(input) {
		for x := range len(input[0]) {
			if saveGetRoll(input, x, y) == 0 {
				continue
			}
			r := evalPosition(input, x, y)
			if r == 1 {
				res++
				input = updateField(input, x, y)
			}
		}
	}
	fmt.Println("---")
	for _, l := range input {
		fmt.Println(l)
	}
	fmt.Println("---")
	return res, input
}

func solve(input []string) []string {
	res := 0
	r := 1
	for r > 0 {
		r, input = solveIter(input)
		res += r
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
