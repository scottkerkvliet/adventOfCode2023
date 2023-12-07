package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	fileReader "github.com/scottkerkvliet/advent-of-code-2023/utils/file-reader"
)

/********** Types **********/

type Seed int
type Soil int
type Fertilizer int
type Water int
type Light int
type Temperature int
type Humidity int
type Location int

type SeedValues struct {
	seed  Seed
	soil  Soil
	fert  Fertilizer
	water Water
	light Light
	temp  Temperature
	humid Humidity
	loc   Location
}

func (sv *SeedValues) String() string {
	return fmt.Sprintf("Seed %v, soil %v, fertilizer %v, water %v, light %v, temperature %v, humidity %v, location %v", sv.seed, sv.soil, sv.fert, sv.water, sv.light, sv.temp, sv.humid, sv.loc)
}

type Mapper[S ~int, D ~int] struct {
	source      S
	destination D
	length      int
}

func (m *Mapper[S, D]) GetMapping(source S) (D, bool) {
	diff := int(source) - int(m.source)
	if diff < 0 || int(diff) >= m.length {
		return 0, false
	}
	return m.destination + D(diff), true
}

func GetDestinationFromMappers[S ~int, D ~int](source S, mappers []*Mapper[S, D]) D {
	for _, mapper := range mappers {
		if d, ok := mapper.GetMapping(source); ok {
			return d
		}
	}
	return D(source)
}

type Almanac struct {
	soilMaps  []*Mapper[Seed, Soil]
	fertMaps  []*Mapper[Soil, Fertilizer]
	waterMaps []*Mapper[Fertilizer, Water]
	lightMaps []*Mapper[Water, Light]
	tempMaps  []*Mapper[Light, Temperature]
	humidMaps []*Mapper[Temperature, Humidity]
	locMaps   []*Mapper[Humidity, Location]
}

func (a *Almanac) FindSeedValues(seed Seed) *SeedValues {
	sv := &SeedValues{seed: seed}
	sv.soil = GetDestinationFromMappers(sv.seed, a.soilMaps)
	sv.fert = GetDestinationFromMappers(sv.soil, a.fertMaps)
	sv.water = GetDestinationFromMappers(sv.fert, a.waterMaps)
	sv.light = GetDestinationFromMappers(sv.water, a.lightMaps)
	sv.temp = GetDestinationFromMappers(sv.light, a.tempMaps)
	sv.humid = GetDestinationFromMappers(sv.temp, a.humidMaps)
	sv.loc = GetDestinationFromMappers(sv.humid, a.locMaps)
	return sv
}

/********** File Functions **********/

const (
	seedsPrefix        = "seeds: "
	seedToSoil         = "seed-to-soil map:"
	soilToFertilizer   = "soil-to-fertilizer map:"
	fertilizerToWater  = "fertilizer-to-water map:"
	waterToLight       = "water-to-light map:"
	lightToTemp        = "light-to-temperature map:"
	tempToHumidity     = "temperature-to-humidity map:"
	humidityToLocation = "humidity-to-location map:"
)

func readSeedsLine(line string) ([]Seed, error) {
	seedParts := strings.Split(strings.TrimPrefix(line, seedsPrefix), " ")
	var seeds []Seed
	for _, seedPart := range seedParts {
		value, err := strconv.Atoi(seedPart)
		if err != nil {
			return nil, fmt.Errorf("Error parsing seed: %w", err)
		}
		seeds = append(seeds, Seed(value))
	}
	return seeds, nil
}

func readMapperLine[S ~int, D ~int](line string) (*Mapper[S, D], error) {
	parts := strings.Split(line, " ")
	if len(parts) != 3 {
		return nil, fmt.Errorf("Expected 3 numbers in line %q, got %v", line, len(parts))
	}

	dest, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, fmt.Errorf("Error parsing destination: %w", err)
	}

	source, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, fmt.Errorf("Error parsing source: %w", err)
	}

	len, err := strconv.Atoi(parts[2])
	if err != nil {
		return nil, fmt.Errorf("Error parsing length: %w", err)
	}

	return &Mapper[S, D]{source: S(source), destination: D(dest), length: len}, nil
}

