package book

import "fmt"

type Book struct {
	Title string
	Author string
	Pages int
}

func (b *Book) printInfo() {
	fmt.printF("Title: %s\nAuthor: %s\nPages: %d\n", b.Title, b.Author, b.Pages)
}