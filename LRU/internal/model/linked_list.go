package model

// LinkedList - сам односвязный список
type LinkedList struct {
	Head *Node // ссылка на первый элемент списка
}

func (list *LinkedList) AddToFront(value int) {
	newNode := &Node{Value: value}
	newNode.Next = list.Head // новый узел указывает на текущую голову
	list.Head = newNode      // голова становится новым узлом
}

func (list *LinkedList) AddToEnd(value int) {
	newNode := &Node{Value: value}
	if list.Head == nil {
		list.Head = newNode
		return
	}
	current := list.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

func (list *LinkedList) Find(value int) *Node {
	current := list.Head
	for current != nil {
		if current.Value == value {
			return current
		}
		current = current.Next
	}
	return nil // элемент не найден
}

func (list *LinkedList) Delete(value int) {
	if list.Head == nil {
		return
	}

	// если удаляем голову
	if list.Head.Value == value {
		list.Head = list.Head.Next
		return
	}

	current := list.Head
	for current.Next != nil && current.Next.Value != value {
		current = current.Next
	}

	// если найден элемент, то пропускаем его
	if current.Next != nil {
		current.Next = current.Next.Next
	}
}
