package main

import "fmt"

type Node struct {
	prev *Node
	next *Node
	key  interface{}
}

type List struct {
	head *Node
	tail *Node
}

func (L *List) Insert(key interface{}) {
	list := &Node{
		next: L.head,
		key:  key,
	}
	if L.head != nil {
		L.head.prev = list
	}
	L.head = list

	l := L.head
	for l.next != nil {
		l = l.next
	}
	L.tail = l
}

func (l *List) Show() {
	list := l.head
	for list != nil {
		fmt.Printf("%+v \n", list.key)
		list = list.next
	}
	fmt.Println()
}

func main() {

	menu()

}

func menu() {
	l := List{}

	var choice int
	var entry string

	for i := 0; i < 1; {

		fmt.Println("Please Choose an Option")
		fmt.Println("1. Add Element to List")
		fmt.Println("2. Exit")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Enter the Value for the Element")
			fmt.Scan(&entry)
			l.Insert(entry)
		case 2:
			return
		default:
			fmt.Println("Invalid Option")
		}

		fmt.Printf("head: %v\n", l.head.key)
		fmt.Printf("tail: %v\n", l.tail.key)
		l.Show()

	}

}
