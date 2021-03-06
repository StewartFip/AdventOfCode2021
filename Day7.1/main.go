package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/StewartFip/AdventOfCode2021/lib"
)

func main() {
	file, err := os.Open("./Day7.1/input.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(lib.SplitCSV)

	var positions []int
	for scanner.Scan() {
		pos, _ := strconv.Atoi(scanner.Text())
		positions = append(positions, pos)
	}

	analysis := make(map[int]int)
	for _, origin := range positions {
		if _, found := analysis[origin]; found {
			continue
		}

		for _, pos := range positions {
			analysis[origin] += int(math.Abs(float64(origin - pos)))
		}
	}

	var minimumFuel int
	for _, f := range analysis {
		if minimumFuel == 0 || f < minimumFuel {
			minimumFuel = f
		}
	}

	fmt.Println(minimumFuel)
}
