package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"unicode"

	fileReader "github.com/scottkerkvliet/advent-of-code-2023/utils/file-reader"
)

/********** Part 1 **********/
func getNumberFromLine(line string) (int, error) {
	var first, last rune
	for _, char := range line {
		if unicode.IsDigit(char) {
			if first == 0 {
				first = char
			}
			last = char
		}
	}

	value, err := strconv.Atoi(string(first) + string(last))
	if err != nil {
		return 0, err
	}
	return value, nil
}

func part1(file string) error {
	values, err := fileReader.ReadFileByLine(file, getNumberFromLine)
	if err != nil {
		return err
	}

	sum := 0
	for _, value := range values {
		sum += value
	}
	fmt.Printf("The sum of all values was %v\n", sum)
	return nil
}

/********** Part 2 **********/
func part2(file string) error {
	fmt.Println("Part 2 is not implemented.")
	return nil
}

/********** main **********/
func main() {
	inputFile := flag.String("file", "input.txt", "the input file to execute")
	part1Flag := flag.Bool("1", false, "if provided, execute puzzle 1")
	part2Flag := flag.Bool("2", false, "if provided, execute puzzle 2")
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
