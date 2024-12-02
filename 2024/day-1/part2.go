package main

import (
	"bufio"
	"fmt"
	"maps"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("input.txt")
	check(err)
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var left []int
	rightCounter := make(map[int]int)

	for fileScanner.Scan() {
		// fmt.Printf(fileScanner.Text())
		result := strings.Split(fileScanner.Text(), "   ")
		// fmt.Println("result: ", result)

		leftInt, err := strconv.Atoi(result[0])
		check(err)

		rightInt, err := strconv.Atoi(result[1])
		check(err)

		left = append(left, leftInt)
		if len(rightCounter) == 0 {
			rightCounter[rightInt] = 1
		} else {
			rightKeys := maps.Keys(rightCounter)
			flag := false
			for k := range rightKeys {
				if rightInt == k {
					rightCounter[k]++
					flag = true
				}
			}
			if flag == false {
				rightCounter[rightInt] = 1
			}
		}
		// fileLines = append(fileLines, fileScanner.Text())
	}
	file.Close()

	totalDistance := 0
	for i := 0; i < 1000; i++ {
		rightKeys := maps.Keys(rightCounter)
		for k := range rightKeys {
			if k == left[i] {
				totalDistance += rightCounter[k] * left[i]
			}
		}

	}
	fmt.Println("total distance: ", totalDistance)

}
