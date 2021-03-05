package slab

type slabEntry struct {
	Value interface{}
	Index uintptr
}

func (entry slabEntry) IsValid() bool {
	return entry.Value != nil
}

// Slab is an implementation of the internal registry free list. A zero-value
// instance is a valid instance. This data structure is not thread-safe.
type Slab struct {
	entries []slabEntry
	free    uintptr
}

func (s *Slab) Put(entry interface{}) uintptr {
	if s.free == uintptr(len(s.entries)) {
		index := uintptr(len(s.entries))
		s.entries = append(s.entries, slabEntry{entry, 0})
		s.free++
		return index
	}

	index := s.free

	s.free = s.entries[index].Index
	s.entries[index] = slabEntry{entry, 0}

	return index
}

func (s *Slab) Get(i uintptr) interface{} {
	// Perform bound check.
	if i >= uintptr(len(s.entries)) {
		return nil
	}
	// Perform validity check in case of invalid ID.
	if entry := s.entries[i]; entry.IsValid() {
		return entry.Value
	}
	return nil
}

func (s *Slab) Pop(i uintptr) interface{} {
	popped := s.entries[i].Value
	s.entries[i] = slabEntry{nil, s.free}
	s.free = i
	return popped
}
