package main

func hash32(s string) uint32 {
	h := uint32(3)
	for _, r := range s {
		h ^= uint32(r)
		h *= 0x80000057
		h ^= h >> 15
	}
	return h
}

type hashTable struct {
	count   int
	buckets []*item
}

type item struct {
	key   string
	value int
	next  *item
}

func (h *hashTable) bucket(key string) uint32 {
	return hash32(key) % uint32(len(h.buckets))
}

func (h *hashTable) grow() {
	oldbuckets := h.buckets

	// alocate twice as much space for new items
	h.buckets = make([]*item, len(oldbuckets)*2)

	// reset count
	h.count = 0

	// add old items to new buckets
	for _, b := range oldbuckets {
		for b != nil {
			h.Put(b.key, b.value)
			b = b.next
		}
	}
}

func NewHashTableSize(size int) *hashTable {
	return &hashTable{
		count:   0,
		buckets: make([]*item, size),
	}
}

func (h *hashTable) Put(key string, value int) {
	// get bucket
	b := &h.buckets[h.bucket(key)]
	for *b != nil {
		// update if value if keys equals
		if (*b).key == key {
			(*b).value = value
			break
		}
		// or try next
		b = &(*b).next
	}
	// add new item
	if *b == nil {
		*b = &item{key, value, nil}
		h.count++
	}
	if h.count > len(h.buckets) {
		h.grow()
	}
}

func (h *hashTable) Get(key string) (int, bool) {
	// get bucket
	b := h.buckets[h.bucket(key)]
	// itterate bucket
	for b != nil {
		// if we have key, return value
		if b.key == key {
			return b.value, true
		}
		b = b.next
	}

	// we dont have key
	return 0, false
}

func (h *hashTable) Contains(key string) bool {
	_, ok := h.Get(key)
	return ok
}
