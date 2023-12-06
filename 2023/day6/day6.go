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
	filePath := "C:\\Users\\cpeverelli\\Desktop\\adventofcode\\2023\\day6\\input.txt"

	// PART 1
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	str := scanner.Text()
	values := strings.Split(strings.Split(str, ":")[1], " ")
	times := []string{}
	for _, c := range values {
		if strings.TrimSpace(c) != "" {
			times = append(times, c)
		}
	}
	scanner.Scan()
	str = scanner.Text()
	values = strings.Split(strings.Split(str, ":")[1], " ")
	distances := []string{}
	for _, c := range values {
		if len(strings.TrimSpace(c)) > 0 {
			distances = append(distances, c)
		}
	}

	sum := 1
	for i := 0; i < len(times); i++ {
		time, _ := strconv.Atoi(times[i])
		distance, _ := strconv.Atoi(distances[i])
		res := calcWays(time, distance)
		sum *= res
		fmt.Println(res)
	}

	fmt.Println("PART 1: ", sum)

	// PART 2
	fmt.Println("PART 2: ", sum)
	fmt.Printf("Execution time: %s\n", time.Now().Sub(startTime))
}

func calcWays(time int, distance int) (count int) {
	for milliseconds := 1; milliseconds < time; milliseconds++ {
		distanceTravelled := milliseconds * (time - milliseconds)
		if distance < distanceTravelled {
			count++
		}
	}
	return
}
