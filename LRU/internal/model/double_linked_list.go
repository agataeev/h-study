package model

// DoublyLinkedList - сам двусвязный список
type DoublyLinkedList struct {
	Head *Node
	Tail *Node
}

// AddToFront - добавляет элемент в начало списка
func (list *DoublyLinkedList) AddToFront(value int) {
	newNode := &Node{Value: value}
	if list.Head == nil {
		list.Head = newNode
		list.Tail = newNode
		return
	}
	newNode.Next = list.Head
	list.Head.Prev = newNode
	list.Head = newNode
}

// AddToEnd - добавляет элемент в конец списка
func (list *DoublyLinkedList) AddToEnd(value int) {
	newNode := &Node{Value: value}
	if list.Head == nil {
		list.Head = newNode
		list.Tail = newNode
		return
	}
	list.Tail.Next = newNode
	newNode.Prev = list.Tail
	list.Tail = newNode
}

// Find - ищет элемент в списке
func (list *DoublyLinkedList) Find(value int) *Node {
	current := list.Head
	for current != nil {
		if current.Value == value {
			return current
		}
		current = current.Next
	}
	return nil
}

// Delete - удаляет элемент из списка
func (list *DoublyLinkedList) Delete(value int) {
	current := list.Head

	// Проходим по списку в поисках элемента
	for current != nil {
		if current.Value == value {
			// Если это голова списка
			if current.Prev == nil {
				list.Head = current.Next
				if list.Head != nil {
					list.Head.Prev = nil
				}
			} else {
				current.Prev.Next = current.Next
			}

			// Если это хвост списка
			if current.Next == nil {
				list.Tail = current.Prev
				if list.Tail != nil {
					list.Tail.Next = nil
				}
			} else {
				current.Next.Prev = current.Prev
			}
			return
		}
		current = current.Next
	}
}
