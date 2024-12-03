package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ArrayAtoi(t []string) []int {
	var output []int
	for _, value := range t {
		y, err := strconv.Atoi(value)
		check(err)
		output = append(output, y)
	}
	return output
}

func main() {
	content, err := os.ReadFile("input.txt")
	check(err)

	total := 0
	lineStr := string(content)
	r, err := regexp.Compile(`don't\(\)|do\(\)|mul\((\d+),(\d+)\)`)
	check(err)
	out := r.FindAllStringSubmatch(lineStr, -1)
	do := true
	for _, match := range out {
		if match[0] == "do()" {
			do = true
		} else if match[0] == "don't()" {
			do = false
		} else if do == true {
			x, _ := strconv.Atoi(match[2])
			y, _ := strconv.Atoi(match[1])
			total += x * y
		}
	}
	fmt.Println("total: ", total)
}
