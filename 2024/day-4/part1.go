package main

import (
	"fmt"
	"os"
	// "strconv"
	"regexp"
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

	lineStr := string(content)
	// fmt.Println(lineStr)

	// get horizontal
	hArray := strings.Split(lineStr, "\n")
	hArrayRev := ReverseArr(hArray)
	// fmt.Println(hArray)
	// fmt.Println(hArrayRev)

	// get vertical
	var vArray []string
	for j, array := range hArray {
		for i := 0; i < len(array); i++ {
			if j == 0 {
				vArray = append(vArray, string(array[i]))
			} else {
				vArray[i] = vArray[i] + string(array[i])
			}
		}
	}
	vArrayRev := ReverseArr(vArray)
	// fmt.Println(vArray)
	// fmt.Println(vArrayRev)

	// get diagonal l start
	dArrayL := GetDiagonalL(hArray)
	dArrayLRev := ReverseArr(dArrayL)
	// fmt.Println(dArrayL)
	// fmt.Println(dArrayLRev)

	// get diagonal R start
	dArrayR := GetDiagonalR(hArray)
	dArrayRRev := ReverseArr(dArrayR)
	// fmt.Println(dArrayR)
	// fmt.Println(dArrayRRev)

	total := 0
	total += FindXMAS(hArray)
	total += FindXMAS(hArrayRev)
	total += FindXMAS(vArrayRev)
	total += FindXMAS(vArray)
	total += FindXMAS(dArrayL)
	total += FindXMAS(dArrayLRev)
	total += FindXMAS(dArrayR)
	total += FindXMAS(dArrayRRev)

	fmt.Println("total: ", total)
}

func FindXMAS(t []string) int {
	total := 0
	r, err := regexp.Compile(`XMAS`)
	check(err)
	for _, array := range t {
		if len(array) < 4 {
			continue
		}
		out := r.FindAllString(array, -1)
		total += len(out)
	}
	return total
}

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func ReverseArr(t []string) []string {
	var result []string
	for _, str := range t {
		rev := Reverse(str)
		result = append(result, rev)
	}
	return result
}

func GetDiagonalL(hArray []string) []string {
	var dArray []string
	yLen := len(hArray) - 1
	m := ""
	for x := 0; x < len(hArray[0]); x++ {
		m += string(hArray[x][x])
		u := ""
		d := ""
		for y := 0; y < yLen; y++ {
			if x+y >= yLen {
				break
			} else if x == 0 {
				continue
			}
			u += string(hArray[y][x+y])
			d += string(hArray[x+y][y])
			// fmt.Println(y, x+y)
			// fmt.Println(y+x, y)
			// 01,12,23
			// 10,21,
		}
		// fmt.Println(u)
		// fmt.Println(d)
		if x != 0 {
			dArray = append(dArray, u)
			dArray = append(dArray, d)
		}
	}
	// fmt.Println(m)
	dArray = append(dArray, m)
	// fmt.Println(dArray)
	return dArray
}

func GetDiagonalR(hArray []string) []string {
	var dArray []string
	yLen := len(hArray) - 1
	// fmt.Println("yLen: ", yLen)
	xLen := len(hArray[0]) - 1
	for x := xLen; x >= 0; x-- {
		u := ""
		d := ""
		for y := 0; y < yLen; y++ {
			if x-y < 0 {
				continue
			}
			u += string(hArray[y][x-y])
			diff := yLen - x - 1
			d += string(hArray[y+diff][yLen-1-y])
			// fmt.Println(y, x-y)
			// fmt.Println(y+diff, 9-y)
		}
		// fmt.Printf("\n")
		// fmt.Println(u)
		// fmt.Println(d)
		dArray = append(dArray, u)
		if x != xLen {
			dArray = append(dArray, d)
		}
	}
	// fmt.Println(dArray)
	return dArray
}
