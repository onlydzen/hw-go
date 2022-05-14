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

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	head *ListItem
	tail *ListItem
	len  int
}

func (l list) Len() int {
	return l.len
}

func (l list) Front() *ListItem {
	return l.head
}

func (l list) Back() *ListItem {
	return l.tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	l.len++

	if l.head == nil {
		return setFirstItem(l, v)
	}

	li := &ListItem{v, l.head, nil}
	l.head.Prev = li
	l.head = li

	return li
}

func (l *list) PushBack(v interface{}) *ListItem {
	l.len++

	if l.tail == nil {
		return setFirstItem(l, v)
	}

	li := &ListItem{v, nil, l.tail}
	l.tail.Next = li
	l.tail = li

	return li
}

func (l *list) Remove(i *ListItem) {
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}

	if i.Next != nil {
		i.Next.Prev = i.Prev
	}

	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if i.Prev == nil {
		return
	}

	if l.tail == i {
		l.tail = i.Prev
	}

	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}

	i.Prev = nil
	i.Next = l.head
	l.head.Prev = i

	l.head = i
}

func NewList() List {
	return new(list)
}

func setFirstItem(l *list, v interface{}) *ListItem {
	li := &ListItem{v, nil, nil}
	l.head = li
	l.tail = li

	return li
}
