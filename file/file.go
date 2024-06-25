package file

import (
	"bufio"
	"os"
	"strconv"
)

func ReadFileLines(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	numbers := make([]int, 0, 1000)
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			continue
		}
		numbers = append(numbers, num)
	}

	if err := scanner.Err(); err != nil {
		return nil
	}

	return numbers
}

func FindValueIndex(numbers []int, target int) int {
	for index, number := range numbers {
		if number == target {
			return index
		}
	}
	return -1
}
