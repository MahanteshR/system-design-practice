package main

import (
	"fmt"
	"library-management/pkg"
)

func main() {
	library := pkg.Library{
		Racks: []pkg.Rack{
			{RackID: "R1", Capacity: 10, Books: make(map[string]int)},
			{RackID: "R2", Capacity: 10, Books: make(map[string]int)},
		},
	}

	book := pkg.Book{ID: "B1", Title: "Go Programming", Authors: []string{"John Doe"}}

	library.AddBook(book, 5)

	err := library.BorrowBook("B1", "Alice", "2024-09-01")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Book borrowed successfully")
	}

}
