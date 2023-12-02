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
	filePath := "C:\\Users\\cpeverelli\\Desktop\\adventofcode\\2023\\day2\\input.txt"

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
	fmt.Printf("Execution time: %s\n", time.Since(startTime))
}
