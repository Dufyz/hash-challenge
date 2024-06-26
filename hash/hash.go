package hash

const (
	HashTableSize = 256
)

type HashTable struct {
	table [HashTableSize]bool
	index map[int]int
}

func NewHashTable(size int) *HashTable {
	return &HashTable{
		index: make(map[int]int, size),
	}
}

func (h *HashTable) Insert(number, i int) {
	if number >= 0 && number < HashTableSize {
		h.table[number] = true
		h.index[number] = i
	}
}

func (h *HashTable) Contains(number int) bool {
	if number >= 0 && number < HashTableSize {
		return h.table[number]
	}
	return false
}

func (h *HashTable) IndexOf(number int) int {
	if index, exists := h.index[number]; exists {
		return index
	}
	return -1
}
