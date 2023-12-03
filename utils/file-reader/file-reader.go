package fileReader

import (
	"bufio"
	"os"
)

/***** Types *****/
type LineReader[T any] func(string) T

/***** Methods *****/
func ReadFileByLine[K any](path string, lr LineReader[K]) ([]K, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	var values []K
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		values = append(values, lr(scanner.Text()))
	}

	return values, nil
}
