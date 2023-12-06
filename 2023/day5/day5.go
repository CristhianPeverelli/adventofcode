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
	filePath := "C:\\Users\\cpeverelli\\Desktop\\adventofcode\\2023\\day5\\input.txt"

	// PART 1
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sum := 0
	scanner.Scan()
	str := scanner.Text()
	values := strings.Split(str, " ")
	seeds := []int64{}
	for _, v := range values {
		s, _ := strconv.ParseInt(v, 10, 64)
		seeds = append(seeds, s)
	}
	step := 0
	for scanner.Scan() {
		str = scanner.Text()
		if len(str) > 0 {
			if isLetter(str[0]) {
				sMap := ""
				for len(str) > 0 {
					scanner.Scan()
					str = scanner.Text()
					sMap += str + " "
				}
				seeds = nextStep(seeds, sMap, step)
				step++
			}
		}
	}
	fmt.Println("PART 1: ", lowest(seeds))
	fmt.Println("PART 2: ", sum)
	fmt.Printf("Execution time: %s\n", time.Since(startTime))
}

func nextStep(seeds []int64, maps string, step int) []int64 {
	sMap := strings.Split(maps, " ")
	values := []int64{}
	for _, s := range sMap {
		if len(s) > 0 {
			n, _ := strconv.ParseInt(s, 10, 64)
			values = append(values, n)
		}
	}
	for j, seed := range seeds {
		for i := 0; i < len(values)-2; i += 3 {
			if seed >= values[i+1] && seed <= (values[i+1]+values[i+2]) {
				seeds[j] = values[i] + (seed - values[i+1])
				break
			}
		}

	}
	return seeds
}

func lowest(vals []int64) int64 {
	min := vals[0]
	for _, v := range vals {
		if v < min {
			min = v
		}
	}
	return min
}

func isLetter(c byte) bool {
	if c >= 'a' && c <= 'z' {
		return true
	}
	return false
}
