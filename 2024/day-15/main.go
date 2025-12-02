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
	fmt.Println("grid:", input[0])
	grid := BuildGrid(input[0])
	robot := LocateRobot(grid)
	fmt.Println("robot:", robot)
	moves := input[1]
	Travel(robot, grid, moves)
}

func BuildGrid(input string) [][]string {
	var out [][]string
	for _, line := range strings.Split(input, "\n") {
		lineArr := strings.Split(line, "")
		out = append(out, lineArr)
	}
	return out
}

func LocateRobot(grid [][]string) []int {
	for y, line := range grid {
		for x, char := range line {
			if string(char) == "@" {
				return []int{x, y}
			}
		}
	}
	return nil
}

func Travel(robot []int, grid [][]string, moves string) {
	x := robot[0]
	y := robot[0]
	for _, c := range moves {
		move := string(c)
		switch move {
		case ">":
			out := MoveRobot(x+1, y, grid, robot)
		case "<":
			out := MoveRobot(x-1, y, grid, robot)
		case "^":
			out := MoveRobot(x, y+1, grid, robot)
		case "v":
			out := MoveRobot(x, y-1, grid, robot)
		}
	}
}

func MoveRobot(x int, y int, grid [][]string, robot []int) int {
	if grid[y][x] == "#" {
		return 0
	} else if grid[y][x] == "." {
		grid[y][x] = "@"
		xRob := robot[0]
		yRob := robot[1]
		grid[yRob][xRob] = "."
		return 0
	} else if grid[y][x] == "O" {
		return 1
	}
	return -1
}

func PushBoxes(x, y)
