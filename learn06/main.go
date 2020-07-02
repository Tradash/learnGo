package main

type myData int

type DoublyLinkedNode struct {
	prev *DoublyLinkedNode
	data myData
	next *DoublyLinkedNode
}

type DoublyLinkedList struct {
	firstNode *DoublyLinkedNode
	lastNode  *DoublyLinkedNode
}

// Добавить новый элемент в конец списка
func (l *DoublyLinkedList) PushBack(d myData) {
	var n DoublyLinkedNode
	n.data = d
	n.prev = nil
	if l.firstNode == nil {
		n.next = nil
		l.firstNode = &n
		l.lastNode = &n
	} else {
		n.prev = l.lastNode
		l.lastNode.next = &n
		l.lastNode = &n
		println("1111")
	}
}

// Добавить новый элемент в начало списка
func (l *DoublyLinkedList) PushFront(d myData) {
	var n DoublyLinkedNode
	n.data = d
	n.prev = nil
	if l.firstNode == nil {
		n.next = nil
		l.firstNode = &n
		l.lastNode = &n
	} else {
		n.next = l.firstNode
		l.firstNode.prev = &n
		l.firstNode = &n

	}
}

// Получить весь список
func (l *DoublyLinkedList) PrintListFromFirst() {
	if l.firstNode == nil {
		println("Список пуст")
	} else {
		var n *DoublyLinkedNode
		n = l.firstNode
		i := 1
		for {
			println("Элемент-", i, "Данные", n.data)
			i++
			if n.next == nil {
				println("Конец списка")
				return
			} else {
				n = n.next
			}
		}
	}
}

// Напечатать елемент элемент
func (n *DoublyLinkedNode) print() {
	if n == nil {
		println("Пустой элемент")
	} else {
		println("Элемент-", 1, "Данные", n.data)
	}
}

// получить первый элемент
func (l *DoublyLinkedList) first() *DoublyLinkedNode {
	if l.firstNode == nil {
		return nil
	} else {
		return l.firstNode
	}
}

// Получить последний элемент
func (l *DoublyLinkedList) last() *DoublyLinkedNode {
	if l.lastNode == nil {
		return nil
	} else {
		return l.lastNode
	}
}

// Получить следующий элемент
func (n *DoublyLinkedNode) nextItem() *DoublyLinkedNode {
	if n.next == nil {
		return nil
	} else {
		return n.next
	}
}

// Получить Предыдущий элемент
func (n *DoublyLinkedNode) prevItem() *DoublyLinkedNode {
	if n.prev == nil {
		return nil
	} else {
		return n.prev
	}
}

// Определить длину списка
func (l *DoublyLinkedList) len() int {
	var n *DoublyLinkedNode
	if l.firstNode == nil {
		return 0
	}
	n = l.firstNode
	i := 1
	for n.next != nil {
		n = n.next
		i++
	}
	return i
}

// Напечатать все элементы с конца
func (l *DoublyLinkedList) PrintListFromLast() {
	if l.lastNode == nil {
		println("Список пуст")
	} else {
		var n *DoublyLinkedNode
		n = l.lastNode
		i := 1
		for {
			println("Элемент-", i, "Данные", n.data)
			i++
			if n.prev == nil {
				println("Конец списка")
				return
			} else {
				n = n.prev
			}
		}
	}
}

func (l *DoublyLinkedList) find(d myData) *DoublyLinkedNode {
	k := l.len()
	n := l.firstNode
	for i := 0; i < k; i++ {
		if n.data == d {
			return n
		}
		n = n.next
	}
	return nil
}
func (l *DoublyLinkedList) Remove(d myData) {
	n := l.find(d)
	if n != nil {
		if l.firstNode == n {
			l.firstNode = n.next
		}
		if l.lastNode == n {
			l.lastNode = n.prev
		}
		if n.prev != nil {
			n.prev.next = n.next
		}
		if n.next != nil {
			n.next.prev = n.prev
		}
	}
}

func main() {
	var myList DoublyLinkedList
	myList.PrintListFromFirst()
	myList.first().print()
	myList.last().print()
	println("Размер списка - ", myList.len())
	myList.PushFront(10)
	myList.PrintListFromFirst()
	myList.PushFront(20)
	myList.PushFront(25)
	println("Размер списка - ", myList.len())
	myList.PrintListFromFirst()
	myList.PushBack(30)
	myList.PrintListFromFirst()
	myList.PushBack(40)
	myList.PrintListFromFirst()
	myList.PrintListFromLast()
	println("Размер списка - ", myList.len())
	println("Получение крайних элементов")
	myList.first().print()
	myList.last().print()
	myList.first().nextItem().nextItem().print()
	myList.first().nextItem().nextItem().prevItem().print()
	println(myList.find(30))
	myList.Remove(30)
	myList.PrintListFromFirst()
	myList.PrintListFromLast()
	println("Удаляю...40")
	myList.Remove(40)
	myList.PrintListFromFirst()
	myList.PrintListFromLast()
	myList.Remove(25)
	myList.PrintListFromFirst()
	myList.PrintListFromLast()
	myList.Remove(20)
	myList.PrintListFromFirst()
	myList.PrintListFromLast()
	myList.Remove(10)
	myList.PrintListFromFirst()
	myList.PrintListFromLast()

}