func readMapperLines[S ~int, D ~int](scanner *bufio.Scanner) ([]*Mapper[S, D], error) {
	var mappers []*Mapper[S, D]
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			return mappers, nil
		}
		mapper, err := readMapperLine[S, D](scanner.Text())
		if err != nil {
			return nil, err
		}
		mappers = append(mappers, mapper)
	}
	return mappers, nil
}

func readAlmanacFile(scanner *bufio.Scanner) ([]Seed, *Almanac, error) {
	if !scanner.Scan() {
		return nil, nil, fmt.Errorf("Empty almanac file")
	}
	seeds, err := readSeedsLine(scanner.Text())
	if err != nil {
		return nil, nil, err
	}

	scanner.Scan() // clear empty line
	almanac := &Almanac{}

	for scanner.Scan() {
		switch scanner.Text() {
		case seedToSoil:
			soilMaps, err := readMapperLines[Seed, Soil](scanner)
			if err != nil {
				return nil, nil, err
			}
			almanac.soilMaps = soilMaps
		case soilToFertilizer:
			fertMaps, err := readMapperLines[Soil, Fertilizer](scanner)
			if err != nil {
				return nil, nil, err
			}
			almanac.fertMaps = fertMaps
		case fertilizerToWater:
			waterMaps, err := readMapperLines[Fertilizer, Water](scanner)
			if err != nil {
				return nil, nil, err
			}
			almanac.waterMaps = waterMaps
		case waterToLight:
			lightMaps, err := readMapperLines[Water, Light](scanner)
			if err != nil {
				return nil, nil, err
			}
			almanac.lightMaps = lightMaps
		case lightToTemp:
			tempMaps, err := readMapperLines[Light, Temperature](scanner)
			if err != nil {
				return nil, nil, err
			}
			almanac.tempMaps = tempMaps
		case tempToHumidity:
			humidMaps, err := readMapperLines[Temperature, Humidity](scanner)
			if err != nil {
				return nil, nil, err
			}
			almanac.humidMaps = humidMaps
		case humidityToLocation:
			locMaps, err := readMapperLines[Humidity, Location](scanner)
			if err != nil {
				return nil, nil, err
			}
			almanac.locMaps = locMaps
		}
	}

	return seeds, almanac, nil
}

/********** Main Functions **********/

func part1(file string) error {
	scanner, err := fileReader.GetFileScanner(file)
	if err != nil {
		return err
	}
	seeds, almanac, err := readAlmanacFile(scanner)
	if err != nil {
		return err
	}

	minLocation := math.MaxInt
	for _, seed := range seeds {
		seedValues := almanac.FindSeedValues(seed)
		// fmt.Println(seedValues.String())
		minLocation = min(minLocation, int(seedValues.loc))
	}

	fmt.Printf("The minimum seed location in part 1 is %v\n", minLocation)
	return nil
}

// Brute force got 12634632
func part2(file string) error {
	scanner, err := fileReader.GetFileScanner(file)
	if err != nil {
		return err
	}
	seedNums, almanac, err := readAlmanacFile(scanner)
	if err != nil {
		return err
	}

	if len(seedNums)%2 != 0 {
		return fmt.Errorf("Expected even number of seeds, got %v", len(seedNums))
	}
	minLocation := math.MaxInt
	for i := 0; i < len(seedNums); i += 2 {
		// Iterate, where index i is the start number and i+1 is the number of cycles
		for seed := seedNums[i]; seed < seedNums[i]+seedNums[i+1]; seed++ {
			seedValues := almanac.FindSeedValues(seed)
			// fmt.Println(seedValues.String())
			minLocation = min(minLocation, int(seedValues.loc))
		}
	}

	fmt.Printf("The minimum seed location in part 2 is %v\n", minLocation)
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
