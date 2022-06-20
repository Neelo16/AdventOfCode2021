package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	Board  [][]int
	Marked [][]bool
	Width  int
	Height int
}

func (board Board) String() string {
	s := ""
	for i, line := range board.Board {
		for j, row := range line {
			if board.Marked[i][j] {
				s += fmt.Sprint(" X ")
			} else {
				s += fmt.Sprintf("%2v ", row)
			}
		}
		s += "\n"
	}
	return s
}

func (board Board) Mark(target int) {
	for i, row := range board.Board {
		for j, value := range row {
			if value == target {
				board.Marked[i][j] = true
				return
			}
		}
	}
}

func (board Board) actualVictory(isTransposed bool) bool {
	for _, row := range board.Marked {
		lost := false
		for _, isMarked := range row {
			if !isMarked {
				lost = true
				break
			}
		}
		if !lost {
			return true
		}
	}
	return !isTransposed && board.Transpose().actualVictory(true)
}

func (board Board) Victory() bool {
	return board.actualVictory(false)
}

func (board Board) Transpose() Board {
	transposed := Board{Width: board.Height, Height: board.Width}
	transposed.Board = make([][]int, board.Width)
	transposed.Marked = make([][]bool, board.Width)
	for i := 0; i < board.Width; i++ {
		transposed.Board[i] = make([]int, board.Height)
		transposed.Marked[i] = make([]bool, board.Height)
	}
	for i := 0; i < board.Width; i++ {
		for j := 0; j < board.Height; j++ {
			transposed.Board[j][i] = transposed.Board[i][j]
			transposed.Marked[j][i] = transposed.Marked[i][j]
		}
	}
	return transposed
}

func (board Board) SumUnmarked() int {
	sum := 0
	for i := 0; i < board.Height; i++ {
		for j := 0; j < board.Width; j++ {
			if !board.Marked[i][j] {
				sum += board.Board[i][j]
			}
		}
	}
	return sum
}

func MapToInt(strings []string) []int {
	ints := make([]int, 0, len(strings))
	for _, s := range strings {
		converted, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("Failed to convert %v to int: %v\n", s, err)
		}
		ints = append(ints, converted)
	}
	return ints
}

func LinesToBoard(lines []string) Board {
	board := Board{Width: len(strings.Fields(lines[0])), Height: len(lines), Board: make([][]int, 0, len(lines))}
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) != board.Width {
			log.Fatalf("Invalid line length: should be %v but was %v\n", board.Width, len(line))
		}
		board.Board = append(board.Board, MapToInt(fields))
	}
	board.Marked = make([][]bool, board.Height)
	for i := range board.Marked {
		board.Marked[i] = make([]bool, board.Width)
	}
	return board
}

func ReadInput() (draws []int, boards []Board) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	draws = MapToInt(strings.Split(scanner.Text(), ","))
	scanner.Scan()

	currentBoard := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			boards = append(boards, LinesToBoard(currentBoard))
			currentBoard = make([]string, 0)
		} else {
			currentBoard = append(currentBoard, line)
		}
	}
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	boards = append(boards, LinesToBoard(currentBoard))

	return
}

func FindNthVictor(draws []int, boards []Board, n int) (winningDraw int, winningBoard Board, err error) {
	victors := 0
	for _, draw := range draws {
		remainingBoards := make([]Board, 0, len(boards))
		for _, board := range boards {
			board.Mark(draw)
			if board.Victory() {
				victors++
				winningDraw, winningBoard = draw, board
			} else {
				remainingBoards = append(remainingBoards, board)
			}
			if victors == n {
				return
			}
		}
		boards = remainingBoards
	}
	return 0, Board{}, errors.New("failed to find board")
}

func main() {
	draws, boards := ReadInput()
	winningDraw, winningBoard, err := FindNthVictor(draws, boards, 1)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("First star: %v\n", winningBoard.SumUnmarked()*winningDraw)
	draws, boards = ReadInput()
	winningDraw, winningBoard, err = FindNthVictor(draws, boards, len(boards)-1)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Second star: %v\n", winningBoard.SumUnmarked()*winningDraw)
}
