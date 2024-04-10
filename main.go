package main

import "fmt"

type Books struct {
	title string
	id    int
}

func main() {
	var book1 Books
	var book2 Books
	book1.title = "Csgeek"
	book1.id = 1
	book2.title = "Programming"
	book2.id = 2
	fmt.Printf("book title is: %s", book1.title)

}
