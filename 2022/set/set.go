package set

type Set[T comparable] struct {
	set map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		set: make(map[T]struct{}),
	}
}

func (s *Set[T]) Set(value T) {
	s.set[value] = struct{}{}
}

func (s *Set[T]) Has(value T) bool {
	_, hasValue := s.set[value]
	return hasValue
}

func (s *Set[T]) Delete(value T) {
	delete(s.set, value)
}

func (s *Set[T]) Values() []T {
	values := []T{}
	for value := range s.set {
		values = append(values, value)
	}
	return values
}
