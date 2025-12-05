package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func parseRange(l string) []int {
	p := strings.Split(l, "-")
	a, _ := strconv.Atoi(p[0])
	b, _ := strconv.Atoi(p[1])
	return []int{a, b}
}

func getRanges(input []string) ([][]int, int) {
	y := 0
	res := [][]int{}
	l := input[y]

	for len(l) > 0 {
		res = append(res, parseRange(l))
		y++
		l = input[y]
	}

	return res, y
}

func checkNumber(ranges [][]int, l string) int {
	v, _ := strconv.Atoi(l)
	for _, r := range ranges {
		if v >= r[0] && v <= r[1] {
			return 1
		}
	}
	return 0
}

func solve(input []string) []string {
	res := 0
	ranges, y := getRanges(input)

	for y < len(input) {
		res += checkNumber(ranges, input[y])
		y++
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
