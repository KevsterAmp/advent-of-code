package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	var right []int

	for fileScanner.Scan() {
		// fmt.Printf(fileScanner.Text())
		result := strings.Split(fileScanner.Text(), "   ")
		// fmt.Println("result: ", result)

		leftInt, err := strconv.Atoi(result[0])
		check(err)

		rightInt, err := strconv.Atoi(result[1])
		check(err)

		// fmt.Println("left: ", leftInt)
		// fmt.Println("right: ", rightInt)

		left = append(left, leftInt)
		right = append(right, rightInt)
		// fileLines = append(fileLines, fileScanner.Text())
	}
	file.Close()

	// fmt.Println("left sorted: ", left[:5])
	// fmt.Println("right sorted: ", right[:5])

	sort.Ints(left)
	sort.Ints(right)

	// fmt.Println("left sorted: ", left[:5])
	// fmt.Println("right sorted: ", right[:5])
	//
	// fmt.Println("left len: ", len(left))
	// fmt.Println("right sorted: ", len(right))
	var totalDistance int
	for i := 0; i < 1000; i++ {
		if right[i]-left[i] < 0 {
			totalDistance += left[i] - right[i]
		} else {
			totalDistance += right[i] - left[i]
		}
	}
	fmt.Println("total distance: ", totalDistance)

}
