package main

import (
	"fmt"
	"os"
	// "regexp"
	"strconv"
	"strings"
)

func main() {
	content, _ := os.ReadFile("input.txt")
	input := string(content)
	input = input[:len(input)-1]
	stones := strings.Split(input, " ")
	stoneRecord := make(map[string]int)

	// initial stoneRecord
	for _, stone := range stones {
		if In(stone, stoneRecord) {
			stoneRecord[stone] = 1
		} else {
			stoneRecord[stone]++
		}
	}
	fmt.Println("stoneRecord", stoneRecord)

	for i := 0; i < 75; i++ {
		newStoneRecord := make(map[string]int)

		for key := range stoneRecord {
			stone := key
			count := stoneRecord[key]
			outStones := blink(stone)
			for _, x := range outStones {
				temp, _ := strconv.Atoi(x)
				x = strconv.Itoa(temp)
				if In(x, newStoneRecord) == false {
					newStoneRecord[x] = 1 * count
				} else {
					newStoneRecord[x] += 1 * count
				}
			}
		}
		stoneRecord = newStoneRecord
		if i == 24 || i == 74 {
			fmt.Println("result: ", SumRecords(stoneRecord))
		}
	}
}

func blink(stone string) []string {
	var newStone []string
	if len(stone)%2 == 0 {
		mid := len(stone) / 2
		newStone = append(newStone, stone[:mid])
		newStone = append(newStone, stone[mid:])
	} else if stone == "0" {
		newStone = append(newStone, "1")
	} else {
		numStone, _ := strconv.Atoi(stone)
		stone = strconv.Itoa(numStone * 2024)
		newStone = append(newStone, stone)
	}
	return newStone
}

func In[S string, V any](key S, record map[S]V) bool {
	_, exists := record[key]
	return exists
}

func SumRecords(record map[string]int) int {
	sum := 0
	for key := range record {
		sum += record[key]
	}
	return sum
}
