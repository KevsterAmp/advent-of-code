package main

import (
	"fmt"
	"os"
	// "strconv"
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
	lineStr := strings.Split(string(content), "\n")
	lineStr = lineStr[:len(lineStr)-1]
	freshGrid := make([]string, len(lineStr))
	copy(freshGrid, lineStr)
	// fmt.Println(string(content))
	// guard := LocateGuard(lineStr)
	// total := Travel(lineStr, guard, true)
	// fmt.Println("cell total: ", total)
	// PartOne(lineStr)
	PartTwo(freshGrid)
}

func PartOne(lineStr []string) {
	guard := LocateGuard(lineStr)
	total := Travel(lineStr, guard, true)
	fmt.Println("cell total: ", total)
}

func PartTwo(lineStr []string) {
	counter := 0
	for y, line := range lineStr {
		for x, char := range line {
			if string(char) == "." {
				freshGrid := make([]string, len(lineStr))
				copy(freshGrid, lineStr)
				freshGrid[y] = freshGrid[y][:x] + "#" + freshGrid[y][x+1:]
				guard := LocateGuard(freshGrid)
				isLoop := Travel(freshGrid, guard, true)
				if isLoop == -1 {
					counter++
				}
			}
		}
	}
	fmt.Println("Counter: ", counter)
}

func LocateGuard(grid []string) []int {
	for y, line := range grid {
		for x, char := range line {
			if string(char) == "^" {
				return []int{x, y}
			}
		}
	}
	return nil
}

func Travel(grid []string, guard []int, checkLoop bool) int {
	loopCounter := 0
	var guardHistory [][]int
	directions := []string{"u", "r", "d", "l"}
	iter := 0
	for {
		for _, d := range directions {
			x := guard[0]
			y := guard[1]
			// fmt.Println("iter: ", iter)
			// fmt.Println("Guard:", guard)
			grid[y] = grid[y][:x] + "X" + grid[y][x+1:]
			// for _, k := range grid {
			// 	fmt.Println(k)
			// }
			grid[y] = grid[y][:x] + "." + grid[y][x+1:]
			guardHistory = append(guardHistory, guard)
			guard = GetNewPos(grid, guard, d)
			if guard == nil {
				// fmt.Println("iter nil: ", iter)
				// for _, k := range grid {
				// 	fmt.Println(k)
				// }
				total := GetTotalPos(grid)
				return total
			}
			if checkLoop && IsLoop(guardHistory, guard) && iter >= 4 {
				loopCounter++
				// fmt.Println("iter nil: ", iter)
				// for _, k := range grid {
				// 	fmt.Println(k)
				// }
			}
			if loopCounter == 4 {
				fmt.Println("iter nil: ", iter)
				for _, k := range grid {
					fmt.Println(k)
				}
				return -1
			}

			iter++

		}
	}
}

func IsLoop(guardHistory [][]int, currentGuard []int) bool {
	for _, guard := range guardHistory {
		if currentGuard[0] == guard[0] && currentGuard[1] == guard[1] {
			return true
		}
	}
	return false
}

func GetNewPos(grid []string, guard []int, d string) []int {
	x := guard[0]
	y := guard[1]

	switch d {
	case "u":
		for i := y; i >= 0; i-- {
			if string(grid[i][x]) == "#" {
				// return -1 element
				return []int{x, i + 1}
			} else {
				grid[i] = grid[i][:x] + "X" + grid[i][x+1:]
			}
		}
	case "r":
		for i := x; i < len(grid[0]); i++ {
			if string(grid[y][i]) == "#" {
				// return -1 element
				return []int{i - 1, y}
			} else {
				grid[y] = grid[y][:i] + "X" + grid[y][i+1:]
			}
		}
	case "d":
		for i := y; i < len(grid); i++ {
			if string(grid[i][x]) == "#" {
				// return -1 element
				return []int{x, i - 1}
			} else {
				grid[i] = grid[i][:x] + "X" + grid[i][x+1:]
			}
		}
	case "l":
		for i := x; i >= 0; i-- {
			if string(grid[y][i]) == "#" {
				// return -1 element
				return []int{i + 1, y}
			} else {
				grid[y] = grid[y][:i] + "X" + grid[y][i+1:]
			}
		}
	}
	return nil
}

func GetTotalPos(grid []string) int {
	total := 0
	for _, line := range grid {
		for _, x := range line {
			if string(x) == "X" {
				total++
			}
		}
	}
	return total
}
