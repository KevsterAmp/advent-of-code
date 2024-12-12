package main

import (
	"fmt"
	"os"
	// "regexp"
	// "strconv"
	// "strings"
)

func main() {
	content, _ := os.ReadFile("input.txt")
	fmt.Println(string(content))
	// count := numIslands(content)
	// fmt.Println("coutn: ", count)
}

func numIslands(input [][]byte) int {
	// data := strings.Split(string(input), "\n")
	var grid []string

	for _, k := range input {
		grid = append(grid, string(k))
	}

	var seen [][]int
	count := 0
	for y, line := range grid {
		for x, c := range line {
			if string(c) == "1" && InOutput([]int{x, y}, seen) == false {
				island := GetIsland([]int{x, y}, grid)
				count++
				for _, box := range island {
					seen = append(seen, box)
				}
			}
		}
	}
	return count
}

func GetIsland(start []int, grid []string) [][]int {
	xMax := len(grid[0])
	yMax := len(grid)
	var island [][]int
	var queue [][]int
	var newQueue [][]int
	queue = append(queue, start)

	for {
		for _, idxs := range queue {
			x := idxs[0]
			y := idxs[1]
			for i := -1; i <= 1; i += 2 {
				if x+i < xMax && x+i >= 0 && string(grid[y][x+i]) == "1" {
					outSpace := []int{x + i, y}
					if InOutput(outSpace, island) == false {
						newQueue = append(newQueue, outSpace)
						island = append(island, outSpace)
					}
				}
				if y+i < yMax && y+i >= 0 && string(grid[y+i][x]) == "1" {
					outSpace := []int{x, y + i}
					if InOutput(outSpace, island) == false {
						newQueue = append(newQueue, outSpace)
						island = append(island, outSpace)
					}
				}
			}
		}
		if len(newQueue) == 0 {
			break
		}
		queue = newQueue
		newQueue = [][]int{}
	}
	return island
}

func InOutput(e []int, t [][]int) bool {
	for _, idx := range t {
		if idx[0] == e[0] && idx[1] == e[1] {
			return true
		}
	}
	return false
}
