package main

import "fmt"

type Node struct {
	prev *Node
	next *Node
	key  string
}

type List struct {
	head *Node
	tail *Node
}

func (L *List) Insert(key string) {
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
	fmt.Println()

	list := l.head
	for list != nil {
		fmt.Printf("%+v \n", list.key)
		list = list.next
	}

	fmt.Println()
}

func (l *List) find() {

	var smallest string
	var largest string

	list := l.head
	for list != nil {

		if len(smallest) == 0 {
			smallest = list.key
		}

		if len(list.key) < len(smallest) {
			smallest = list.key
		} else if len(list.key) > len(largest) {
			largest = list.key
		}

		list = list.next
	}
	fmt.Println()
	fmt.Println("The Smallest Element is: ", smallest)
	fmt.Println("The largest Element is: ", largest)
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
		fmt.Println()
		fmt.Println("Please Choose an Option")
		fmt.Println("\t1. Add Element to List")
		fmt.Println("\t2. Display Largest and Smallest Elements (By Character Length)")
		fmt.Println("\t3. Display the Entire List")
		fmt.Println("\t4. Exit")
		fmt.Println()
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Enter the Value for the Element")
			fmt.Scan(&entry)
			l.Insert(entry)
		case 2:
			l.find()
		case 3:
			l.Show()
		case 4:
			return
		default:
			fmt.Println()
			fmt.Println("INVALID OPTION")
			fmt.Println()
		}

	}
}
