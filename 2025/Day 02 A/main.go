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

func invalidSum(a, b string) int {
	if len(a) == len(b) && len(a)%2 == 1 {
		return 0
	}
	res := 0
	na, _ := strconv.Atoi(a)
	nb, _ := strconv.Atoi(b)

	// dumbest possible approach, because I am lazy
	for na <= nb {
		astr := strconv.Itoa(na)
		last := astr[len(astr)/2:]
		first := astr[:len(astr)/2]
		nlast, _ := strconv.Atoi(last)
		nfirst, _ := strconv.Atoi(first)

		if len(astr)%2 == 1 {
			na = int(math.Pow10(len(astr)))
			continue
		}

		if len(first) > len(strconv.Itoa(nlast)) {
			nlast = int(math.Pow10(len(first) - 1))
			na, _ = strconv.Atoi(first + strconv.Itoa(nlast))
			continue
		}

		if nlast == nfirst {
			fmt.Println(na)
			res += na
		}
		na++
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
