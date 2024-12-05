package main

import (
	"fmt"
	"os"
	"strconv"
	// "regexp"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	content, err := os.ReadFile("input.txt")
	check(err)
	lineStr := string(content)

	data := strings.Split(lineStr, "\n\n")
	guide := strings.Split(data[0], "\n")
	numbers := strings.Split(data[1], "\n")
	numbers = numbers[:len(numbers)-1]
	// fmt.Println("guide: ", guide)
	// fmt.Println("numbers: ", numbers)

	total := 0
	for _, line := range numbers {
		numSlice := strings.Split(line, ",")
		pairs := CreatePairs(numSlice)
		// fmt.Println("x: ", x)
		filteredGuide := FilterGuide(guide, numSlice)
		// fmt.Println("filteredguide: ", filteredGuide)
		if Verify(filteredGuide, pairs) {
			mid := GetMiddle(numSlice)
			total += mid
		}
	}
	fmt.Println("total: ", total)
}

func In(s string, t []string) bool {
	for _, i := range t {
		if s == i {
			return true
		}
	}
	return false
}

// func filterguide
// remove guides that don't contain 2 of the numbers
func FilterGuide(guide []string, numSlice []string) []string {
	var out []string
	for _, x := range guide {
		num := strings.Split(x, "|")
		if In(num[0], numSlice) && In(num[1], numSlice) {
			out = append(out, x)
		}
	}
	return out
}

// func verify
// for x in filtered guides
// if there are filetered guides that does not exist on created pairs
// number is not properly sorted
func Verify(g []string, n []string) bool {
	for _, x := range g {
		if In(x, n) == false {
			return false
		}
	}
	return true
}

// func createpairs
// function to create pairs of numbers with each other
func CreatePairs(numSlice []string) []string {
	var out []string

	for i, x := range numSlice {
		for j, y := range numSlice {
			if x == y || i > j {
				continue
			}
			// fmt.Println(x + "|" + y)
			out = append(out, string(x)+"|"+string(y))
		}
	}
	return out
}

func GetMiddle(s []string) int {
	res, _ := strconv.Atoi(s[len(s)/2])
	return res

}
