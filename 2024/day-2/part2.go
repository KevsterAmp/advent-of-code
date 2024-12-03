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

		safety, _ := DetermineSafety(line)
		if safety == false {
			for i := 0; i < len(line); i++ {
				reducedLine := RemoveElement(line, i)
				safety, _ := DetermineSafety(reducedLine)
				if safety == true {
					safetyNum++
					break
				}
			}
		}
		if safety == true {
			safetyNum++
		}
	}
	fmt.Println("safetyNum: ", safetyNum)
}

func DetermineSafety(line []int) (bool, []int) {
	var inc bool
	var diff int
	for i := 0; i < len(line)-1; i++ {
		pairIndex := []int{i, i + 1}
		diff = line[i] - line[i+1]
		if diff <= 3 && diff >= 1 {
			if i == 0 {
				inc = true
			} else if inc != true {
				// fmt.Println(line[i], line[i+1])
				return false, pairIndex
			}
		} else if diff >= -3 && diff <= -1 {
			if i == 0 {
				inc = false
			} else if inc != false {
				// fmt.Println(line[i], line[i+1])
				return false, pairIndex
			}
		} else {
			// fmt.Println("ping")
			// fmt.Println(line[i], line[i+1])
			return false, pairIndex
		}
	}
	// fmt.Println("Safety: ", safety)
	// fmt.Println("Inc: ", inc)
	// fmt.Println("diff: ", diff)
	return true, nil
}

func RemoveElement(t []int, i int) []int {
	slice := make([]int, len(t))
	copy(slice, t)
	return append(slice[:i], slice[i+1:]...)
}
