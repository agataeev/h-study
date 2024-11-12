package model

// Определяем тип Set для int
type IntSet map[int]struct{}

// Создаем новый пустой set
func NewIntSet() IntSet {
	return make(IntSet)
}

// Добавляем элемент в set
func (s IntSet) Add(value int) {
	s[value] = struct{}{}
}

// Удаляем элемент из set
func (s IntSet) Remove(value int) {
	delete(s, value)
}

// Проверяем, есть ли элемент в set
func (s IntSet) Contains(value int) bool {
	_, exists := s[value]
	return exists
}

// Возвращаем объединение двух множеств
func (s IntSet) Union(other IntSet) IntSet {
	result := NewIntSet()
	for value := range s {
		result.Add(value)
	}
	for value := range other {
		result.Add(value)
	}
	return result
}

// Возвращаем пересечение двух множеств
func (s IntSet) Intersection(other IntSet) IntSet {
	result := NewIntSet()
	for value := range s {
		if other.Contains(value) {
			result.Add(value)
		}
	}
	return result
}

// Возвращаем разность двух множеств
func (s IntSet) Difference(other IntSet) IntSet {
	result := NewIntSet()
	for value := range s {
		if !other.Contains(value) {
			result.Add(value)
		}
	}
	return result
}
