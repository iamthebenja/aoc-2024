package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
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

	reports := GetReports(*scanner)

	safeReports := 0
	for _, report := range reports {
		isSafe := CheckReport1(report)
		if isSafe {
			safeReports++
		}
	}

	fmt.Println(safeReports)
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

	reports := GetReports(*scanner)

	safeReports := 0
	for _, report := range reports {
		isSafe := CheckReport2(report)
		if isSafe {
			safeReports++
		}
	}

	fmt.Println(safeReports)
}

func GetReports(scanner bufio.Scanner) [][]int {
	var reports [][]int

	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Fields(line)

		if len(nums) == 0 {
			continue
		}

		var report []int
		for _, num := range nums {
			convertedNum, err := strconv.Atoi(num)
			if err != nil {
				log.Fatal(err)
			}

			report = append(report, convertedNum)
		}

		reports = append(reports, report)

	}

	return reports
}

func CheckReport1(report []int) bool {
	isDecreasing := report[0] > report[1]

	for idx := range report {
		if idx == 0 {
			continue
		}
		delta := math.Abs(float64(report[idx]) - float64(report[idx-1]))

		if delta == 0 || delta > 3 {
			return false
		}

		if isDecreasing != (report[idx] < report[idx-1]) {
			return false
		}
	}

	return true
}

func CheckReport2(report []int) bool {
	if CheckReport1(report) {
		return true
	}

	for idx := range report {
		newReport := make([]int, len(report)-1)
		copy(newReport[:idx], report[:idx])
		if idx < len(report)-1 {
			copy(newReport[idx:], report[idx+1:])
		}
		if CheckReport1(newReport) {
			return true
		}
	}

	return false
}
