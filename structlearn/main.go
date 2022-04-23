package main

import (
	"structlearn/book"
	"fmt"
)

func main()  {
	var book = book.Book{}
	book.Title = "The Go Programming Language"
	book.Pages = 34
	fmt.Println(book)
}