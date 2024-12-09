package main

import (
	"fmt"
	"os"
	"strconv"
	// "regexp"
	// "strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	content, err := os.ReadFile("input.txt")
	check(err)
	input := string(content)
	// fmt.Println("input: ", input)
	id := 0
	var phaseOne []int
	for i := 0; i < len(input); i += 2 {
		blockStorage, _ := strconv.Atoi(string(input[i]))
		freeSpace, _ := strconv.Atoi(string(input[i+1]))
		// fmt.Println("Blockstorage", blockStorage)
		// fmt.Println("freespace", freeSpace)
		for j := 0; j < blockStorage; j++ {
			phaseOne = append(phaseOne, id)
		}
		for k := 0; k < freeSpace; k++ {
			phaseOne = append(phaseOne, -1)
		}
		id++
	}
	// fmt.Println("phase one: ", phaseOne)
	// var phaseTwo []int
	// copy(phaseTwo, phaseOne)
	counter := 0
	for {
		fmt.Println("iter: ", counter)
		idxNum := GetLeftMostNum(phaseOne)
		idxGap := GetGap(phaseOne)
		// fmt.Println("idxs", idxNum, idxGap)
		if idxNum != -1 && idxGap != -1 {
			phaseOne[idxGap] = phaseOne[idxNum]
			phaseOne[idxNum] = -1
		} else {
			break
		}
		counter++
	}
	// fmt.Println("phase two: ", phaseOne)
	total := 0
	for i, n := range phaseOne {
		if n == -1 {
			break
		}
		total += n * i
	}
	fmt.Println("total: ", total)
}

func GetLeftMostNum(t []int) int {
	for i := len(t) - 1; i >= 0; i-- {
		if t[i] != -1 {
			// fmt.Println("Leftmost: ", t[i])
			return i
		}
	}
	return -1
}

func GetGap(t []int) int {
	idx := GetLeftMostNum(t)
	tConcat := t[:idx+1]
	// fmt.Println("concat: ", tConcat)
	for i, n := range tConcat {
		if n == -1 {
			return i
		}
	}
	return -1
}
