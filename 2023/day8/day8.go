package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Node struct {
	Value    string
	Children []*Node
}

func main() {
	startTime := time.Now()
	filePath := "C:\\Users\\cpeverelli\\Desktop\\adventofcode\\2023\\day8\\input.txt"

	// PART 1
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	steps := scanner.Text()
	nodes := make(map[string]string)
	for scanner.Scan() {
		str := scanner.Text()
		vals := strings.Split(str, "=")
		nodes[strings.TrimSpace(vals[0])] = strings.TrimSpace(vals[1][2 : len(vals[1])-1])
	}
	selectedNode := "AAA"
	index := 0
	count := 0
	for selectedNode != "ZZZ" {
		if index >= len(steps) {
			index = 0
		}
		str := nodes[selectedNode]
		vals := strings.Split(str, ",")
		if steps[index] == 'L' {
			selectedNode = strings.TrimSpace(vals[0])
		} else {
			selectedNode = strings.TrimSpace(vals[1])
		}
		count++
		index++
	}
	fmt.Println("PART 1: ", count)
	fmt.Printf("Execution time: %s\n", time.Since(startTime))
}
