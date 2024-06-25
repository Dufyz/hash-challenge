package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Dufyz/hash-challenge/file"
	"github.com/Dufyz/hash-challenge/hash"
)

func main() {
	start := time.Now()
	hashTable := hash.NewHashTable()

	linesFromFile := file.ReadFileLines("RandomNumbers.txt")
	if linesFromFile == nil {
		fmt.Println("Error reading file")
		return
	}

	for _, num := range linesFromFile {
		hashTable.Insert(num)
	}

	foundedNumbers := make([]int, 0, 100)

	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	for i := 0; i < 100; i++ {
		randomNumber := rng.Intn(hash.HasTableSize)
		if hashTable.Contains(randomNumber) {
			foundedNumbers = append(foundedNumbers, randomNumber)
		}
	}

	counter := 0
	for _, number := range foundedNumbers {
		index := file.FindValueIndex(linesFromFile, number)
		defer fmt.Printf("O número %d foi encontrado na posição original %d\n", number, index+1)
		counter++
	}

	duration := time.Since(start)
	fmt.Println("Benchmark:", duration)
	fmt.Println("Counter:", counter)
}
