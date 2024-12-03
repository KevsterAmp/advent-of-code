package main

import (
	"bufio"
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
	file, err := os.Open("input.txt")
	check(err)
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	total := 0
	for fileScanner.Scan() {
		lineStr := fileScanner.Text()
		r, err := regexp.Compile(`mul\((\d+),(\d+)\)`)
		check(err)
		out := r.FindAllStringSubmatch(lineStr, -1)
		for _, match := range out {
			x, _ := strconv.Atoi(match[2])
			y, _ := strconv.Atoi(match[1])
			total += x * y
		}
	}
	fmt.Println("total: ", total)
	file.Close()
}
