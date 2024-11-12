package main

func main() {

	/*	set := model.NewSet[int]()
		set.Add(1)
		set.Add(2)
		set.Add(3)
		set.Add(3)
		set.Add(4)
		set.Add(5)

		setFromSlice := model.NewSetFromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
		fmt.Println(setFromSlice.Elements())

		set.Remove(3)

		set1 := model.NewSet[int]()
		set1.Add(3)
		set1.Add(4)
		set1.Add(5)

		union := set.Union(set1)

		union.Elements()

		fmt.Println(union.Elements())*/

	/*list := model.LinkedList{}

	// Добавляем элементы
	list.AddToFront(3)
	list.AddToFront(2)
	list.AddToFront(1)
	list.AddToEnd(4)

	// Поиск элемента
	found := list.Find(3)
	if found != nil {
		fmt.Println("Найдено:", found.Value) // Вывод: Найдено: 3
	} else {
		fmt.Println("Элемент не найден")
	}

	// Удаление элемента
	list.Delete(2)

	// Выводим оставшиеся элементы
	current := list.Head
	for current != nil {
		fmt.Print(current.Value, " ")
		current = current.Next
	}
	// Вывод: 1 3 4*/

	/*doubleList := model.DoublyLinkedList{}

	// Добавляем элементы
	doubleList.AddToFront(2)
	doubleList.AddToFront(1)
	doubleList.AddToEnd(3)
	doubleList.AddToEnd(4)

	// Поиск элемента
	found := doubleList.Find(3)
	if found != nil {
		fmt.Println("Найдено:", found.Value) // Вывод: Найдено: 3
	} else {
		fmt.Println("Элемент не найден")
	}

	// Удаление элемента
	doubleList.Delete(2)

	// Выводим элементы от головы к хвосту
	current := doubleList.Head
	for current != nil {
		fmt.Print(current.Value, " ")
		current = current.Next
	}
	// Вывод: 1 3 4

	fmt.Println()

	// Выводим элементы от хвоста к голове
	current = doubleList.Tail
	for current != nil {
		fmt.Print(current.Value, " ")
		current = current.Prev
	}
	// Вывод: 4 3 1*/
}

func IncrementPointer(a *int) {
	*a++
}

func SwapPointers(a, b *int) {
	*a, *b = *b, *a
}

type User struct {
	Name string
}

func (u *User) UpdateName(name string) {
	u.Name = name
}

type SimpleLinkedList struct {
}
