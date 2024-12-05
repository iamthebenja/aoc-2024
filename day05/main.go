package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
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
	orderRulesMap, updates := SplitInputs(scanner)

	correctUpdates := make([][]int, 0)
updatesRange:
	for _, update := range updates {
		for idx, page := range update {
			if yVals, exists := orderRulesMap[page]; exists {
				for _, y := range yVals {
					if yIdx := slices.Index(update, y); yIdx != -1 && yIdx < idx {
						continue updatesRange
					}
				}
			}
		}

		correctUpdates = append(correctUpdates, update)
	}

	sumOfMiddlePages := 0
	for _, update := range correctUpdates {
		sumOfMiddlePages += update[int(math.Floor(float64(len(update))/2.0))]
	}

	fmt.Printf("Sum of middle pages: %d\n", sumOfMiddlePages)
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
	orderRulesMap, updates := SplitInputs(scanner)

	incorrectUpdates := make([][]int, 0)
updatesRange:
	for _, update := range updates {
		for idx, page := range update {
			if yVals, exists := orderRulesMap[page]; exists {
				for _, y := range yVals {
					if yIdx := slices.Index(update, y); yIdx != -1 && yIdx < idx {
						incorrectUpdates = append(incorrectUpdates, update)
						continue updatesRange
					}
				}
			}
		}
	}

	fmt.Println(incorrectUpdates)

	for _, incorrectUpdate := range incorrectUpdates {
		fmt.Printf("\n\nIncorrect update: %v\n", incorrectUpdate)
		for _, page := range incorrectUpdate {
			fmt.Println(orderRulesMap[page])
		}
	}

	// sumOfMiddlePages := 0
	// for _, update := range incorrectUpdates {
	// 	sumOfMiddlePages += update[int(math.Floor(float64(len(update))/2.0))]
	// }
	//
	// fmt.Printf("Sum of middle pages: %d\n", sumOfMiddlePages)
}

func SplitInputs(scanner *bufio.Scanner) (map[int][]int, [][]int) {
	createRuleMap := true
	orderRulesMap := make(map[int][]int, 0)
	updates := make([][]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			createRuleMap = false
			continue
		}

		if createRuleMap {
			rule := strings.Split(line, "|")

			ruleX, err := strconv.Atoi(rule[0])
			if err != nil {
				log.Fatal(err)
			}

			ruleY, err := strconv.Atoi(rule[1])
			if err != nil {
				log.Fatal(err)
			}

			if _, exists := orderRulesMap[ruleX]; !exists {
				orderRulesMap[ruleX] = make([]int, 0)
			}

			orderRulesMap[ruleX] = append(orderRulesMap[ruleX], ruleY)
		} else {
			update := make([]int, 0)
			for _, page := range strings.Split(line, ",") {
				pageInt, err := strconv.Atoi(page)
				if err != nil {
					log.Fatal(err)
				}

				update = append(update, pageInt)
			}

			if len(update) > 0 {
				updates = append(updates, update)
			}
		}
	}

	return orderRulesMap, updates
}
