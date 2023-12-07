package main

import (
	"flag"
	"fmt"
	"log"
	"slices"
	"strconv"
	"unicode"

	fileReader "github.com/scottkerkvliet/advent-of-code-2023/utils/file-reader"
)

type Number struct {
	value int
}

func BuildMatrix(file string) ([][]any, error) {
	scanner, err := fileReader.GetFileScanner(file)
	if err != nil {
		return nil, err
	}
	var matrix [][]any
	for scanner.Scan() {
		line := make([]any, len(scanner.Text()))
		var currentNumber string
		for i, char := range scanner.Text() {
			if unicode.IsDigit(char) {
				currentNumber += string(char)
				continue
			}
			if len(currentNumber) > 0 {
				value, err := strconv.Atoi(currentNumber)
				if err != nil {
					return nil, err
				}
				for j := i - len(currentNumber); j < i; j++ {
					line[j] = &value
				}
				currentNumber = ""
			}
			if char != '.' {
				line[i] = char
			}
		}
		if len(currentNumber) > 0 {
			value, err := strconv.Atoi(currentNumber)
			if err != nil {
				return nil, err
			}
			for i := len(line) - len(currentNumber); i < len(line); i++ {
				line[i] = &value
			}
		}
		matrix = append(matrix, line)
	}

	return matrix, nil
}

func getNumbersAdjacentToSymbols(matrix [][]any) []*int {
	var uniqueNums []*int
	for i, row := range matrix {
		for j, cell := range row {
			if _, ok := cell.(rune); ok {
				// Look for all numbers around symbol
				for m := max(0, i-1); m < min(len(matrix), i+2); m++ {
					for n := max(0, j-1); n < min(len(matrix[m]), j+2); n++ {
						if value, ok := matrix[m][n].(*int); ok && !slices.Contains(uniqueNums, value) {
							uniqueNums = append(uniqueNums, value)
						}
					}
				}
			}
		}
	}

	return uniqueNums
}

func getSumOfGearRatios(matrix [][]any) int {
	var sum int
	for i, row := range matrix {
		for j, cell := range row {
			if char, ok := cell.(rune); ok && char == '*' {
				// Look for exactly 2 numbers beside the gear
				var surroundingNums []*int
				for m := max(0, i-1); m < min(len(matrix), i+2); m++ {
					for n := max(0, j-1); n < min(len(matrix[m]), j+2); n++ {
						if value, ok := matrix[m][n].(*int); ok && !slices.Contains(surroundingNums, value) {
							surroundingNums = append(surroundingNums, value)
						}
					}
				}
				if len(surroundingNums) == 2 {
					sum += (*surroundingNums[0] * *surroundingNums[1])
				}
			}
		}
	}
	return sum
}

func checkMatrix(matrix [][]any, matched []*int) {
	var unmatched []*int
	for _, row := range matrix {
		for _, cell := range row {
			if value, ok := cell.(*int); ok && !slices.Contains(matched, value) && !slices.Contains(unmatched, value) {
				unmatched = append(unmatched, value)
			}
		}
	}

	fmt.Println("Matched values")
	for _, val := range matched {
		fmt.Printf("%v\n", *val)
	}
	fmt.Println("\nUnmatched values")
	for _, val := range unmatched {
		fmt.Printf("%v\n", *val)
	}
}

func part1(file string) error {
	matrix, err := BuildMatrix(file)
	if err != nil {
		return err
	}

	nums := getNumbersAdjacentToSymbols(matrix)
	var sum int
	for _, num := range nums {
		sum += *num
	}

	// checkMatrix(matrix, nums)
	fmt.Printf("The sum of numbers adjacent to symbols (part 1) is %v\n", sum)

	return nil
}

func part2(file string) error {
	matrix, err := BuildMatrix(file)
	if err != nil {
		return err
	}

	sum := getSumOfGearRatios(matrix)

	fmt.Printf("The sum of the gear ratios (part 1) is %v\n", sum)
	return nil
}

func main() {
	inputFile := flag.String("file", "input.txt", "the input file to execute")
	part1Flag := flag.Bool("1", false, "whether to execute puzzle 1")
	part2Flag := flag.Bool("2", false, "whether to execute puzzle 2")
	flag.Parse()
	if !(*part1Flag || *part2Flag) {
		fmt.Println("Nothing to do, specify a puzzle to solve")
		return
	}

	if *part1Flag {
		if err := part1(*inputFile); err != nil {
			log.Fatal(err)
		}
	}
	if *part2Flag {
		if err := part2(*inputFile); err != nil {
			log.Fatal(err)
		}
	}

	return
}
