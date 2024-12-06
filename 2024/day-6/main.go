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
	// directions := []string{"u", "r", "d", "l"}
	content, err := os.ReadFile("test.txt")
	check(err)
	lineStr := strings.Split(string(content), "\n")
	lineStr = lineStr[:len(lineStr)-1]
	// fmt.Println(string(content))
	guard := LocateGuard(lineStr)
	Travel(lineStr, guard)

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

func Travel(grid []string, guard []int) {
	// var diff int
	directions := []string{"u", "r", "d", "l"}
	iter := 0
	for {
		if guard == nil {
			fmt.Println("iter nil: ", iter)
			// for _, k := range grid {
			// 	fmt.Println(k)
			// }
			total := GetTotalPos(grid)
			fmt.Println(total)
			break
		}
		for _, d := range directions {
			x := guard[0]
			y := guard[1]
			// fmt.Println("iter: ", iter)
			// fmt.Println("Guard:", guard)
			grid[y] = grid[y][:x] + "x" + grid[y][x+1:]
			// for _, k := range grid {
			// 	fmt.Println(k)
			// }
			grid[y] = grid[y][:x] + "." + grid[y][x+1:]
			guard = GetNewPos(grid, guard, d)
			if guard == nil {
				break
			}
			// fmt.Println("diff: ", diff)
			iter++
		}
	}
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
