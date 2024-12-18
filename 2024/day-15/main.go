package main

import (
	"fmt"
	"os"
	// "regexp"
	// "strconv"
	"strings"
)

func main() {
	content, _ := os.ReadFile("test")
	input := strings.Split(string(content), "\n\n")
	grid := strings.Split(input[0], "\n")
	robot := LocateRobot(grid)
	fmt.Println(robot)
	moves := input[1]
	Travel(robot, grid, moves)
}

func LocateRobot(grid []string) []int {
	for y, line := range grid {
		for x, char := range line {
			if string(char) == "@" {
				return []int{x, y}
			}
		}
	}
	return nil
}

func Travel(robot []int, grid []string, moves string) {
	x := robot[0]
	y := robot[0]
	for _, c := range moves {
		move := string(c)
		switch move {
		case ">":
			if string(grid[y][x+1]) == "#" {
				continue
			} else if string(grid[y][x+1]) == "0" {
				fmt.Println()
			}
		}
	}
}

func MoveLinearBoxes(x int, y int, grid []string, move string) {
	boxLen := 0
	switch move {
	case ">":
		for i := x; i < len(grid[0]); x++ {
			currCell := string(grid[y][i])
			if currCell == "0" {
				boxLen++
			} else if currCell == "#" {
				return
			} else {
				break
			}
		}

		grid[y] = grid[y][:x] + "." + strings.Repeat("#", boxLen) + grid[y][x+boxLen+2:]
	}
}
