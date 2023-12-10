package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	fileReader "github.com/scottkerkvliet/advent-of-code-2023/utils/file-reader"
)

type History []int

func (h History) PredictNextValue() int {
	if len(h) == 0 {
		return 0
	}
	allZero := h[0] == 0
	var newHistory History
	for i := 1; i < len(h); i++ {
		allZero = allZero && h[i] == 0
		newHistory = append(newHistory, h[i]-h[i-1])
	}
	if allZero {
		return 0
	}
	return newHistory.PredictNextValue() + h[len(h)-1]
}

func (h History) PredictPreviousValue() int {
	if len(h) == 0 {
		return 0
	}
	allZero := h[0] == 0
	var newHistory History
	for i := 1; i < len(h); i++ {
		allZero = allZero && h[i] == 0
		newHistory = append(newHistory, h[i]-h[i-1])
	}
	if allZero {
		return 0
	}
	return h[0] - newHistory.PredictPreviousValue()
}

func readHistoryLine(line string) (History, error) {
	valueParts := strings.Split(line, " ")
	var h History
	for _, valuePart := range valueParts {
		value, err := strconv.Atoi(valuePart)
		if err != nil {
			return nil, err
		}
		h = append(h, value)
	}
	return h, nil
}

func part1(file string) error {
	histories, err := fileReader.ReadFileByLine(file, readHistoryLine)
	if err != nil {
		return err
	}

	var sum int
	for _, h := range histories {
		sum += h.PredictNextValue()
	}

	fmt.Printf("The sum of the next values is %v\n", sum)
	return nil
}

func part2(file string) error {
	histories, err := fileReader.ReadFileByLine(file, readHistoryLine)
	if err != nil {
		return err
	}

	var sum int
	for _, h := range histories {
		sum += h.PredictPreviousValue()
	}

	fmt.Printf("The sum of the previous values is %v\n", sum)
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
