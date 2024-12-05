package main

import (
	"fmt"
	"os"
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

	contentStr := string(content)
	// fmt.Println(contentStr)

	// get horizontal
	data := strings.Split(contentStr, "\n")
	data = data[:len(data)-1]
	// fmt.Println("data: ", data)

	match := 0
	dataLen := len(data)
	for i := 1; i < dataLen-1; i++ {
		lineLen := len(data[i])
		for j := 1; j < lineLen-1; j++ {
			if SoftMatch(i, j, data) == true && HardMatch(i, j, data) == true {
				match++
			}
		}
	}
	fmt.Println("Match: ", match)

}

func SoftMatch(i int, j int, t []string) bool {
	if string(t[i][j]) != "A" {
		return false
	}
	ops := []int{-1, 1}
	for _, x := range ops {
		for _, y := range ops {
			if IsMatch(string(t[i+x][j+y])) == false {
				return false
			}
			// fmt.Println(x, y)
		}
	}
	return true
}

func IsMatch(c string) bool {
	return c == "M" || c == "S"
}

func HardMatch(i int, j int, t []string) bool {
	combination := [][]int{
		{i - 1, j + 1, i + 1, j - 1},
		{i + 1, j + 1, i - 1, j - 1},
	}
	for _, x := range combination {
		if XmasMatch(t, x) == false {
			return false
		}
	}
	return true
}

func XmasMatch(t []string, x []int) bool {
	return (string(t[x[0]][x[1]]) == "M" && string(t[x[2]][x[3]]) == "S") || (string(t[x[0]][x[1]]) == "S" && string(t[x[2]][x[3]]) == "M")
}
