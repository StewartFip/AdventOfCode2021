package main

import (
	"bufio"
	"fmt"
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
}
