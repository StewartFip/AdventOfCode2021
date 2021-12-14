package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/StewartFip/AdventOfCode2021/lib"
)

func main() {
	file, err := os.Open("./Day6.1/input.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(lib.SplitCSV)

	fishes := make([]int, 0)
	for scanner.Scan() {
		timer, _ := strconv.Atoi(scanner.Text())
		fishes = append(fishes, timer)
	}

	newFish := make([]int, 0)

	for day := 1; day <= 80; day++ {
		fishes = append(fishes, newFish...)
		newFish = make([]int, 0)

		for i := range fishes {
			if fishes[i] == 0 {
				fishes[i] = 6
			} else {
				fishes[i]--
			}

			if fishes[i] == 0 {
				newFish = append(newFish, 9)
			}
		}
	}

	fmt.Printf("Total fish after 80 days is %d\n", len(fishes))
}
