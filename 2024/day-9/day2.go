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
	phaseTwo := SplitBlocks(phaseOne)
	// fmt.Println("phase two: ", phaseTwo)
	phaseThree := AttemptBlockSwap(phaseTwo)
	phaseFour := MatrixToSlice(phaseThree)
	// fmt.Println("phase three: ", MatrixToSlice(phaseThree))
	total := 0
	for i, n := range phaseFour {
		if n == -1 {
			continue
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

func SplitBlocks(t []int) [][]int {
	start := -1
	key := -1
	var out [][]int
	var block []int

	for i, n := range t {
		if start == -1 || i == 0 {
			start = i
			key = n
		} else if n != key {
			block = t[start:i]
			out = append(out, block)
			start = i
			key = n
		} else if i == len(t)-1 {
			block = t[start:]
			out = append(out, block)
		}
	}
	return out
}

func VerifySwapability(blocks [][]int) bool {
	maxLen := 0
	minBlock := 999
	for _, block := range blocks {
		if block[0] == -1 && len(block) > maxLen {
			maxLen = len(block)
		} else if block[0] != -1 && len(block) < minBlock {
			minBlock = len(block)
		}
	}
	// fmt.Println("Block_temp: ", blocks)
	// fmt.Println("maxGap", maxLen)
	// fmt.Println("minBlock", minBlock)
	if minBlock <= maxLen && minBlock != 0 {
		return true
	}
	return false
}

func AttemptBlockSwap(blocks [][]int) [][]int {
	skipCounter := 0
	for {
		idxBlock := GetLastBlock(blocks, skipCounter)
		block := blocks[idxBlock]
		blockLen := len(block)
		// fmt.Println("block: ", block)
		// fmt.Println("blocks:", blocks)
		idxGap := GetBlockGap(blocks[:idxBlock+1], blockLen)
		// if idxBlock < idxGap {
		// 	break
		// }
		if VerifySwapability(blocks[:idxBlock+1]) == false {
			break
		} else if idxGap == -1 {
			skipCounter++
			continue
		} else if GetGap(MatrixToSlice(blocks[:idxBlock])) == -1 {
			break
		}
		gapBlock := blocks[idxGap]
		for j := 0; j < blockLen; j++ {
			gapBlock[j] = block[j]
			block[j] = -1
		}

		//rebuild blocks
		a := MatrixToSlice(blocks)
		blocks = SplitBlocks(a)
	}

	return blocks
}

func MatrixToSlice(matrix [][]int) []int {
	var out []int
	for _, slice := range matrix {
		out = append(out, slice...)
	}
	// fmt.Println("out: ", out)
	return out
}

func GetLastBlock(blocks [][]int, skip int) int {
	for i := len(blocks) - 1; i >= 0; i-- {
		if blocks[i][0] != -1 {
			if skip == 0 {
				return i
			} else {
				skip--
			}
		}
	}
	return -1
}

func GetBlockGap(t [][]int, l int) int {
	for i, n := range t {
		if n[0] == -1 && len(n) >= l {
			return i
		}
	}
	return -1
}
