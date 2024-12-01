package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	puzzle := flag.Int("puzzle", 1, "Which puzzle to run")
	test := flag.Bool("test", false, "Use test inputs")
	flag.Parse()

	if *puzzle == 1 {
		Run1(*test)
	} else {
		Run2(*test)
	}
}

func Run1(useTest bool) {
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

	firstOfPair, secondOfPair := ConvertPairsToArrays(*scanner)

	if len(firstOfPair) != len(secondOfPair) {
		log.Fatal("Pair lengths don't match!")
	}

	sort.Ints(firstOfPair)
	sort.Ints(secondOfPair)

	var diffs []float64
	for idx, value := range firstOfPair {
		diffs = append(diffs, math.Abs(float64(value)-float64(secondOfPair[idx])))
	}

	fmt.Println()
	fmt.Println(diffs)

	sum := 0.0
	for _, value := range diffs {
		sum += value
	}

	fmt.Println(sum)
}

func Run2(useTest bool) {
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

	firstOfPair, secondOfPair := ConvertPairsToArrays(*scanner)

	if len(firstOfPair) != len(secondOfPair) {
		log.Fatal("Pair lengths don't match!")
	}

	numOfTimes := make(map[int]int)
	for _, value := range secondOfPair {
		numOfTimes[value]++
	}

	var similarityScores []int
	for _, value := range firstOfPair {
		similarityScores = append(similarityScores, value*numOfTimes[value])
	}

	fmt.Println(similarityScores)

	sum := 0
	for _, value := range similarityScores {
		sum += value
	}

	fmt.Println(sum)

}

func ConvertPairsToArrays(scanner bufio.Scanner) ([]int, []int) {
	var firstOfPair []int
	var secondOfPair []int

	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Fields(line)

		if len(nums) == 0 {
			continue
		}

		num1, err := strconv.Atoi(nums[0])
		if err != nil {
			log.Fatal(err)
		}

		firstOfPair = append(firstOfPair, num1)

		num2, err := strconv.Atoi(nums[1])
		if err != nil {
			log.Fatal(err)
		}

		secondOfPair = append(secondOfPair, num2)
	}

	return firstOfPair, secondOfPair
}
