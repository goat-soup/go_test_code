package set

type Set[T comparable] struct {
	e map[T]bool
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{e: make(map[T]bool)}
}

func (s *Set[T]) Contains(key T) bool {
	_, v := s.e[key]
	return v
}

func (s *Set[T]) Add(key T) {
	s.e[key] = true
}

func (s *Set[T]) Remove(key T) {
	delete(s.e, key)
}

func (s *Set[T]) Size() int {
	return len(s.e)
}

func (s *Set[T]) ToSlice() []T {
	keys := make([]T, 0, s.Size())
	for k := range s.e {
		keys = append(keys, k)
	}
	return keys
}
