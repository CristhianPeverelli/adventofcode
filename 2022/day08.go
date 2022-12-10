package main

import (
	"fmt"
	"strconv"
	"strings"
	"io/ioutil"
	"time"
)

var input = readInput("input.txt")

func main() {
	start := time.Now()
	firstPartScore := len(input)*2 + len(input[0])*2 - 4 //This is the corner of all the matrix (itopScore all visible)
	secondPartScore := 0
	for i, s := range input {
		if i == 0 || i == len(input)-1 { continue }
		for j := range s {
			if j == 0 || j == len(s)-1 { continue }
			isVisible, score := getValues(i, j)
			if isVisible { firstPartScore++ }
			if score > secondPartScore { secondPartScore = score }
		}
	}
	fmt.Println("FIRST PART: ", firstPartScore)
	fmt.Println("SECOND PART: ", secondPartScore)
	fmt.Println("EXECUTION TIME: ",time.Since(start))
}

func readInput(f string) [][]int {
	input,_ := ioutil.ReadFile(f)
	s := strings.Split(string(input), "\n")
	var i [][]int
	for _, elem := range s {
		var t []int
		for _, elem := range strings.Split(elem, "") {
			num, _ := strconv.Atoi(elem)
			t = append(t, num)
		}
		i = append(i, t)
	}
	return i
}

func getValues(col, row int) (bool, int) {
	value := input[col][row]
	top, topScore := checkTop(col, row, value)
	bottom, bottomScore := checkBottom(col, row, value)
	left, leftScore := checkLeft(col, row, value)
	right, rightScore := checkRight(col, row, value)
	return top || bottom || left || right, topScore * bottomScore * leftScore * rightScore
}

func checkTop(col, row, value int) (bool, int) {
	score := 0
	for i := col - 1; i >= 0; i-- {
		score++
		if input[i][row] >= value {
			return false, score
		}
	}
	return true, score
}

func checkBottom(col, row, value int) (bool, int) {
	score := 0
	for i := col + 1; i <= len(input)-1; i++ {
		score++
		if input[i][row] >= value {
			return false, score
		}
	}
	return true, score
}

func checkLeft(col, row, value int) (bool, int) {
	score := 0
	for i := row - 1; i >= 0; i-- {
		score++
		if input[col][i] >= value {
			return false, score
		}
	}
	return true, score
}

func checkRight(col, row, value int) (bool, int) {
	score := 0
	for i := row + 1; i <= len(input[row])-1; i++ {
		score++
		if input[col][i] >= value {
			return false, score
		}
	}
	return true, score
}
