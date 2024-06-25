package main

import (
	"fmt"
	// "log"
	"math/rand"
	// "os"
	// "runtime/pprof"
	"time"

	"github.com/Dufyz/hash-challenge/file"
	"github.com/Dufyz/hash-challenge/hash"
)

func main() {
	// f, err := os.Create("cpu.prof")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// pprof.StartCPUProfile(f)
	// defer pprof.StopCPUProfile()

	start := time.Now()

	linesFromFile := file.ReadFileLines("RandomNumbers.txt")
	if linesFromFile == nil {
		fmt.Println("Error reading file")
		return
	}

	hashTable := hash.NewHashTable(len(linesFromFile))
	for i, num := range linesFromFile {
		hashTable.Insert(num, i)
	}

	randomNumbers := make([]int, 0, 100)
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	for i := 0; i < 100; i++ {
		randomNumbers = append(randomNumbers, rng.Intn(hash.HasTableSize))
	}

	counter := 0
	for _, number := range randomNumbers {
		index := hashTable.IndexOf(number)
		if index != -1 {
			fmt.Printf("O número %d foi encontrado na posição original %d\n", number, index+1)
			counter++
		}
	}

	duration := time.Since(start)
	fmt.Println("Benchmark:", duration)
	fmt.Println("Counter:", counter)
}
