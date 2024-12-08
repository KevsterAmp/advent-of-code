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
	grid := strings.Split(string(content), "\n")
	grid = grid[:len(grid)-1]
	freshGrid := make([]string, len(grid))
	copy(freshGrid, grid)
	fmt.Println(string(content))
	m := GetAntenna(grid)
	fmt.Println(m)
	for _, antennas := range m {
		pairs := CreatePairs(antennas)
		// fmt.Println("pairs: ", pairs)
		CreateAntinodes(pairs, freshGrid)
	}
	count := GetAntinodeCount(freshGrid)
	fmt.Println("Count: ", count)
}

func GetAntenna(grid []string) map[string][][]int {
	m := make(map[string][][]int)
	for y, line := range grid {
		for x, a := range line {
			char := string(a)
			if char != "." && char != "#" {
				if In(char, m) {
					m[char] = append(m[char], []int{x, y})
				} else {
					m[char] = [][]int{{x, y}}
				}
			}
		}
	}
	return m
}

func In(element string, m map[string][][]int) bool {
	for x := range m {
		if x == element {
			return true
		}
	}
	return false
}

func CreatePairs(antennas [][]int) [][][]int {
	var out [][][]int

	for i, x := range antennas {
		for j, y := range antennas {
			if i == j || i > j {
				continue
			}
			// fmt.Println(x + "|" + y)
			out = append(out, [][]int{x, y})
		}
	}
	return out
}

func CreateAntinodes(pairs [][][]int, grid []string) {
	for _, pair := range pairs {
		aPair := pair[0]
		bPair := pair[1]
		// fmt.Println("Pairs: ", pair)
		xDiff := GetDiff(aPair[0], bPair[0])
		yDiff := GetDiff(aPair[1], bPair[1])
		// fmt.Println("diffs: ", xDiff, yDiff)
		antinodes := [][]int{
			make([]int, 2),
			make([]int, 2),
		}
		for i := 0; i < 2; i++ {
			g := aPair[i] >= bPair[i]
			if g && i == 0 {
				antinodes[0][i] = aPair[i] + xDiff
				antinodes[1][i] = bPair[i] - xDiff
			} else if g == false && i == 0 {
				antinodes[0][i] = aPair[i] - xDiff
				antinodes[1][i] = bPair[i] + xDiff
			} else if g && i == 1 {
				antinodes[0][i] = aPair[i] + yDiff
				antinodes[1][i] = bPair[i] - yDiff
			} else if g == false && i == 1 {
				antinodes[0][i] = aPair[i] - yDiff
				antinodes[1][i] = bPair[i] + yDiff
			}
		}
		for _, node := range antinodes {
			// fmt.Println("Node: ", node)
			x := node[0]
			y := node[1]
			if OutOfBounds(x, "x", grid) == false && OutOfBounds(y, "y", grid) == false && string(grid[y][x]) != "#" {
				grid[y] = string(grid[y][:x]) + "#" + string(grid[y][x+1:])
				// fmt.Println("Grid")
				// for _, line := range grid {
				// fmt.Println(line)
				// }
			}
		}
	}
}

func GetDiff(x int, y int) int {
	diff := x - y
	if diff < 0 {
		diff += (diff * -2)
	}
	return diff
}

func OutOfBounds(c int, t string, grid []string) bool {
	if t == "y" && c >= 0 && c < len(grid) {
		return false
	} else if t == "x" && c >= 0 && c < len(grid[0]) {
		return false
	}
	return true
}

func GetAntinodeCount(grid []string) int {
	count := 0
	for _, line := range grid {
		for _, c := range line {
			if string(c) == "#" {
				count++
			}
		}
	}
	return count
}
