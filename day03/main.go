package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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
	input := GetAllLines(*scanner)

	fmt.Println(SumMultiples(input))
}

func Puzzle2(useTest bool) {
	filePath := "./inputs/01.txt"
	if useTest {
		filePath = "./inputs/02-test.txt"
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := GetAllLines(*scanner)

	dos := GetDoMultiples(input)
	fmt.Println(dos)
}

func GetAllLines(scanner bufio.Scanner) []string {
	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func SumMultiples(lines []string) int {
	validMultipliers := regexp.MustCompile(`(?:mul\()(\d+)(?:[,])(\d+)(?:\))`)
	sum := 0

	for _, line := range lines {
		multipliers := validMultipliers.FindAllStringSubmatch(line, -1)

		for _, val := range multipliers {
			fmt.Println(val[0])
			val1, err := strconv.Atoi(val[1])
			if err != nil {
				log.Fatal(err)
			}

			val2, err := strconv.Atoi(val[2])
			if err != nil {
				log.Fatal(err)
			}

			sum += val1 * val2
		}
	}

	return sum
}

func GetDoMultiples(lines []string) int {
	validMultipliers := regexp.MustCompile(`(?:mul\()(\d+)(?:[,])(\d+)(?:\))|do\(\)|don't\(\)`)
	do := true
	sum := 0

	for _, line := range lines {
		multipliers := validMultipliers.FindAllStringSubmatch(line, -1)

		for _, val := range multipliers {
			if val[0] == "do()" {
				do = true
				continue
			}

			if val[0] == "don't()" {
				do = false
				continue
			}

			if do {
				val1, err := strconv.Atoi(val[1])
				if err != nil {
					log.Fatal(err)
				}

				val2, err := strconv.Atoi(val[2])
				if err != nil {
					log.Fatal(err)
				}

				sum += val1 * val2
			}
		}
	}

	// fmt.Println(matches)
	return sum
}
