package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var horizontalPosition, depth, aim, count int

	file, err := os.Open("./Day2.2/input.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		direction, magnitude := vector(scanner.Text())

		switch direction {
		case "forward":
			horizontalPosition += magnitude
			depth += aim * magnitude
			count++
		case "up":
			aim -= magnitude
			count++
		case "down":
			aim += magnitude
			count++
		}
	}

	fmt.Printf("The result of horizontalPosition (%d) * depth (%d) is: %d (%d vectors analyzed)\n", horizontalPosition, depth, horizontalPosition*depth, count)
}

func vector(value string) (string, int) {
	vals := strings.Split(value, " ")

	direction := vals[0]
	magnitude, _ := strconv.Atoi(vals[1])

	return direction, magnitude
}
