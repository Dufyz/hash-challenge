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
	numbers := make([]int, 0, 1000000)

	const maxCapacity = 64 * 1024 * 20
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

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
