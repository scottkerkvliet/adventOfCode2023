package main

import (
	"flag"
	"fmt"
	"log"
	"sort"
	"strconv"
	"unicode"

	fileReader "github.com/scottkerkvliet/advent-of-code-2023/utils/file-reader"
)

type HandType int

const (
	HighCard HandType = iota
	Pair
	TwoPair
	Triple
	FullHouse
	Quadruple
	Quintuple
)

type Hand struct {
	cards    [5]int
	handType HandType
	bid      int
}

func (h *Hand) IsBetterThan(o *Hand) bool {
	if h.handType != o.handType {
		return h.handType > o.handType
	}
	for i := range h.cards {
		if h.cards[i] != o.cards[i] {
			return h.cards[i] > o.cards[i]
		}
	}
	return false
}

func getHandReader(jokers bool) func(line string) (*Hand, error) {
	return func(line string) (*Hand, error) {
		if len(line) < 7 {
			return nil, fmt.Errorf("Line is too short for a hand %q", line)
		}
		cardString := line[0:5]
		bidString := line[6:]

		bid, err := strconv.Atoi(bidString)
		if err != nil {
			return nil, fmt.Errorf("Error parsing bid: %w", err)
		}

		var cards [5]int
		cardMap := make(map[int]int)
		for i, char := range cardString {
			switch {
			case unicode.IsDigit(char):
				cards[i] = int(char - '0')
			case char == 'A':
				cards[i] = 14
			case char == 'K':
				cards[i] = 13
			case char == 'Q':
				cards[i] = 12
			case char == 'J' && jokers:
				cards[i] = -1
			case char == 'J':
				cards[i] = 11
			case char == 'T':
				cards[i] = 10
			default:
				return nil, fmt.Errorf("Got unexpected rune in cards: %q", string(char))
			}
			cardMap[cards[i]]++
		}

		highestCount, secondHighestCount := 0, 0
		for card, count := range cardMap {
			if card == -1 {
				continue
			}
			if count >= highestCount {
				secondHighestCount = highestCount
				highestCount = count
			} else if count > secondHighestCount {
				secondHighestCount = count
			}
		}
		highestCount += cardMap[-1]

		hand := &Hand{cards: cards, bid: bid}
		switch {
		case highestCount == 5:
			hand.handType = Quintuple
		case highestCount == 4:
			hand.handType = Quadruple
		case highestCount == 3 && secondHighestCount == 2:
			hand.handType = FullHouse
		case highestCount == 3:
			hand.handType = Triple
		case highestCount == 2 && secondHighestCount == 2:
			hand.handType = TwoPair
		case highestCount == 2:
			hand.handType = Pair
		default:
			hand.handType = HighCard
		}
		return hand, nil
	}
}

func printCards(cards [5]int) (cardString string) {
	for _, card := range cards {
		switch {
		case card <= 9:
			cardString += strconv.Itoa(card)
		case card == 10:
			cardString += "T"
		case card == 11:
			cardString += "J"
		case card == 12:
			cardString += "Q"
		case card == 13:
			cardString += "K"
		case card == 14:
			cardString += "A"
		}
	}
	return
}

func getTotalWinnings(hands []*Hand) int {
	sort.Slice(hands, func(i, j int) bool {
		return !hands[i].IsBetterThan(hands[j])
	})

	var totalWinnings int
	for i, hand := range hands {
		rank := i + 1
		winnings := hand.bid * rank
		totalWinnings += winnings
		// fmt.Printf("Hand %v type %v bid %4d rank %4d winnings %v\n", printCards(hand.cards), hand.handType, hand.bid, rank, winnings)
	}
	return totalWinnings
}

func part1(file string) error {
	hands, err := fileReader.ReadFileByLine(file, getHandReader(false))
	if err != nil {
		return err
	}

	totalWinnings := getTotalWinnings(hands)

	fmt.Printf("The total winnings from part 1 are %v\n", totalWinnings)
	return nil
}

func part2(file string) error {
	hands, err := fileReader.ReadFileByLine(file, getHandReader(true))
	if err != nil {
		return err
	}

	totalWinnings := getTotalWinnings(hands)

	fmt.Printf("The total winnings from part 2 are %v\n", totalWinnings)
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
