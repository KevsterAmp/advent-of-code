package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	content, _ := os.ReadFile("input")
	input := strings.Trim(string(content), "\n")
	regex := regexp.MustCompile(`.=(-?\d+),(-?\d+) .=(-?\d+),(-?\d+)`)
	matches := regex.FindAllStringSubmatch(input, -1)
	x := 101
	y := 103
	futureBots := make(map[string]int)

	for _, match := range matches {
		currPos, velocity := BuildConfig(match)
		outPos := CalculatePos(currPos, velocity, x, y, 100)
		key := strconv.Itoa(outPos[0]) + " " + strconv.Itoa(outPos[1])
		futureBots[key] += 1
	}

	quadrants := GetQuadrants(x, y)
	test := InQuadrant(futureBots, quadrants)

	fmt.Println("total:", test)

	// part two
	for i := 0; ; i++ {
		futureBots := make(map[string]int)
		for _, match := range matches {
			currPos, velocity := BuildConfig(match)
			outPos := CalculatePos(currPos, velocity, x, y, i)
			key := strconv.Itoa(outPos[0]) + " " + strconv.Itoa(outPos[1])
			futureBots[key] += 1
		}
		if FindChristmasTree(futureBots) {

			fmt.Println("I: ", i)
			break
		}
	}
}

func FindChristmasTree(futureBots map[string]int) bool {
	for _, value := range futureBots {
		if value > 1 {
			return false
		}
	}
	return true
}

func InPos(pos []int, futurePos [][]int) bool {
	for _, a := range futurePos {
		if a[0] == pos[0] && a[1] == pos[1] {
			return true
		}
	}
	return false
}

func CalculatePos(currPos []int, veloc []int, x, y, i int) []int {
	outPos := make([]int, 2)
	outPos[0] = currPos[0] + (veloc[0] * i)
	outPos[1] = currPos[1] + (veloc[1] * i)
	if outPos[0] > x {
		outPos[0] = outPos[0] % x
	} else if outPos[0] < -(x - 1) {
		if outPos[0]%x != 0 {
			outPos[0] = (outPos[0] % x) + x
		} else {
			outPos[0] = 0
		}
	} else if outPos[0] < 0 {
		outPos[0] += x
	}
	if outPos[1] > y-1 {
		outPos[1] = outPos[1] % y
	} else if outPos[1] < -(y - 1) {
		if outPos[1]%y != 0 {
			outPos[1] = (outPos[1] % y) + y
		} else {
			outPos[1] = 0
		}
	} else if outPos[1] < 0 {
		outPos[1] += y
	}
	return outPos
}

func BuildConfig(match []string) ([]int, []int) {
	var configs []int
	for i, x := range match {
		if i != 0 {
			a, _ := strconv.Atoi(x)
			configs = append(configs, a)
		}
	}
	currPos := []int{configs[0], configs[1]}
	velocity := []int{configs[2], configs[3]}
	return currPos, velocity
}

func GetQuadrants(x, y int) map[string][4]int {
	var midX, midY int
	if x%2 == 1 {
		midX = x/2 - 1
	} else {
		midX = x / 2
	}
	if y%2 == 1 {
		midY = y/2 - 1
	} else {
		midY = y / 2
	}
	maxX := x - 1
	maxY := y - 1

	quadrants := map[string][4]int{
		"Quadrant I":   {midX + 2, maxX, midY + 2, maxY}, // Top-right
		"Quadrant II":  {0, midX, midY + 2, maxY},        // Top-left
		"Quadrant III": {0, midX, 0, midY},               // Bottom-left
		"Quadrant IV":  {midX + 2, maxX, 0, midY},        // Bottom-right
	}

	return quadrants
}

func InQuadrant(futureBots map[string]int, quadrants map[string][4]int) int {
	var minX, maxX, minY, maxY, x, y, count int
	var quadrantsCount []int
	total := 1
	for _, quadrant := range quadrants {
		// fmt.Println(name, quadrant)
		count = 0
		minX = quadrant[0]
		maxX = quadrant[1]
		minY = quadrant[2]
		maxY = quadrant[3]
		for key, value := range futureBots {
			parts := strings.Split(key, " ")
			x, _ = strconv.Atoi(parts[0])
			y, _ = strconv.Atoi(parts[1])
			if x >= minX && x <= maxX && y >= minY && y <= maxY {
				// fmt.Println(x, y, value)
				count += value
				// fmt.Println("count:", count)
			}
		}
		quadrantsCount = append(quadrantsCount, count)
		total *= count
	}
	fmt.Println("all count:", quadrantsCount)
	return total
}
