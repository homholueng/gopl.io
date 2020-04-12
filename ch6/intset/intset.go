package intset

func popCount(x uint) int {
	bitCount := 0
	for ; x != 0; bitCount++ {
		x &= (x - 1)
	}
	return bitCount
}

const offset = 32 << (^uint(0) >> 63)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/offset, uint(x%offset)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/offset, uint(x%offset)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// Len return the number of elements
func (s *IntSet) Len() int {
	len := 0
	for _, word := range s.words {
		len += popCount(word)
	}
	return len
}

// Remove remove x from the set
func (s *IntSet) Remove(x int) {
	word, bit := x/offset, uint(x%offset)
	if !(word < len(s.words)) {
		return
	}
	s.words[word] &= ^(1 << bit)
}

// Clear remove all elements form the set
func (s *IntSet) Clear() {
	for i := 0; i < len(s.words); i++ {
		s.words[i] &= 0
	}
}

// Copy return a copy of the set
func (s *IntSet) Copy() *IntSet {
	c := IntSet{words: make([]uint, len(s.words))}
	copy(c.words, s.words)
	return &c
}

// AddAll add all number in params to s
func (s *IntSet) AddAll(numbers ...int) {
	for _, n := range numbers {
		s.Add(n)
	}
}

// UnionWith sets s to the union of s and t
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// IntersectWith sets s to the intersect of s and t
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		} else {
			s.words = append(s.words, 0)
		}
	}
}

// DifferenceWith sets s to the difference of s and t
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		} else {
			s.words = append(s.words, 0)
		}
	}
}

// SymmetricDifference sets s to the symmetric difference of s and t
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// Elems return all int store in s
func (s *IntSet) Elems() []int {
	elems := make([]int, 0)
	for base, word := range s.words {
		for bit := 0; bit < offset; bit++ {
			if (word & (1 << uint(bit))) != 0 {
				elems = append(elems, offset*base+bit)
			}
		}
	}
	return elems
}
