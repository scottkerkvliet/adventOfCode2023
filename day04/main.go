package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"slices"
	"strconv"
	"strings"

	fileReader "github.com/scottkerkvliet/advent-of-code-2023/utils/file-reader"
)

type Scratchcard struct {
	id               int
	winningNumbers   []int
	scratchedNumbers []int
}

func (s *Scratchcard) GetNumberOfMatches() int {
	var matches int
	for _, scratchedNum := range s.scratchedNumbers {
		if slices.Contains(s.winningNumbers, scratchedNum) {
			matches++
		}
	}
	return matches
}

func parseSpaceSeparatedNumbers(s string) ([]int, error) {
	var nums []int
	numberParts := strings.Split(strings.TrimSpace(s), " ")
	for _, numberPart := range numberParts {
		if len(numberPart) == 0 {
			continue
		}
		value, err := strconv.Atoi(numberPart)
		if err != nil {
			return nil, fmt.Errorf("Error parsing number %q in string %q: %w", numberPart, s, err)
		}
		nums = append(nums, value)
	}
	return nums, nil
}

func readCard(line string) (*Scratchcard, error) {
	lineParts := strings.Split(line, ":")
	if len(lineParts) != 2 {
		return nil, fmt.Errorf("Expected line to have 1 colon, got %v", len(lineParts)-1)
	}
	cardPart := lineParts[0]
	numbersPart := lineParts[1]

	cardIdPart := strings.TrimPrefix(cardPart, "Card ")
	cardId, err := strconv.Atoi(strings.TrimSpace(cardIdPart))
	if err != nil {
		return nil, fmt.Errorf("Could not parse card id: %w", err)
	}

	sectionParts := strings.Split(numbersPart, " | ")
	if len(sectionParts) != 2 {
		return nil, fmt.Errorf("Expected line to have 1 bar, got %v", len(sectionParts)-1)
	}
	winningNumbers, err := parseSpaceSeparatedNumbers(sectionParts[0])
	if err != nil {
		return nil, err
	}
	scratchedNumbers, err := parseSpaceSeparatedNumbers(sectionParts[1])
	if err != nil {
		return nil, err
	}

	return &Scratchcard{id: cardId, winningNumbers: winningNumbers, scratchedNumbers: scratchedNumbers}, nil
}

func part1(file string) error {
	cards, err := fileReader.ReadFileByLine(file, readCard)
	if err != nil {
		return err
	}

	var sum int
	for _, card := range cards {
		points := int(math.Floor(math.Pow(2, float64(card.GetNumberOfMatches()-1))))
		sum += points
	}

	fmt.Printf("The scratchcards have a total of %v points\n", sum)
	return nil
}

func part2(file string) error {
	cards, err := fileReader.ReadFileByLine(file, readCard)
	if err != nil {
		return err
	}

	cardCopies := make([]int64, len(cards))
	for i := range cardCopies {
		cardCopies[i] = 1
	}

	var totalCopies int64
	for i, card := range cards {
		matches := card.GetNumberOfMatches()
		copies := cardCopies[i]
		for j := i + 1; j < min(i+1+matches, len(cardCopies)); j++ {
			cardCopies[j] += copies
		}
		totalCopies += copies
	}

	fmt.Printf("There are a total of %v cards in part 2\n", totalCopies)
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
