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
	matrix := ToMap(input)
	zeroes := FindStart(matrix)
	GetTrailheads(zeroes, matrix)
}

func GetTrailheads(zeroes [][]int, matrix [][]int) {
	totalPartOne := 0
	totalPartTwo := 0
	for _, zero := range zeroes {
		results := [][]int{zero}
		for i := 0; i < 9; i++ {
			temp_res := [][]int{}
			for _, result := range results {
				res := SearchNextLevel(matrix, result)
				for _, x := range res {
					temp_res = append(temp_res, x)
				}
			}
			results = temp_res
			if i == 8 {
				totalPartOne += len(removeDuplicates(results))
				totalPartTwo += len(results)
			}
		}
	}
	fmt.Println("part 1: ", totalPartOne)
	fmt.Println("part 2: ", totalPartTwo)
}

func ToMap(t string) [][]int {
	x := strings.Split(t, "\n")
	var out [][]int
	var outLine []int
	for _, line := range x {
		if len(line) == 0 {
			continue
		}
		for _, c := range line {
			cInt, _ := strconv.Atoi(string(c))
			outLine = append(outLine, cInt)
		}
		out = append(out, outLine)
		outLine = []int{}
	}
	return out
}

func SearchNextLevel(matrix [][]int, start []int) [][]int {
	// find +1 of the current level up to 9.k
	x := start[0]
	y := start[1]
	currLevel := matrix[y][x]
	nextLevel := currLevel + 1
	// fmt.Println("currlvl: ", currLevel)
	// fmt.Println("next level: ", nextLevel)
	a := []int{-1, 1}
	var result [][]int
	for _, j := range a {
		if x+j >= 0 && x+j < len(matrix[0]) && matrix[y][x+j] == nextLevel {
			result = append(result, []int{x + j, y})
		}
		if y+j >= 0 && y+j < len(matrix) && matrix[y+j][x] == nextLevel {
			result = append(result, []int{x, y + j})
		}
	}
	return result
}

func FindStart(matrix [][]int) [][]int {
	var zeroes [][]int
	for y, line := range matrix {
		for x, n := range line {
			if n == 0 {
				zeroes = append(zeroes, []int{x, y})
			}
		}
	}
	return zeroes
}

func removeDuplicates(slice [][]int) [][]int {
	// Use a map to track seen elements
	var out [][]int

	var seen bool
	for _, value := range slice {
		seen = false
		for _, x := range out {
			if value[0] == x[0] && value[1] == x[1] {
				seen = true
				break
			}
		}
		if seen == false {
			out = append(out, value)
		}
	}
	return out
}
