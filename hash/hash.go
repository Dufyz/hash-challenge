package hash

const (
	HasTableSize = 10000
)

type HashTable struct {
	table map[int]int
}

func NewHashTable(len int) *HashTable {
	return &HashTable{
		table: make(map[int]int, len),
	}
}

func (h HashTable) Insert(number int, i int) {
	h.table[number] = i
}

func (h *HashTable) Contains(number int) bool {
	_, exists := h.table[number]
	return exists
}

func (h *HashTable) IndexOf(number int) int {
	n, exists := h.table[number]

	if !exists {
		return -1
	}

	return n
}
