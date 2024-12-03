package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	safetyNum := 0
	for fileScanner.Scan() {
		lineStr := strings.Split(fileScanner.Text(), " ")
		line := ArrayAtoi(lineStr)
		fmt.Println(line)

		var inc bool
		var diff int
		safety := true
		for i := 0; i < len(line)-1; i++ {
			diff = line[i] - line[i+1]
			if diff <= 3 && diff >= 1 {
				if i == 0 {
					inc = true
				} else if inc != true {
					// fmt.Println(line[i], line[i+1])
					safety = false
					break
				}
			} else if diff >= -3 && diff <= -1 {
				if i == 0 {
					inc = false
				} else if inc != false {
					// fmt.Println(line[i], line[i+1])
					safety = false
					break
				}
			} else {
				// fmt.Println("ping")
				// fmt.Println(line[i], line[i+1])
				safety = false
				break
			}
		}
		// fmt.Println("Safety: ", safety)
		// fmt.Println("Inc: ", inc)
		// fmt.Println("diff: ", diff)
		if safety == true {
			safetyNum++
		}
	}
	fmt.Println("safetyNum: ", safetyNum)
}
