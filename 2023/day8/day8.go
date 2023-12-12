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

	// PART 2
	aNodes := []string{}
	for key := range nodes {
		if key[2] == 'A' {
			aNodes = append(aNodes, key)
		}
	}
	fmt.Println(aNodes)
	index := 0
	count2 := 0
	for !completed(aNodes) {
		if index >= len(steps) {
			index = 0
		}
		for i, node := range aNodes {
			str := nodes[strings.TrimSpace(node)]
			vals := strings.Split(str, ",")
			if steps[index] == 'L' {
				aNodes[i] = vals[0]
			} else {
				aNodes[i] = vals[1]
			}
		}
		count2++
		index++
	}
	fmt.Println("PART 1: check previous commit")
	fmt.Println("PART 2: ", count2)
	fmt.Printf("Execution time: %s\n", time.Since(startTime))
}

func completed(a []string) bool {
	for _, str := range a {
		if str[2] != 'Z' {
			return false
		}
	}
	return true
}
