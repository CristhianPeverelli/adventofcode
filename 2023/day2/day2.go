package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	startTime := time.Now()
	filePath := "C:\\Users\\Peve\\Desktop\\adventofcode\\2023\\day2\\input.txt"

	// PART 1
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, ":")
		gameNumber, _ := strconv.Atoi(values[0][5:])
		counters := make(map[string]int)
		sets := strings.Split(values[1], ";")
		correct := true
		for _, set := range sets {
			parts := strings.Split(set, ",")
			for _, part := range parts {
				vals := strings.Split(part, " ")
				num, _ := strconv.Atoi(vals[1])
				counters[string(vals[2])] += num
			}
			if counters["red"] > 12 || counters["green"] > 13 || counters["blue"] > 14 {
				correct = false
			}
			counters = make(map[string]int)
		}
		if correct {
			sum += gameNumber
		}
	}
	fmt.Println("PART 1:", sum)

	// PART 2
	file, _ = os.Open(filePath)
	defer file.Close()
	scanner = bufio.NewScanner(file)
	sum = 0
	for scanner.Scan() {
		line := scanner.Text()
		counters := make(map[string]int)
		sets := strings.Split(strings.Split(line, ":")[1], ";")
		for _, set := range sets {
			parts := strings.Split(set, ",")
			for _, part := range parts {
				vals := strings.Split(part, " ")
				num, _ := strconv.Atoi(vals[1])
				if num > counters[string(vals[2])] {
					counters[string(vals[2])] = num
				}
			}
		}
		sum += counters["green"] * counters["red"] * counters["blue"]
	}
	fmt.Println("PART 2:", sum)
	fmt.Printf("Execution time: %s\n", time.Since(startTime))
}
