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
	filePath := "C:\\Users\\Peve\\Desktop\\adventofcode\\2023\\day4\\input.txt"

	// PART 1
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	sum := 0
	sum2 := 0
	games := make([]int, 197)
	for i := range games {
		games[i] = 1
	}
	for scanner.Scan() {
		str := scanner.Text()
		vals := strings.Split(str, ":")
		game := vals[1]
		next := strings.Split(vals[0], " ")
		gameNumber, _ := strconv.Atoi(next[len(next)-1])
		winning, actual := strings.Split(game, "|")[0], strings.Split(game, "|")[1]
		n1 := convert(winning)
		n2 := convert(actual)
		points := getPoints(n1, n2)
		points2 := getPoints2(n1, n2)
		for i := 0; i < points2; i++ {
			if gameNumber+i < len(games) {
				games[gameNumber+i] += games[gameNumber-1]
			}
		}
		sum += points
		sum2 += games[gameNumber-1]
	}

	fmt.Println("PART 1: ", sum)
	fmt.Println("PART 2: ", sum2)
	fmt.Printf("Execution time: %s\n", time.Now().Sub(startTime))
}

func convert(a string) (res []int) {
	valori := strings.Split(a, " ")
	for _, n := range valori {
		num, _ := strconv.Atoi(n)
		if num != 0 {
			res = append(res, num)
		}
	}
	return
}

func getPoints(first []int, second []int) int {
	count := 0
	for _, num := range second {
		if present(num, first) {
			if count == 0 {
				count = 1
			} else {
				count *= 2
			}
		}
	}
	return count
}

func getPoints2(first []int, second []int) int {
	count := 0
	for _, num := range second {
		if present(num, first) {
			count++
		}
	}
	return count
}

func present(num int, a []int) bool {
	for _, n := range a {
		if n == num {
			return true
		}
	}
	return false
}
