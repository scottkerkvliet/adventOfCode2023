package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	fileReader "github.com/scottkerkvliet/advent-of-code-2023/utils/file-reader"
)

type Draw struct {
	green, blue, red int
}

type Game struct {
	id    int
	draws []Draw
}

func (g *Game) MaxDraw() Draw {
	maxDraw := Draw{}
	for _, draw := range g.draws {
		maxDraw.green = max(maxDraw.green, draw.green)
		maxDraw.red = max(maxDraw.red, draw.red)
		maxDraw.blue = max(maxDraw.blue, draw.blue)
	}
	return maxDraw
}

func getIntFromString(s string) (int, error) {
	return strconv.Atoi(strings.TrimSpace(s))
}

func getDrawFromString(s string) (Draw, error) {
	greenSuffix := " green"
	redSuffix := " red"
	blueSuffix := " blue"

	draw := Draw{}
	for _, cubePart := range strings.Split(s, ",") {
		switch {
		case strings.HasSuffix(cubePart, greenSuffix):
			numCubes, err := getIntFromString(strings.TrimSuffix(cubePart, greenSuffix))
			if err != nil {
				return draw, err
			}
			draw.green = numCubes
		case strings.HasSuffix(cubePart, redSuffix):
			numCubes, err := getIntFromString(strings.TrimSuffix(cubePart, redSuffix))
			if err != nil {
				return draw, err
			}
			draw.red = numCubes
		case strings.HasSuffix(cubePart, blueSuffix):
			numCubes, err := getIntFromString(strings.TrimSuffix(cubePart, blueSuffix))
			if err != nil {
				return draw, err
			}
			draw.blue = numCubes
		default:
			return draw, fmt.Errorf("Draw string had no expected suffix %q", s)
		}
	}
	return draw, nil
}

func getGameFromLine(line string) (*Game, error) {
	gamePrefix := "Game "

	lineParts := strings.Split(line, ":")
	if len(lineParts) != 2 {
		return nil, fmt.Errorf("Got %v colons in line, wanted 1", len(lineParts)-1)
	}
	idPart := lineParts[0]
	drawPart := lineParts[1]

	if !strings.HasPrefix(idPart, gamePrefix) {
		return nil, fmt.Errorf("Line does not begin with %q", gamePrefix)
	}
	id, err := getIntFromString(strings.TrimPrefix(idPart, gamePrefix))
	if err != nil {
		return nil, err
	}
	game := &Game{id, []Draw{}}

	drawParts := strings.Split(drawPart, ";")
	if len(drawParts) == 0 {
		return nil, fmt.Errorf("Line has no draws %q", line)
	}
	for _, drawString := range drawParts {
		draw, err := getDrawFromString(drawString)
		if err != nil {
			return nil, err
		}
		game.draws = append(game.draws, draw)
	}

	return game, nil
}

func part1(file string) error {
	games, err := fileReader.ReadFileByLine(file, getGameFromLine)
	if err != nil {
		return err
	}

	var gameIdSum int
	for _, game := range games {
		minCubes := game.MaxDraw()
		if minCubes.red <= 12 && minCubes.green <= 13 && minCubes.blue <= 14 {
			gameIdSum += game.id
		}
	}

	fmt.Printf("The sum of all valid games in part 1 is %v\n", gameIdSum)
	return nil
}

func part2(file string) error {
	games, err := fileReader.ReadFileByLine(file, getGameFromLine)
	if err != nil {
		return err
	}

	var powerSum int
	for _, game := range games {
		minCubes := game.MaxDraw()
		power := minCubes.blue * minCubes.green * minCubes.red
		powerSum += power
	}

	fmt.Printf("The sum of all powers in part 2 is %v\n", powerSum)
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
