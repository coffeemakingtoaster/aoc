package main

import (
	"bufio"
	"fmt"
	"os"
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

func buildInitial(l string) []int {
	res := []int{}
	for _, v := range []rune(l) {
		if v == 'S' {
			res = append(res, 1)
		} else {
			res = append(res, 0)
		}
	}
	return res
}

func applyLine(rays []int, line string) ([]int, int) {
	splitters := 0
	newRays := make([]int, len(rays))
	for i, v := range []rune(line) {
		if rays[i] == 0 {
			continue
		}
		switch v {
		case '^':
			splitters++
			newRays[i-1] = 1
			newRays[i] = max(newRays[i], 0)
			newRays[i+1] = 1
		default:
			newRays[i] = 1
		}
	}

	return newRays, splitters
}

func solve(input []string) []string {
	res := 0
	rays := buildInitial(input[0])
	for i := 1; i < len(input); i++ {
		var incr int
		rays, incr = applyLine(rays, input[i])
		res += incr

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
