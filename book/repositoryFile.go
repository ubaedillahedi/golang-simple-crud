package book

import "fmt"

type repositoryFile struct {
}

func NewRepositoryFile() *repositoryFile {
	return &repositoryFile{}
}

func (r *repositoryFile) FindAll() ([]Book, error) {
	var books []Book

	fmt.Println("FindAll")

	return books, nil
}

func (r *repositoryFile) FindByID(id int) (Book, error) {
	var book Book
	fmt.Println("FindById")
	return book, nil

}

func (r *repositoryFile) Create(book Book) (Book, error) {
	fmt.Println("Create")

	return book, nil
}