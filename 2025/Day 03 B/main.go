package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const SIZE = 12

type Buffer struct {
	values []int
}

func makeBuffer(s string) *Buffer {
	b := Buffer{}
	for i := range SIZE {
		n, _ := strconv.Atoi(string(rune(s[len(s)-1-i])))
		b.values = append([]int{n}, b.values...)
	}
	return &b
}

func (b *Buffer) pushback(val int) {
	for i := range b.values {
		if val >= b.values[i] {
			old := b.values[i]
			b.values[i] = val
			val = old
		} else {
			break
		}
	}
}

func (b *Buffer) getVal() int {
	fmt.Printf("End value: %v (%d)\n", b.values, len(b.values))
	i := len(b.values) - 1
	res := 0
	for i >= 0 {
		res += b.values[i] * int(math.Pow(10, float64(len(b.values)-1-i)))
		i--
	}
	return res
}

func evalRow(row string) int {
	buffer := makeBuffer(row)
	fmt.Printf("Start value: %v\n", buffer.values)
	i := len(row) - SIZE - 1
	for i >= 0 {
		n, _ := strconv.Atoi(string(rune(row[i])))
		buffer.pushback(n)
		i--
	}
	return buffer.getVal()
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
