package main

import (
	"fmt"
	"structlearn/book"
)

//显式初始化，
//不建议使用
// var book2 = book.Book{"The Go Program Language",700,make(map[string]int)}
//建议使用
var book3 = book.Book{
	Title:"The Go Program Language",
	Pages: 20,
	Indexs: make(map[string]int),
}



func main() {
	var book = book.Book{}
	book.Title = "The Go Programming Language"
	book.Pages = 34
	fmt.Println(book)
}
