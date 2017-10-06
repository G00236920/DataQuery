package main

import "fmt"

type Node struct { //Struct called node
	prev *Node
	next *Node
	key  string
}

type List struct { //List double linked
	head *Node
	tail *Node
}

func (L *List) Insert(key string) { //add values to the node placed in a list
	list := &Node{
		next: L.head,
		key:  key,
	}
	if L.head != nil { //if the head value is null, then you are at the start of the list
		L.head.prev = list
	}
	L.head = list

	l := L.head
	for l.next != nil { //if you havent reached the end of the list
		l = l.next
	}
	L.tail = l
}

func (l *List) Show() { //Display the list
	fmt.Println()

	list := l.head
	for list != nil { //if the list is not empty
		fmt.Printf("%+v \n", list.key)
		list = list.next
	}

	fmt.Println()
}

func (l *List) find() { //determine the smallest and largest values based on how many characters in the string

	var smallest string
	var largest string

	list := l.head
	for list != nil {

		if len(smallest) == 0 { // find smallest string
			smallest = list.key
		}

		if len(list.key) < len(smallest) {
			smallest = list.key
		} else if len(list.key) > len(largest) { //find largest string
			largest = list.key
		}

		list = list.next //move list to next node
	}
	fmt.Println()
	fmt.Println("The Smallest Element is: ", smallest)
	fmt.Println("The largest Element is: ", largest)
	fmt.Println()
}

func main() {

	menu()

}

func menu() { //Menu Function
	l := List{}

	var choice int
	var entry string

	//output information to the screen
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
			fmt.Scan(&entry) //take user input
			l.Insert(entry)  //insert that input to the list
		case 2:
			l.find() //find function
		case 3:
			l.Show() //display function
		case 4:
			return
		default:
			fmt.Println()
			fmt.Println("INVALID OPTION") //user input invalid info
			fmt.Println()
		}

	}
}
