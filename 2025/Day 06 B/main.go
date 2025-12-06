package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Block struct {
	values []int
	op     func(int, int) int
}

func (b Block) solve() int {
	res := b.values[0]
	for i := 1; i < len(b.values); i++ {
		res = b.op(res, b.values[i])
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

func getMaxLengthOfLine(l string) int {
	i := 0
	curr := 0
	for i < len(l) {
		if l[i] == ' ' {
			return curr
		}
		curr++
		i++

	}
	return curr
}

func getBlock(field []string, x int) (Block, int) {
	space_count := 0
	b := Block{}
	values := []string{}
	for x < len(field[0]) {
		space_count = 0
		buffer := ""
		for i := range len(field) {
			v := field[i][x]
			switch v {
			case '*':
				b.op = func(a, b int) int {
					return a * b
				}
				continue
			case '+':
				b.op = func(a, b int) int {
					return a + b
				}
				continue
			case ' ':
				space_count++
				fallthrough
			default:
				buffer += string(v)
			}
		}
		x++
		if space_count == len(field) {
			break
		}
		values = append(values, buffer)
		buffer = ""
	}
	for _, v := range values {
		n, err := strconv.Atoi(strings.TrimSpace(v))
		if err != nil {
			panic(err)
		}
		b.values = append(b.values, n)
	}
	return b, x
}

func horField(field []string) []Block {
	res := []Block{}
	x := 0
	for x < len(field[0]) {
		b, i := getBlock(field, x)
		res = append(res, b)
		x = i
	}
	return res
}

func solve(input []string) []string {
	res := 0
	blocks := horField(input)
	for _, b := range blocks {
		res += b.solve()
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
