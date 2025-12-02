package main

import (
	"bufio"
	"fmt"
	"math"
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

func incr(s string, v int) string {
	n, _ := strconv.Atoi(s)
	n += v
	return strconv.Itoa(n)
}

func isPerm(s string, n int) int {
	if len(s)%n != 0 {
		return 1
	}
	if n == 1 {
		if strings.Count(s, string(rune(s[0]))) == len(s) {
			return 0
		}
		return 1
	}
	a, _ := strconv.Atoi(s)
	b, _ := strconv.Atoi(s[:n])
	if int(math.Log10(float64(a)))%1 == 0 && strings.Count(s, "0") == len(s)-1 {
		return 1
	}
	if a%b == 0 {
		//fmt.Printf("%d : %d\t", a, b)
		if strings.Count(s, strconv.Itoa(b)) == len(s)/n {
			return 0
		}
		return 1
	}
	return 1
}

func invalidSum(a, b string) int {
	res := 0
	na, _ := strconv.Atoi(a)
	nb, _ := strconv.Atoi(b)

	// dumbest possible approach, because I am lazy
	for na <= nb {
		astr := strconv.Itoa(na)
		m := 1
		for i := range len(astr)/2 + 1 {
			if i == 0 {
				continue
			}
			p := isPerm(astr, i)
			if p == 0 {
				fmt.Printf("Found %d for %d\n", na, i)
				res += na
				break
			}
			if i == 1 {
				m = max(p, m)
			}
		}
		na += m
	}
	return res
}

func solve(input []string) []string {
	res := 0
	ranges := strings.Split(input[0], ",")
	for _, r := range ranges {
		//fmt.Printf("%d/%d\n", i, len(ranges))
		numbers := strings.Split(r, "-")
		res += invalidSum(numbers[0], numbers[1])
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
