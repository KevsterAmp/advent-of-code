package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	total := 0
	content, _ := os.ReadFile("input")
	input := strings.Split(string(content), "\n\n")
	regex := regexp.MustCompile(`\d+`)
	for _, config := range input {
		match := regex.FindAllString(config, -1)
		out := CramersRule(match)
		if out != nil {
			total += (out[0] * 3) + out[1]
		}
	}
	fmt.Println(total)
}

func CramersRule(t []string) []int {
	ax, _ := strconv.Atoi(t[0])
	bx, _ := strconv.Atoi(t[1])
	ay, _ := strconv.Atoi(t[2])
	by, _ := strconv.Atoi(t[3])
	c1, _ := strconv.Atoi(t[4])
	c2, _ := strconv.Atoi(t[5])
	c1 += 10000000000000
	c2 += 10000000000000

	d := (ax * by) - (bx * ay)
	dx := (c1 * by) - (c2 * ay)
	dy := (c2 * ax) - (c1 * bx)
	if d != 0 && dx%d == 0 && dy%d == 0 {
		return []int{dx / d, dy / d}
	}
	return nil
}
