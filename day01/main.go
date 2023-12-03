package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
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
	fmt.Printf("The sum of all values for part 1 was %v\n", sum)
	return nil
}

/********** Part 2 **********/
func getTextNumbersFromLine(line string) (int, error) {
	firstIndex, lastIndex := len(line), -1
	var first, last rune
	// Handle digits
	for _, digit := range []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'} {
		firstDigitIndex := strings.IndexRune(line, digit)
		if firstDigitIndex != -1 && firstDigitIndex < firstIndex {
			first = digit
			firstIndex = firstDigitIndex
		}
		lastDigitIndex := strings.LastIndexFunc(line, func(r rune) bool { return r == digit })
		if lastDigitIndex != -1 && lastDigitIndex > lastIndex {
			last = digit
			lastIndex = lastDigitIndex
		}
	}
	// Handle strings
	for text, char := range map[string]rune{"one": '1', "two": '2', "three": '3', "four": '4', "five": '5', "six": '6', "seven": '7', "eight": '8', "nine": '9'} {
		firstTextIndex := strings.Index(line, text)
		if firstTextIndex != -1 && firstTextIndex < firstIndex {
			first = char
			firstIndex = firstTextIndex
		}
		lastTextIndex := strings.LastIndex(line, text)
		if lastTextIndex != -1 && lastTextIndex > lastIndex {
			last = char
			lastIndex = lastTextIndex
		}
	}

	value, err := strconv.Atoi(string(first) + string(last))
	if err != nil {
		// fmt.Printf("====> Error converting. firstIndex: %v, first: %v, lastIndex: %v, last: %v\n", firstIndex, first, lastIndex, last)
		return 0, fmt.Errorf("error parsing line %q: %w", line, err)
	}
	return value, nil
}

func part2(file string) error {
	values, err := fileReader.ReadFileByLine(file, getTextNumbersFromLine)
	if err != nil {
		return err
	}

	sum := 0
	for _, value := range values {
		sum += value
	}
	fmt.Printf("The sum of all values for part 2 was %v\n", sum)
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
