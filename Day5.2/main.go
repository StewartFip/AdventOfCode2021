package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./Day5.2/input.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	coords := coordCount(make(map[string]int))
	for scanner.Scan() {
		coords.append(scanner.Text())
	}

	var count int

	for _, c := range coords {
		if c >= 2 {
			count++
		}
	}

	fmt.Printf("The number of intersecting coordinates >= 2 is %d\n", count)
}

type coordCount map[string]int

func (c coordCount) append(coordRange string) {
	for _, coord := range split(coordRange) {
		c[coord]++
	}
}

func split(coordRange string) []string {
	coords := make([]string, 0)

	beg, end := parse(strings.Split(coordRange, " -> "))

	switch {
	case beg[0] == end[0] || beg[1] == end[1]:
		for x := beg[0]; x <= end[0]; x++ {
			for y := beg[1]; y <= end[1]; y++ {
				coords = append(coords, fmt.Sprintf("%d,%d", x, y))
			}
		}
	case beg[0] < end[0]:
		for x := beg[0]; x <= end[0]; x++ {
			for y := beg[1]; y <= end[1]; y++ {
				coords = append(coords, fmt.Sprintf("%d,%d", x, y))
				x++
			}
		}
	case beg[0] > end[0]:
		for x := beg[0]; x >= end[0]; x-- {
			for y := beg[1]; y <= end[1]; y++ {
				coords = append(coords, fmt.Sprintf("%d,%d", x, y))
				x--
			}
		}
	}

	return coords
}

func parse(begEnd []string) ([]int, []int) {
	beg := make([]int, 2)
	end := make([]int, 2)

	beg[0], _ = strconv.Atoi(begEnd[0][:strings.Index(begEnd[0], ",")])
	beg[1], _ = strconv.Atoi(begEnd[0][strings.Index(begEnd[0], ",")+1:])
	end[0], _ = strconv.Atoi(begEnd[1][:strings.Index(begEnd[1], ",")])
	end[1], _ = strconv.Atoi(begEnd[1][strings.Index(begEnd[1], ",")+1:])

	switch {
	case beg[0] == end[0] && beg[1] > end[1]:
		return end, beg
	case beg[1] == end[1] && beg[0] > end[0]:
		return end, beg
	case beg[0] < end[0] && beg[1] > end[1]:
		return end, beg
	case beg[0] > end[0] && beg[1] > end[1]:
		return end, beg
	default:
		return beg, end
	}
}
