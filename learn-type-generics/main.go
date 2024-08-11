package main

import "fmt"

// Node rappresenta un singolo nodo nella lista collegata
type Node[T any] struct {
	value T
	next  *Node[T]
}

// LinkedList rappresenta la lista collegata
type LinkedList[T any] struct {
	head *Node[T]
}

// Add aggiunge un nuovo valore alla lista
func (l *LinkedList[T]) Add(value T) {
	newNode := &Node[T]{value: value}
	if l.head == nil {
		l.head = newNode
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

// Print stampa tutti i valori nella lista
func (l *LinkedList[T]) Print() {
	current := l.head
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
}

func main() {
	// Creiamo una lista di numeri interi
	list := LinkedList[int]{}
	list.Add(10)
	list.Add(20)
	list.Add(30)
	list.Print()

	// Creiamo una lista di stringhe
	strList := LinkedList[string]{}
	strList.Add("Hello")
	strList.Add("World")
	strList.Print()
}
