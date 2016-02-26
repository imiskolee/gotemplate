// Template Set type
//
// Tries to be similar to Python's set type
package main

// template type Set(A)

// Store struct{} which are 0 size as the members in the map
type mySetNothing struct{}

type mySet struct {
	m map[string]mySetNothing
}

// Returns a new empty set with the given capacity
func newSizedMySet(capacity int) *mySet {
	return &mySet{
		m: make(map[string]mySetNothing, capacity),
	}
}

// Returns a new empty set
func newMySet() *mySet {
	return newSizedMySet(0)
}

// Returns the number of elements in the set
func (s *mySet) Len() int {
	return len(s.m)
}

// Contains returns whether elem is in the set or not
func (s *mySet) Contains(elem string) bool {
	_, found := s.m[elem]
	return found
}

// Add adds elem to the set, returning the set
//
// If the element already exists then it has no effect
func (s *mySet) Add(elem string) *mySet {
	s.m[elem] = mySetNothing{}
	return s
}

// AddList adds a list of elems to the set
//
// If the elements already exists then it has no effect
func (s *mySet) AddList(elems []string) *mySet {
	for _, elem := range elems {
		s.m[elem] = mySetNothing{}
	}
	return s
}

// Discard removes elem from the set
//
// If it wasn't in the set it does nothing
//
// It returns the set
func (s *mySet) Discard(elem string) *mySet {
	delete(s.m, elem)
	return s
}

// Remove removes elem from the set
//
// It returns whether the elem was in the set or not
func (s *mySet) Remove(elem string) bool {
	_, found := s.m[elem]
	if found {
		delete(s.m, elem)
	}
	return found
}

// Pop removes elem from the set and returns it
//
// It also returns whether the elem was found or not
func (s *mySet) Pop(elem string) (string, bool) {
	_, found := s.m[elem]
	if found {
		delete(s.m, elem)
	}
	return elem, found
}

// AsList returns all the elements as a slice
func (s *mySet) AsList() []string {
	elems := make([]string, len(s.m))
	i := 0
	for elem := range s.m {
		elems[i] = elem
		i++
	}
	return elems
}

// Clear removes all the elements
func (s *mySet) Clear() *mySet {
	s.m = make(map[string]mySetNothing)
	return s
}

// Copy returns a shallow copy of the Set
func (s *mySet) Copy() *mySet {
	newSet := newSizedMySet(len(s.m))
	for elem := range s.m {
		newSet.m[elem] = mySetNothing{}
	}
	return newSet
}

// Difference returns a new set with all the elements that are in this
// set but not in the other
func (s *mySet) Difference(other *mySet) *mySet {
	newSet := newSizedMySet(len(s.m))
	for elem := range s.m {
		if _, found := other.m[elem]; !found {
			newSet.m[elem] = mySetNothing{}
		}
	}
	return newSet
}

// DifferenceUpdate removes all the elements that are in the other set
// from this set.  It returns the set.
func (s *mySet) DifferenceUpdate(other *mySet) *mySet {
	m := s.m
	for elem := range other.m {
		delete(m, elem)
	}
	return s
}

// Intersection returns a new set with all the elements that are only in this
// set and the other set. It returns the new set.
func (s *mySet) Intersection(other *mySet) *mySet {
	newSet := newSizedMySet(len(s.m) + len(other.m))
	for elem := range s.m {
		if _, found := other.m[elem]; found {
			newSet.m[elem] = mySetNothing{}
		}
	}
	for elem := range other.m {
		if _, found := s.m[elem]; found {
			newSet.m[elem] = mySetNothing{}
		}
	}
	return newSet
}

// IntersectionUpdate changes this set so that it only contains
// elements that are in both this set and the other set.  It returns
// the set.
func (s *mySet) IntersectionUpdate(other *mySet) *mySet {
	for elem := range s.m {
		if _, found := other.m[elem]; !found {
			delete(s.m, elem)
		}
	}
	return s
}

// Union returns a new set with all the elements that are in either
// set. It returns the new set.
func (s *mySet) Union(other *mySet) *mySet {
	newSet := newSizedMySet(len(s.m) + len(other.m))
	for elem := range s.m {
		newSet.m[elem] = mySetNothing{}
	}
	for elem := range other.m {
		newSet.m[elem] = mySetNothing{}
	}
	return newSet
}

// Update adds all the elements from the other set to this set.
// It returns the set.
func (s *mySet) Update(other *mySet) *mySet {
	for elem := range other.m {
		s.m[elem] = mySetNothing{}
	}
	return s
}

// IsSuperset returns a bool indicating whether this set is a superset of other set.
func (s *mySet) IsSuperset(strict bool, other *mySet) bool {
	if strict && len(other.m) >= len(s.m) {
		return false
	}
string:
	for v := range other.m {
		for i := range s.m {
			if v == i {
				continue string
			}
		}
		return false
	}
	return true
}

// IsSubset returns a bool indicating whether this set is a subset of other set.
func (s *mySet) IsSubset(strict bool, other *mySet) bool {
	if strict && len(s.m) >= len(other.m) {
		return false
	}
string:
	for v := range s.m {
		for i := range other.m {
			if v == i {
				continue string
			}
		}
		return false
	}
	return true
}

// IsDisjoint returns a bool indicating whether this set and other set have no elements in common.
func (s *mySet) IsDisjoint(other *mySet) bool {
	for v := range s.m {
		if other.Contains(v) {
			return false
		}
	}
	return true
}

// SymmetricDifference returns a new set of all elements that are a member of exactly
// one of this set and other set(elements which are in one of the sets, but not in both).
func (s *mySet) SymmetricDifference(other *mySet) *mySet {
	work1 := s.Union(other)
	work2 := s.Intersection(other)
	for v := range work2.m {
		delete(work1.m, v)
	}
	return work1
}

// SymmetricDifferenceUpdate modifies this set to be a set of all elements that are a member
// of exactly one of this set and other set(elements which are in one of the sets,
// but not in both) and returns this set.
func (s *mySet) SymmetricDifferenceUpdate(other *mySet) *mySet {
	work := s.SymmetricDifference(other)
	*s = *work
	return s
}
