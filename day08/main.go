package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"

	fileReader "github.com/scottkerkvliet/advent-of-code-2023/utils/file-reader"
)

type Node struct {
	name, left, right string
}

func followDirectionsToZZZ(directions string, nodeMap map[string]*Node) (int, error) {
	const start, end = "AAA", "ZZZ"
	currentNode := nodeMap[start]
	var totalSteps int
	for currentNode.name != end {
		for _, direction := range directions {
			switch direction {
			case 'L':
				currentNode = nodeMap[currentNode.left]
			case 'R':
				currentNode = nodeMap[currentNode.right]
			default:
				return 0, fmt.Errorf("Got invalid direction: %q", string(direction))
			}
			totalSteps++
			if currentNode.name == end {
				break
			}
		}
	}
	return totalSteps, nil
}

func readNodeLine(line string) (*Node, error) {
	if len(line) != 16 {
		return nil, fmt.Errorf("Expected line in format \"xxx = (yyy, zzz)\", got %q", line)
	}

	return &Node{name: line[0:3], left: line[7:10], right: line[12:15]}, nil
}

func readMap(scanner *bufio.Scanner) (string, map[string]*Node, error) {
	if !scanner.Scan() {
		return "", nil, fmt.Errorf("File was empty")
	}
	directions := scanner.Text()

	scanner.Scan() // Skip blank line
	nodeMap := make(map[string]*Node)
	for scanner.Scan() {
		node, err := readNodeLine(scanner.Text())
		if err != nil {
			return "", nil, err
		}
		nodeMap[node.name] = node
	}
	return directions, nodeMap, nil
}

func part1(file string) error {
	scanner, err := fileReader.GetFileScanner(file)
	if err != nil {
		return err
	}
	directions, nodeMap, err := readMap(scanner)
	if err != nil {
		return err
	}

	totalSteps, err := followDirectionsToZZZ(directions, nodeMap)
	if err != nil {
		return err
	}

	fmt.Printf("In part 1, you will reach ZZZ in %v steps\n", totalSteps)
	return nil
}

func part2(file string) error {
	fmt.Println("Part 2 is not implemented.")
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
