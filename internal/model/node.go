package model

// Node - узел двух связного списка
type Node struct {
	Key   int
	Value int   // значение, хранящееся в узле
	Next  *Node // указатель на следующий узел
	Prev  *Node // указатель на предыдущий узел
}

// Node - узел LRU-Cache
type LRUNode struct {
	Key   int
	Value int
}
