package hash_test

import (
	"testing"

	"github.com/Dufyz/hash-challenge/hash"
)

func TestHashTable(t *testing.T) {
	ht := hash.NewHashTable()

	// Teste para verificar a inserção e a existência de um número
	testNumber := 12345
	ht.Insert(testNumber)
	if !ht.Contains(testNumber) {
		t.Errorf("HashTable.Contains() = false; want true for number %d", testNumber)
	}

	// Teste para verificar que números não inseridos não são encontrados
	uninsertedNumber := 54321
	if ht.Contains(uninsertedNumber) {
		t.Errorf("HashTable.Contains() = true; want false for number %d", uninsertedNumber)
	}
}

func TestHashTableMultipleInserts(t *testing.T) {
	ht := hash.NewHashTable()

	// Inserir múltiplos números e verificar a existência
	numbersToInsert := []int{100, 200, 300, 400, 500}
	for _, num := range numbersToInsert {
		ht.Insert(num)
		if !ht.Contains(num) {
			t.Errorf("HashTable.Contains() = false; want true for number %d", num)
		}
	}

	// Verificar um número que não foi inserido
	if ht.Contains(600) {
		t.Errorf("HashTable.Contains() = true; want false for number not inserted")
	}
}
