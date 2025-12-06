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

func convertField(field []string) map[int][]string {
	res := make(map[int][]string)
	for _, l := range field {
		parts := strings.Split(l, " ")
		ind := 0
		for _, p := range parts {
			if len(p) == 0 {
				continue
			}
			if v, ok := res[ind]; ok {
				res[ind] = append(v, p)
			} else {
				res[ind] = []string{p}
			}
			ind++
		}
	}
	return res
}

func solveCalc(in []string) int {
	res, _ := strconv.Atoi(in[0])
	op := func(a, b int) int {
		return a + b
	}
	if in[len(in)-1] == "*" {
		op = func(a, b int) int {
			return a * b
		}
	}
	for i := range len(in) - 1 {
		if i == 0 {
			continue
		}
		c, _ := strconv.Atoi(in[i])
		res = op(res, c)
	}
	return res
}

func solve(input []string) []string {
	res := 0
	m := convertField(input)
	for _, v := range m {
		res += solveCalc(v)
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
