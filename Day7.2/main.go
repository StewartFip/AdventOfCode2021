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
	file, err := os.Open("./Day7.2/input.txt")
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

	maxOrigin := positions[0]
	for _, p := range positions {
		if p > maxOrigin {
			maxOrigin = p
		}
	}

	analysis := make(map[int]int)
	for origin := 0; origin <= maxOrigin; origin++ {
		if _, found := analysis[origin]; found {
			continue
		}

		for _, pos := range positions {
			steps := int(math.Abs(float64(origin - pos)))
			var cost int

			for i := 1; i <= steps; i++ {
				cost += i
			}

			analysis[origin] += cost
		}
	}

	var minimumFuel int
	for _, f := range analysis {
		if minimumFuel == 0 || f < minimumFuel {
			minimumFuel = f
		}
	}

	// 96,799,391 is too high
	fmt.Println(minimumFuel)
}
