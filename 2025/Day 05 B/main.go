package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func sortAndParseRanges(ranges [][]int) [][]int {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})
	returnRanges := [][]int{ranges[0]}
	for i, s := range ranges {
		if i == 0 {
			continue
		}
		comp := returnRanges[len(returnRanges)-1]
		if comp[1] >= s[0] {
			returnRanges[len(returnRanges)-1] = []int{comp[0], max(s[1], comp[1])}
		} else {
			returnRanges = append(returnRanges, s)
		}
	}
	fmt.Printf("Received %d seeds, merged into %d seeds!\n", len(ranges), len(returnRanges))
	return returnRanges
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
		if y == len(input) {
			break
		}
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
	ranges, _ := getRanges(input)
	ranges = sortAndParseRanges(ranges)

	for _, r := range ranges {
		res += r[1] - r[0] + 1
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
