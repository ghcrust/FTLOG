package bookstore

import (
	"errors"
)

type Book struct {
	ID              int
	Title           string
	Author          string
	Copies          int
	PriceCents      int
	DiscountPercent int
}

type Catalog map[int]Book

func (c Catalog) GetAllBooks() []Book {
	var books []Book
	for _, v := range c {
		books = append(books, v)
	}
	return books

}

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("No copies left!")
	}
	b.Copies--
	return b, nil
}

func (c Catalog) GetBook(id int) (Book, error) {
	if b, ok := c[id]; ok {
		return b, nil
	}
	return Book{}, errors.New("index not in catalog!")
}

/*
func GetAllBooks(catalog map[int]Book) []Book {
	//books := []Book{}
	var books []Book
	for _, book := range catalog {
		books = append(books, book)
	}
	return books
}

func GetBook(catalog map[int]Book, id int) (Book, error) {
	for _, book := range catalog {
		if book.ID == id {
			return book, nil
		}
	}
	if book ,ok := catalog[id];ok {
		return book, nil
	}
	return Book{}, errors.New("Id not found in catalog!")
}
*/

func (b Book) NetPriceCents() int {
	return b.PriceCents - (b.PriceCents*b.DiscountPercent)/100
}
