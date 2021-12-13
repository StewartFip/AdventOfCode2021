package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./Day4.2/input.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Scan once to get the Draw Numbers
	scanner.Scan()
	drawNumbers := strings.Split(scanner.Text(), ",")

	// Scan one more time to get past the blank line
	scanner.Scan()

	boards := make([]string, 0, 100)

	var board string

	for scanner.Scan() {
		switch scanner.Text() {
		case "":
			boards = append(boards, board)

			board = ""
		default:
			board += scanner.Text() + "\n"
		}
	}

	winner, winningDraw, i := play(boards[0:], drawNumbers)
	if winner == "" {
		fmt.Println("No winner was found!")
		os.Exit(1)
	}

	winningNumber, err := strconv.Atoi(winningDraw)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	s, err := score(winner, winningNumber)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("The winning board number is %d on a draw of %s\n...and looks like this afterward:\n%s\nThe final score is: %d\n", i, winningDraw, winner, s)
}

func play(boards, drawNumbers []string) (string, string, int) {
	fmt.Println(len(boards))

	wb := newWinningBoards()

	for _, number := range drawNumbers {
		repl := regexp.MustCompile(fmt.Sprintf("%2s( |\n)", number))
		for i, board := range boards {
			if wb.contains(i) {
				continue
			}

			boards[i] = mark(board, repl)

			if winner(boards[i]) {
				wb.winners = append(wb.winners, i)

				fmt.Printf("Winner Winner Chicken Dinner!\n%s\n", boards[i])
				fmt.Println(number)

				fmt.Println(len(wb.winners))

				if len(wb.winners) == 99 {
					return boards[i], number, i
				}
			}
		}
	}

	return "", "", 0
}

func mark(board string, repl *regexp.Regexp) string {
	matches := repl.FindAllStringIndex(board, -1)

	for _, match := range matches {
		board = board[0:match[0]] + "  " + board[match[1]-1:]
	}

	return board
}

func winner(board string) bool {
	rows := strings.Split(board, "\n")
	for i := 0; i <= 12; i += 3 {
		var loser bool

		for _, row := range rows[:len(rows)-1] {
			if row[i:i+2] != "  " {
				loser = true

				break
			}
		}

		if !loser {
			return true
		}
	}

	return strings.Contains(board, fmt.Sprintf("%15s", "\n"))
}

func score(board string, winningNumber int) (int, error) {
	var s int

	rows := strings.Split(board, "\n")
	for _, row := range rows[:len(rows)-1] {
		for i := 0; i <= 12; i += 3 {
			if row[i:i+2] != "  " {
				val, err := strconv.Atoi(strings.TrimSpace(row[i : i+2]))
				if err != nil {
					return 0, err
				}

				s += val
			}
		}
	}

	return s * winningNumber, nil
}

type winningBoards struct {
	winners []int
}

func newWinningBoards() *winningBoards {
	return &winningBoards{
		winners: make([]int, 0, 100),
	}
}

func (wb *winningBoards) contains(number int) bool {
	for _, b := range wb.winners {
		if number == b {
			return true
		}
	}

	return false
}
