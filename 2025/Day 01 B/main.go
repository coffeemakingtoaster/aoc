package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
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

func turn(dial, turn int) (int, int) {
	if turn == 0 {
		return dial, 0
	}
	startDial := dial
	dial = (dial + turn)
	if dial < 0 {
		dial = 100 + dial
	}
	if dial >= 100 {
		dial = dial - 100
	}
	/*
		This does nothing here...why?
			dial = dial % 100
	*/
	if dial == 0 || startDial == 0 {
		return dial, 0
	}

	if turn < 0 && dial >= startDial {
		fmt.Printf("%d->%d (%d) [1]\n", startDial, dial, turn)
		return dial, 1
	}
	if turn > 0 && dial <= startDial {
		fmt.Printf("%d->%d (%d) [1]\n", startDial, dial, turn)
		return dial, 1
	}
	fmt.Printf("%d->%d (%d) [0]\n", startDial, dial, turn)
	return dial, 0
}

func solve(input []string) []string {
	res := 0
	dial := 50
	for _, l := range input {
		var incr int
		n, _ := strconv.Atoi(l[1:])
		if n >= 100 {
			fmt.Printf("Larger than 100: %d (%d)\n", int(math.Floor(float64(n/100))), n)
			res += int(math.Floor(float64(n / 100)))
			n = n % 100
		}
		if l[0] == 'L' {
			n *= -1
		}
		dial, incr = turn(dial, n)
		res += incr
		if dial == 0 {
			res++
		}
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
