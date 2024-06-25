package hash

const (
	HasTableSize = 10000
)

type HashTable struct {
	table map[int]bool
}

func NewHashTable() *HashTable {
	return &HashTable{
		table: make(map[int]bool),
	}
}

func (h HashTable) Insert(number int) {
	h.table[number] = true
}

func (h *HashTable) Contains(number int) bool {
	_, exists := h.table[number]
	return exists
}
