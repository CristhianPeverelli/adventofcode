package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Node struct {
	symbol  rune
	visited bool
}

var maxSteps int

func main() {
	startTime := time.Now()
	filePath := "C:\\Users\\cpeverelli\\Desktop\\adventofcode\\2023\\day10\\input.txt"

	// PART 1
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	str := scanner.Text()
	matrix := make([][]Node, len(str))
	for i, _ := range matrix {
		matrix[i] = make([]Node, len(matrix))
	}
	row := 0
	for i, v := range str {
		matrix[row][i] = Node{symbol: v, visited: false}
	}
	iStart := 0
	jStart := 0
	for scanner.Scan() {
		str = scanner.Text()
		row++
		for i, v := range str {
			if v == '.' {
				matrix[row][i] = Node{symbol: v, visited: false}
			} else if v == 'S' {
				iStart = row
				jStart = i
				matrix[row][i] = Node{symbol: v, visited: true}
			} else {
				matrix[row][i] = Node{symbol: v, visited: false}

			}
		}
	}

	maxSteps := []int{}
	if matrix[iStart-1][jStart].symbol != '.' {
		maxSteps = append(maxSteps, 1+checkAdiacenti(iStart-1, jStart, matrix))
	}
	if matrix[iStart+1][jStart].symbol != '.' {
		maxSteps = append(maxSteps, 1+checkAdiacenti(iStart+1, jStart, matrix))
	}
	if matrix[iStart][jStart+1].symbol != '.' {
		maxSteps = append(maxSteps, 1+checkAdiacenti(iStart, jStart+1, matrix))
	}
	if matrix[iStart][jStart-1].symbol != '.' {
		maxSteps = append(maxSteps, 1+checkAdiacenti(iStart, jStart-1, matrix))
	}
	max := 0
	for _, c := range maxSteps {
		if c > max {
			max = c
		}
	}
	fmt.Println("PART 1: ", max)
	fmt.Printf("Execution time: %s\n", time.Since(startTime))
}

func checkAdiacenti(i int, j int, matrix [][]Node) int {
	if i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix[0]) || matrix[i][j].visited {
		return 0
	}
	matrix[i][j].visited = true
	switch matrix[i][j].symbol {
	case 'F':
		if matrix[i+1][j].visited {
			return 1 + checkAdiacenti(i, j+1, matrix)
		} else {
			return 1 + checkAdiacenti(i+1, j, matrix)
		}
	case 'J':
		if matrix[i-1][j].visited {
			return 1 + checkAdiacenti(i, j-1, matrix)
		} else {
			return 1 + checkAdiacenti(i-1, j, matrix)
		}
	case 'L':
		if matrix[i-1][j].visited {
			return 1 + checkAdiacenti(i, j+1, matrix)
		} else {
			return 1 + checkAdiacenti(i-1, j, matrix)
		}
	case '7':
		if matrix[i+1][j].visited {
			return 1 + checkAdiacenti(i, j-1, matrix)
		} else {
			return 1 + checkAdiacenti(i+1, j, matrix)
		}
	case '-':
		if matrix[i][j-1].visited {
			return 1 + checkAdiacenti(i, j+1, matrix)
		} else {
			return 1 + checkAdiacenti(i, j-1, matrix)
		}
	case '|':
		if matrix[i-1][j].visited {
			return 1 + checkAdiacenti(i+1, j, matrix)
		} else {
			return 1 + checkAdiacenti(i-1, j, matrix)
		}
	default:
		return 0
	}
}

// func printMatrix(m [][]Node) {
// 	for i, _ := range m {
// 		for j, _ := range m {
// 			fmt.Print(string(m[i][j].symbol), " ")
// 		}
// 		fmt.Println()
// 	}
// }
