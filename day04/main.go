package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	puzzle := flag.Int("puzzle", 1, "Which puzzle to run")
	test := flag.Bool("test", false, "Use test inputs")
	flag.Parse()

	if *puzzle == 1 {
		Puzzle1(*test)
	} else {
		Puzzle2(*test)
	}
}

func Puzzle1(useTest bool) {
	filePath := "./inputs/01.txt"
	if useTest {
		filePath = "./inputs/01-test.txt"
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := CreateGrid(scanner)

	count := CountXMAS(grid)
	fmt.Println(count)
}

func Puzzle2(useTest bool) {
	filePath := "./inputs/01.txt"
	if useTest {
		filePath = "./inputs/01-test.txt"
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := CreateGrid(scanner)
	count := CountX_MAS(grid)
	fmt.Println(count)
}

func CreateGrid(scanner *bufio.Scanner) [][]rune {
	grid := make([][]rune, 0)
	for scanner.Scan() {
		row := make([]rune, 0)
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		for _, char := range line {
			row = append(row, char)
		}
		grid = append(grid, row)
	}

	return grid
}

func CountXMAS(input [][]rune) int {
	count := 0
	for i := range input {
		for j := range input[i] {
			if input[i][j] == 'X' {
				if CheckHorizontal(input, i, j) {
					count++
				}

				if CheckVerticalDown(input, i, j) {
					count++
				}

				if CheckVerticalUp(input, i, j) {
					count++
				}

				if CheckDiagonalDownRight(input, i, j) {
					count++
				}

				if CheckDiagonalDownLeft(input, i, j) {
					count++
				}

				if CheckDiagonalUpRight(input, i, j) {
					count++
				}

				if CheckDiagonalUpLeft(input, i, j) {
					count++
				}

				if CheckBackwords(input, i, j) {
					count++
				}
			}
		}
	}

	return count
}

func CountX_MAS(input [][]rune) int {
	count := 0
	for i := range input {
		for j := range input[i] {
			if input[i][j] == 'A' {
				if j > len(input[i])-2 ||
					i > len(input)-2 ||
					j < 1 ||
					i < 1 {
					continue
				}

				if (input[i-1][j-1] == 'M' && input[i+1][j+1] == 'S' ||
					input[i-1][j-1] == 'S' && input[i+1][j+1] == 'M') &&
					(input[i-1][j+1] == 'M' && input[i+1][j-1] == 'S' ||
						input[i-1][j+1] == 'S' && input[i+1][j-1] == 'M') {
					count++
				}
			}
		}
	}

	return count
}

func CheckHorizontal(input [][]rune, row int, col int) bool {
	if col > len(input[row])-4 {
		return false
	}
	word := string(input[row][col : col+4])
	return word == "XMAS"
}

func CheckVerticalDown(input [][]rune, row int, col int) bool {
	if row > len(input)-4 {
		return false
	}

	var builder strings.Builder
	for i := 0; i < 4; i++ {
		builder.WriteRune(input[row+i][col])
	}

	word := builder.String()
	return word == "XMAS"
}

func CheckVerticalUp(input [][]rune, row int, col int) bool {
	if row < 3 {
		return false
	}

	var builder strings.Builder
	for i := 0; i < 4; i++ {
		builder.WriteRune(input[row-i][col])
	}

	word := builder.String()
	return word == "XMAS"
}

func CheckDiagonalDownRight(input [][]rune, row int, col int) bool {
	if row > len(input)-4 || col > len(input[row])-4 {
		return false
	}

	var builder strings.Builder
	for i := 0; i < 4; i++ {
		builder.WriteRune(input[row+i][col+i])
	}

	word := builder.String()
	return word == "XMAS"
}

func CheckDiagonalDownLeft(input [][]rune, row int, col int) bool {
	if row > len(input)-4 || col < 3 {
		return false
	}

	var builder strings.Builder
	for i := 0; i < 4; i++ {
		builder.WriteRune(input[row+i][col-i])
	}

	word := builder.String()
	return word == "XMAS"
}

func CheckDiagonalUpRight(input [][]rune, row int, col int) bool {
	if row < 3 || col > len(input[row])-4 {
		return false
	}

	var builder strings.Builder
	for i := 0; i < 4; i++ {
		builder.WriteRune(input[row-i][col+i])
	}

	word := builder.String()
	return word == "XMAS"
}

func CheckDiagonalUpLeft(input [][]rune, row int, col int) bool {
	if row < 3 || col < 3 {
		return false
	}

	var builder strings.Builder
	for i := 0; i < 4; i++ {
		builder.WriteRune(input[row-i][col-i])
	}

	word := builder.String()
	return word == "XMAS"
}

func CheckBackwords(input [][]rune, row int, col int) bool {
	if col < 3 {
		return false
	}

	var builder strings.Builder
	for i := 0; i < 4; i++ {
		builder.WriteRune(input[row][col-i])
	}
	word := builder.String()
	return word == "XMAS"
}
