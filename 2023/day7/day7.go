package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	startTime := time.Now()
	filePath := "C:\\Users\\cpeverelli\\Desktop\\adventofcode\\2023\\day7\\input.txt"

	// PART 1
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sum := 0
	rankedHands := []string{}
	for scanner.Scan() {
		str := scanner.Text()
		rankedHands = insertSort(rankedHands, fmt.Sprint(getRank(replaceJokers((str[:6]))))+" - "+(str[:6])+"- "+str[6:])
	}
	currentRank := 1
	lastRanked := 0
	for i, str := range rankedHands {
		c, _ := strconv.Atoi(string(str[0]))
		if c > currentRank {
			fmt.Println("Ordino:", rankedHands[lastRanked:i])
			orderCards(rankedHands[lastRanked:i])
			currentRank += c - currentRank
			lastRanked = i
		}
	}
	fmt.Println("Ordino:", rankedHands[lastRanked:])
	orderCards(rankedHands[lastRanked:])
	fmt.Println(rankedHands)
	for i, str := range rankedHands {
		values := strings.Split(str, "-")
		num, _ := strconv.Atoi(strings.TrimSpace(values[2]))
		fmt.Println(num, "*", i+1)
		sum += num * (i + 1)
	}
	fmt.Println("PART 1: check previous commit")
	fmt.Println("PART 2: ", sum)
	fmt.Printf("Execution time: %s\n", time.Since(startTime))

}

func insertSort(array []string, str string) []string {
	inserted := false
	for i, n := range array {
		if str[0] < n[0] {
			secondPart := append([]string{}, array[i:]...)
			array = append(array[:i], str)
			array = append(array, secondPart...)
			inserted = true
			break
		}
	}
	if !inserted {
		array = append(array, str)
	}
	return array
}

func orderCards(arr []string) {
	sort.Slice(arr, func(i, j int) bool {
		return compareStrings(arr[i], arr[j])
	})
}

func compareStrings(str1, str2 string) bool {
	order := "23456789TJQKA"
	for i := 0; i < len(str1) && i < len(str2); i++ {
		index1 := getIndex(str1[i], order)
		index2 := getIndex(str2[i], order)
		if index1 < index2 {
			return true
		} else if index1 > index2 {
			return false
		}
	}
	return len(str1) < len(str2)
}

func getIndex(char byte, order string) int {
	for i, c := range order {
		if byte(c) == char {
			return i
		}
	}
	return -1
}
func getRank(hand string) int {
	cardCounts := make(map[rune]int)
	for _, card := range hand {
		cardCounts[card]++
	}
	sortedCards := make([]rune, 0, len(cardCounts))
	for card := range cardCounts {
		sortedCards = append(sortedCards, card)
	}
	sort.Slice(sortedCards, func(i, j int) bool {
		if cardCounts[sortedCards[i]] == cardCounts[sortedCards[j]] {
			return sortedCards[i] > sortedCards[j]
		}
		return cardCounts[sortedCards[i]] > cardCounts[sortedCards[j]]
	})
	if cardCounts[sortedCards[0]] == 5 {
		return 7 // Five of a kind
	} else if cardCounts[sortedCards[0]] == 4 {
		return 6 // Four of a kind
	} else if cardCounts[sortedCards[0]] == 3 && cardCounts[sortedCards[1]] == 2 {
		return 5 // Full house
	} else if cardCounts[sortedCards[0]] == 3 {
		return 4 // Three of a kind
	} else if cardCounts[sortedCards[0]] == 2 && cardCounts[sortedCards[1]] == 2 {
		return 3 // Two pair
	} else if cardCounts[sortedCards[0]] == 2 {
		return 2 // One pair
	} else {
		return 1 // High card
	}
}

func replaceJokers(hand string) string {
	jokerCount := strings.Count(hand, "J")
	return replaceJokersRecursive(hand, jokerCount)
}

func replaceJokersRecursive(hand string, jokerCount int) string {
	if jokerCount == 0 {
		return hand
	}
	ranks := "23456789TJQKA"
	maxHand := hand
	maxRank := getRank(hand)
	for _, r := range ranks {
		newHand := strings.Replace(hand, "J", string(r), 1)
		newRank := getRank(newHand)

		if newRank > maxRank {
			maxRank = newRank
			maxHand = newHand
		}
	}
	return replaceJokersRecursive(maxHand, jokerCount-1)
}
