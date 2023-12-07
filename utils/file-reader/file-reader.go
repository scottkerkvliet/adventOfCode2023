package fileReader

import (
	"bufio"
	"os"
)

/***** Types *****/

type LineReader[T any] func(string) (T, error)

/***** Methods *****/

func GetFileScanner(path string) (*bufio.Scanner, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return bufio.NewScanner(file), nil
}

func ReadFileByLine[K any](path string, lr LineReader[K]) ([]K, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	var values []K
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, err := lr(scanner.Text())
		if err != nil {
			return nil, err
		}
		values = append(values, value)
	}

	return values, nil
}
