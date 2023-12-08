package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	fileReader "github.com/scottkerkvliet/advent-of-code-2023/utils/file-reader"
)

type Race struct {
	duration, distance int
}

func (r *Race) delayWins(d int) bool {
	return (d * (r.duration - d)) > r.distance
}

func winningStrategies(r *Race) (strategies int) {
	for i := 1; i < r.duration; i++ {
		if r.delayWins(i) {
			strategies++
		}
	}
	return
}

func winningStrategiesBinarySearch(r *Race) int {
	middle := r.duration / 2
	if !r.delayWins(middle) {
		return 0
	}
	minDelay := sort.Search(middle, r.delayWins)
	strategiesBelowMiddle := (middle - minDelay) * 2
	strategiesAtMiddle := 1
	if (r.duration % 2) == 1 {
		strategiesAtMiddle = 2
	}

	return strategiesBelowMiddle + strategiesAtMiddle
}

func readRaceFile(scanner *bufio.Scanner, ignoreSpaces bool) ([]*Race, error) {
	if !scanner.Scan() {
		return nil, fmt.Errorf("Could not get first line in race file")
	}
	timeLine := scanner.Text()
	if !scanner.Scan() {
		return nil, fmt.Errorf("Could not get second line in race file")
	}
	distanceLine := scanner.Text()

	var times, distances []int
	if ignoreSpaces {
		timeLine = strings.ReplaceAll(timeLine, " ", "")
		distanceLine = strings.ReplaceAll(distanceLine, " ", "")
	}
	timeParts := strings.Split(strings.TrimPrefix(timeLine, "Time:"), " ")
	for _, timePart := range timeParts {
		if len(timePart) != 0 {
			value, err := strconv.Atoi(timePart)
			if err != nil {
				return nil, fmt.Errorf("Error parsing time: %w", err)
			}
			times = append(times, value)
		}
	}

	distanceParts := strings.Split(strings.TrimPrefix(distanceLine, "Distance:"), " ")
	for _, distancePart := range distanceParts {
		if len(distancePart) != 0 {
			value, err := strconv.Atoi(distancePart)
			if err != nil {
				return nil, fmt.Errorf("Error parsing distance: %w", err)
			}
			distances = append(distances, value)
		}
	}

	if len(times) != len(distances) {
		return nil, fmt.Errorf("Got different quantity of values for time (%v) and distance (%v)", len(times), len(distances))
	}

	var races []*Race
	for i := range times {
		races = append(races, &Race{duration: times[i], distance: distances[i]})
	}

	return races, nil
}

func part1(file string) error {
	scanner, err := fileReader.GetFileScanner(file)
	if err != nil {
		return err
	}
	races, err := readRaceFile(scanner, false)
	if err != nil {
		return err
	}

	product := 1
	for _, race := range races {
		product = product * winningStrategiesBinarySearch(race)
	}

	fmt.Printf("The product of winning strategies in part 1 is %v\n", product)
	return nil
}

func part2(file string) error {
	scanner, err := fileReader.GetFileScanner(file)
	if err != nil {
		return err
	}
	races, err := readRaceFile(scanner, true)
	if err != nil {
		return err
	}
	if len(races) != 1 {
		return fmt.Errorf("Expected one race, got %v", len(races))
	}

	strategies := winningStrategiesBinarySearch(races[0])

	fmt.Printf("The number of winning strategies in part 2 is %v\n", strategies)
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
