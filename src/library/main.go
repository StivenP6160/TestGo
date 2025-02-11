package main

import "library/book"

func main()  {
    var myBook = book.Book{
        Title: "Moby Dick",
        Author: "Herman Melville",
        Pages: 300,
    }

    myBook.printInfo()
}