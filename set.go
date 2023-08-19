package go_set

import "fmt"

type Set[T comparable] struct {
	set map[T]interface{}
}

func New[T comparable]() *Set[T] {
	return &Set[T]{set: make(map[T]interface{})}
}

func NewFrom[T comparable](list *[]T) *Set[T] {
	result := make(map[T]interface{})
	for i := 0; i < len(*list); i++ {
		result[(*list)[i]] = nil
	}
	return &Set[T]{set: result}
}

func (s *Set[T]) Len() int {
	return len(s.set)
}

func (s *Set[T]) Add(element T) {
	s.set[element] = nil
}

func (s *Set[T]) Clear() {
	s.set = make(map[T]interface{})
}

func (s *Set[T]) Copy() *Set[T] {
	newSet := make(map[T]interface{})
	for k := range s.set {
		newSet[k] = nil
	}
	return &Set[T]{set: newSet}
}

func (s *Set[T]) Difference(set *Set[T]) *Set[T] {
	result := make(map[T]interface{})
	for k := range s.set {
		if _, ok := set.set[k]; !ok {
			result[k] = nil
		}
	}
	return &Set[T]{set: result}
}

func (s *Set[T]) DifferenceUpdate(set *Set[T]) {
	for k := range s.set {
		if _, ok := set.set[k]; ok {
			delete(s.set, k)
		}
	}
}

func (s *Set[T]) Discard(element T) {
	delete(s.set, element)
}

func (s *Set[T]) Intersection(set *Set[T]) *Set[T] {
	result := make(map[T]interface{})
	for k := range s.set {
		if _, ok := set.set[k]; ok {
			result[k] = nil
		}
	}
	return &Set[T]{set: result}
}

func (s *Set[T]) IntersectionUpdate(set *Set[T]) {
	for k := range s.set {
		if _, ok := set.set[k]; !ok {
			delete(s.set, k)
		}
	}
}

func (s *Set[T]) IsDisjoint(set *Set[T]) bool {
	return 0 == (s.Intersection(set)).Len()
}

func (s *Set[T]) IsSubset(set *Set[T]) bool {
	for k := range s.set {
		if _, ok := set.set[k]; !ok {
			return false
		}
	}
	return true
}

func (s *Set[T]) IsSuperset(set *Set[T]) bool {
	for k := range set.set {
		if _, ok := s.set[k]; !ok {
			return false
		}
	}
	return true
}

func (s *Set[T]) Remove(element T) error {
	if _, ok := s.set[element]; ok {
		delete(s.set, element)
		return nil
	}
	return fmt.Errorf("KeyError: %v", element)
}

func (s *Set[T]) Union(set *Set[T]) *Set[T] {
	result := make(map[T]interface{})
	for k := range s.set {
		result[k] = nil
	}
	for k := range set.set {
		result[k] = nil
	}
	return &Set[T]{set: result}
}

func (s *Set[T]) UnionUpdate(set *Set[T]) {
	for k := range set.set {
		s.set[k] = nil
	}
}
