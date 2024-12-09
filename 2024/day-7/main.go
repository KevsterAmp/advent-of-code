package main

import (
	"fmt"
	"os"
	"strconv"
	// "regexp"
	"strings"
)

func main() {
	resultOne := 0
	resultTwo := 0
	content, _ := os.ReadFile("input.txt")
	data := strings.Split(string(content), "\n")
	data = data[:len(data)-1]
	for i, line := range data {
		l := strings.Split(line, ": ")
		expected, _ := strconv.Atoi(l[0])
		nums := strings.Split(l[1], " ")
		// fmt.Println("expected: ", expected)
		// fmt.Println("nums: ", nums)
		if PartOne(nums, expected) {
			resultOne += expected
		}
		if PartTwo(nums, expected) {
			resultTwo += expected
		}
		if i == 1 {
		}
	}
	fmt.Println("Part one result: ", resultOne)
	fmt.Println("Part two result: ", resultTwo)
}

func PartTwo(nums []string, expected int) bool {
	symbols := []string{"+", "*", "|"}
	calibration := generateCombinations(symbols, len(nums)-1, nums, expected)
	if calibration {
		return true
	}
	return false
}

func PartOne(nums []string, expected int) bool {
	symbols := []string{"+", "*"}
	calibration := generateCombinations(symbols, len(nums)-1, nums, expected)
	if calibration {
		return true
	}
	return false
}

func DetermineCalibration(nums []string, operators []string, expected int) bool {
	result := generateCombinations(operators, len(nums)-1, nums, expected)
	if result == true {
		return true
	}
	return false
}

func ApplyOperator(x int, y int, operator string) int {
	var out int
	switch operator {
	case "+":
		// fmt.Println(x, "+", y)
		return x + y
	case "*":
		// fmt.Println(x, "*", y)
		return x * y
	case "|":
		xStr := strconv.Itoa(x)
		yStr := strconv.Itoa(y)
		out, _ = strconv.Atoi(xStr + yStr)
		return out
	}
	return -1
}

func generateCombinations(symbols []string, length int, nums []string, expected int) bool {
	totalCombinations := 1
	for i := 0; i < length; i++ {
		totalCombinations *= len(symbols)
	}

	combinations := make([]string, 0, totalCombinations)

	for i := 0; i < totalCombinations; i++ {
		combination := make([]string, length)

		current := i
		for j := length - 1; j >= 0; j-- {
			combination[j] = symbols[current%len(symbols)]
			current /= len(symbols)

		}
		// fmt.Println("Combo: ", combination)
		// fmt.Println("Nums: ", nums)
		total, _ := strconv.Atoi(nums[0])
		for i := 1; i < len(nums); i++ {
			n, _ := strconv.Atoi(nums[i])
			total = ApplyOperator(total, n, combination[i-1])
		}
		if total == expected {
			// fmt.Println(combo, nums)
			return true
		}

		combinations = append(combinations, strings.Join(combination, ""))
	}

	return false
}
