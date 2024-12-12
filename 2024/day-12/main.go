package main

import (
	"fmt"
	"os"
	// "regexp"
	// "strconv"
	"strings"
)

func main() {
	content, _ := os.ReadFile("input")
	input := strings.Trim(string(content), "\n")
	grid := strings.Split(input, "\n")

	for _, line := range grid {
		fmt.Println(line)
	}

	PartOne(grid)
}

func PartOne(grid []string) {
	var seen [][]int
	totalPrice := 0
	totalPrice2 := 0

	for y, line := range grid {
		for x, c := range line {
			idx := []int{x, y}
			if InRegion(idx, seen) == false {
				test, area, perimiter, sides := Bsf(idx, grid)
				fmt.Println("region:", string(c), "area:", area, "perimiter:", perimiter, "sides", sides)
				// fmt.Println("cells", test)
				seen = append(seen, test...)
				totalPrice += area * perimiter
				totalPrice2 += area * sides
			}
		}
	}
	fmt.Println("Total price:", totalPrice)
	fmt.Println("Total price2:", totalPrice2)
}

func Bsf(start []int, grid []string) ([][]int, int, int, int) {
	totalPerimiter := 0
	sides := 0
	xSides := make(map[int][][]int)
	ySides := make(map[int][][]int)
	regions := [][]int{start}
	queue := [][]int{start}
	for n := 0; ; n++ {
		newQueue := [][]int{}
		for _, cell := range queue {
			x := cell[0]
			y := cell[1]
			region := GetRegion(cell, grid)
			currPerimiter := 4
			for i := -1; i <= 1; i += 2 {
				xAdjacent := []int{x + i, y}
				yAdjacent := []int{x, y + i}

				// fmt.Println("xadj:", GetRegion(xAdjacent, grid), "idx", xAdjacent)
				// fmt.Println("yadj:", GetRegion(yAdjacent, grid), "idx", yAdjacent)
				if GetRegion(xAdjacent, grid) == region {
					currPerimiter--
					if InRegion(xAdjacent, regions) == false {
						newQueue = append(newQueue, xAdjacent)
						regions = append(regions, xAdjacent)
					}
				} else if InMap(x+i, xSides) == false || InRegion(cell, xSides[x+i]) == false {
					fmt.Println("x+i", x+i)
					fmt.Println("cell", cell)
					sides++
					sideRegions := GetSideRegions(cell, i, grid, "x")
					fmt.Println("sideRegion:", sideRegions)
					for _, sideRegion := range sideRegions {
						if InRegion(sideRegion, xSides[x+i]) == false {
							xSides[x+i] = append(xSides[x+i], sideRegion)
						} else {
							fmt.Println("sideregion:", sideRegion)
						}
					}
				}
				if GetRegion(yAdjacent, grid) == region {
					currPerimiter--
					if InRegion(yAdjacent, regions) == false {
						newQueue = append(newQueue, yAdjacent)
						regions = append(regions, yAdjacent)
					}
				} else if InMap(y+i, ySides) == false || InRegion(cell, ySides[y+i]) == false {
					sides++
					sideRegions := GetSideRegions(cell, i, grid, "y")
					for _, sideRegion := range sideRegions {
						if InRegion(sideRegion, ySides[y+i]) == false {
							ySides[y+i] = append(ySides[y+i], sideRegion)
						}
					}
				}
			}
			// fmt.Println("queue", queue)
			// fmt.Println("Newqueue", newQueue)
			queue = newQueue
			totalPerimiter += currPerimiter
		}

		if len(queue) == 0 {
			fmt.Println("xSides", xSides)
			fmt.Println("ySides", ySides)
			break
		}
	}
	return regions, len(regions), totalPerimiter, sides
}

// this can be optimized but it does the job
func GetSideRegions(cell []int, i int, grid []string, direction string) [][]int {
	x := cell[0]
	y := cell[1]
	region := GetRegion(cell, grid)

	sideRegions := [][]int{cell}
	if direction == "y" {
		for j := x + 1; j < len(grid); j++ {
			if GetRegion([]int{j, y}, grid) == region && GetRegion([]int{j, y + i}, grid) != region {
				sideRegions = append(sideRegions, []int{j, y})
			} else {
				break
			}
		}
		for j := x - 1; j >= 0; j-- {
			if GetRegion([]int{j, y}, grid) == region && GetRegion([]int{j, y + i}, grid) != region {
				sideRegions = append(sideRegions, []int{j, y})
			} else {
				break
			}
		}
	} else if direction == "x" {
		for j := y + 1; j < len(grid); j++ {
			if GetRegion([]int{x, j}, grid) == region && GetRegion([]int{x + i, j}, grid) != region {
				sideRegions = append(sideRegions, []int{x, j})
			} else {
				break
			}
		}
		for j := y - 1; j >= 0; j-- {
			if GetRegion([]int{x, j}, grid) == region && GetRegion([]int{x + i, j}, grid) != region {
				sideRegions = append(sideRegions, []int{x, j})
			} else {
				break
			}
		}

	}
	return sideRegions
}

func GetRegion(idxs []int, grid []string) string {
	x := idxs[0]
	y := idxs[1]
	xMax := len(grid[0])
	yMax := len(grid)
	if x < xMax && x >= 0 && y < yMax && y >= 0 {
		return string(grid[y][x])
	} else {
		return ""
	}
}

func InRegion(cell []int, t [][]int) bool {
	for _, x := range t {
		if x[0] == cell[0] && x[1] == cell[1] {
			return true
		}
	}
	return false
}

func In(x int, t []int) bool {
	for _, n := range t {
		if n == x {
			return true
		}
	}
	return false
}

func InMap(key int, record map[int][][]int) bool {
	_, exists := record[key]
	return exists
}
