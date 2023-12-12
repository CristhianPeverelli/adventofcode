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
	filePath := "C:\\Users\\cpeverelli\\Desktop\\adventofcode\\2023\\day9\\input.txt"

	// PART 1
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sum, sum2 := 0, 0
	for scanner.Scan() {
		str := scanner.Text()
		vals := strings.Split(str, " ")
		values := []int{}
		for _, v := range vals {
			n, _ := strconv.Atoi(v)
			values = append(values, n)

		}
		sum += getValue(values)
		sum2 += getValue2(values)
	}
	fmt.Println("PART 1: ", sum)
	fmt.Println("PART 2: ", sum2)
	fmt.Printf("Execution time: %s\n", time.Since(startTime))
}

func getValue2(values []int) int {
	if len(values) == 0 {
		return 0
	}
	return values[0] - getValue2(convertedValues(values))
}

func getValue(values []int) int {
	if len(values) == 0 {
		return 0
	}
	return values[len(values)-1] + getValue(convertedValues(values))
}

func convertedValues(values []int) []int {
	res := []int{}
	for i := 0; i < len(values)-1; i++ {
		res = append(res, values[i+1]-values[i])
	}
	return res
}
