package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

// ListItem Предоставляет элемент списка
type ListItem struct {
	Value interface{} // Значение элемента
	Next  *ListItem   // Следующий элемент
	Prev  *ListItem   // Предыдущий элемент
}

// list представляет двусвязный список
type list struct {
	front *ListItem
	back  *ListItem
	len   int
}

// Len предоставляет длину списка
func (l *list) Len() int {
	return l.len
}

// Front возвращает первый элемент списка
func (l *list) Front() *ListItem {
	return l.front
}

// Back возвращает последний элемент списка
func (l *list) Back() *ListItem {
	return l.back
}

// PushFront добавляет элемент в начало списка
func (l *list) PushFront(v interface{}) *ListItem {
	item := &ListItem{Value: v}

	if l.len == 0 {
		l.front = item
		l.back = item
	} else {
		item.Next = l.front
		l.front.Prev = item
		l.front = item
	}
	l.len++

	return item
}

// PushBack добавляет элемент в начало списка
func (l *list) PushBack(v interface{}) *ListItem {
	item := &ListItem{Value: v}

	if l.len == 0 {
		l.front = item
		l.back = item
	} else {
		item.Prev = l.back
		l.back.Next = item
		l.back = item
	}
	l.len++

	return item
}

// Remove удаление элемента из списка
func (l *list) Remove(i *ListItem) {
	if i == l.front {
		l.front = i.Next
	} else {
		i.Prev.Next = i.Next
	}

	if i == l.back {
		l.back = i.Prev
	} else {
		i.Next.Prev = i.Prev
	}

	i.Next = nil
	i.Prev = nil
	l.len--
}

// MoveToFront перемещение в начало списка
func (l *list) MoveToFront(i *ListItem) {
	if i == l.front {
		return
	}

	// Убираем элемент из текущего положения
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}
	if i == l.back {
		l.back = i.Prev
	}

	// Перемещаем элемент в начало
	i.Prev = nil
	i.Next = l.front
	l.front.Prev = i
	l.front = i
}

func NewList() List {
	return new(list)
}
