package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"time"
	"unicode"
)

func main() {
	startTime := time.Now()
	filePath := "C:\\Users\\Peve\\Desktop\\adventofcode\\2023\\day1\\input.txt"

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
		var firstDigit, lastDigit int
		for _, char := range line {
			if unicode.IsDigit(char) {
				if firstDigit == 0 {
					firstDigit = int(char - '0')
				}
				lastDigit = int(char - '0')
			}
		}
		calibrationValue := firstDigit*10 + lastDigit
		sum += calibrationValue
	}
	fmt.Println("PART 1:", sum)

	// PART 2
	file, err = os.Open(filePath)
	defer file.Close()
	scanner = bufio.NewScanner(file)

	sum = 0
	mapp := make(map[string]int)
	mapp["one"] = 1
	mapp["two"] = 2
	mapp["three"] = 3
	mapp["four"] = 4
	mapp["five"] = 5
	mapp["six"] = 6
	mapp["seven"] = 7
	mapp["eight"] = 8
	mapp["nine"] = 9

	for scanner.Scan() {
		str := scanner.Text()
		keyword := []rune{}
		match := []int{}
		regularExp := regexp.MustCompile(`^one|^two|^three|^four|^five|^six|^seven|^eight|^nine`)
		for k, c := range str {
			if c-48 < 10 {
				keyword = append(keyword, c-48)
				continue
			}
			match = regularExp.FindStringIndex(str[k:])
			if match != nil {
				keyword = append(keyword, rune(mapp[str[k+match[0]:k+match[1]]]))
			}
		}
		if len(keyword) == 1 {
			sum += int(keyword[0]*10 + keyword[0])
		} else {
			sum += int(keyword[0]*10 + keyword[len(keyword)-1])
		}
	}
	fmt.Println("PART 2: ", sum)
	fmt.Printf("Execution time: %s\n", time.Now().Sub(startTime))
}
