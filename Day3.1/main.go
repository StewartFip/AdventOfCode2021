package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./Day3.1/input.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer file.Close()

	pos := make([][]int, 12)
	for i := range pos {
		pos[i] = make([]int, 2)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		for i, b := range line {
			switch b {
			case '0':
				pos[i][0]++
			case '1':
				pos[i][1]++
			}
		}
	}

	var gammaBinary, epsilonBinary string
	for _, ge := range pos {
		switch {
		case ge[0] > ge[1]:
			gammaBinary += "0"
			epsilonBinary += "1"
		case ge[1] > ge[0]:
			gammaBinary += "1"
			epsilonBinary += "0"
		}

		fmt.Println(ge)
	}

	gamma, err := strconv.ParseInt(gammaBinary, 2, 64)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	epsilon, err := strconv.ParseInt(epsilonBinary, 2, 64)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("The result of gamma (%s, %d) * epsilon (%s, %d) is: %d\n", gammaBinary, gamma, epsilonBinary, epsilon, gamma*epsilon)
}
