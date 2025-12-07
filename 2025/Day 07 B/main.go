package main

import (
	"bufio"
	"fmt"
	"os"
)

func sum(s []int) int {
	res := 0
	for _, v := range s {
		res += v
	}
	return res
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

func applyLine(rays []int, line string) []int {
	newRays := make([]int, len(rays))
	for i, v := range []rune(line) {
		if rays[i] == 0 {
			continue
		}
		switch v {
		case '^':
			newRays[i-1] += rays[i]
			newRays[i] = max(newRays[i], 0)
			newRays[i+1] += rays[i]
		default:
			newRays[i] += rays[i]
		}
	}

	return newRays
}

func solve(input []string) []string {
	rays := buildInitial(input[0])
	for i := 1; i < len(input); i++ {
		rays = applyLine(rays, input[i])

	}
	return []string{fmt.Sprintf("%d", sum(rays))}
}

func main() {
	content, _ := readLines("./in.txt")
	solution := solve(content)
	for _, line := range solution {
		fmt.Println(line)
	}
}
