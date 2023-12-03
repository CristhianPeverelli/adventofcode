package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {
	startTime := time.Now()
	filePath := "C:\\Users\\Peve\\Desktop\\adventofcode\\2023\\day3\\input.txt"

	// PART 1
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	sum := 0
	for i, l := range lines {
		re := regexp.MustCompile(`\d+`)
		matches := re.FindAllStringIndex(l, -1)
		for _, m := range matches {
			for j := m[0]; j < m[len(m)-1]; j++ {
				if isAdjacent(i, j, lines) {
					n := l[m[0]:m[len(m)-1]]
					num, _ := strconv.Atoi(n)
					sum += num
					break
				}
			}
		}
	}

	// PART 2
	fmt.Println("PART 2: ", sum)
	fmt.Printf("Execution time: %s\n", time.Now().Sub(startTime))
}

func isAdjacent(i, j int, lines []string) bool {
	for x := i - 1; x <= i+1; x++ {
		for y := j - 1; y <= j+1; y++ {
			if x < 0 || x >= len(lines) || y < 0 || y >= len(lines[i]) {
				continue
			}
			if lines[x][y] != '.' && (lines[x][y] < '0' || lines[x][y] > '9') {
				return true
			}
		}
	}
	return false
}
