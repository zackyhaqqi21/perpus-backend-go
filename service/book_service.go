package service

import (
	"perpusbackendgo/model"
	"perpusbackendgo/repository"
)

// BookService tempat naruh business logic/validasi tambahan nanti
type BookService struct {
	Repo *repository.BookRepository
}

func NewBookService(repo *repository.BookRepository) *BookService {
	return &BookService{Repo: repo}
}

func (s *BookService) GetAllBooks() ([]model.Book, error) {
	return s.Repo.FindAll()
}

func (s *BookService) GetBookByID(id int) (model.Book, error) {
	return s.Repo.FindByID(id)
}

func (s *BookService) CreateBook(book *model.Book) error {
	// Contoh tempat nambah validasi nanti, misal:
	// if book.Stock < 0 { return errors.New("stok tidak boleh negatif") }
	return s.Repo.Create(book)
}

func (s *BookService) UpdateBook(id int, newData model.Book) (model.Book, error) {
	book, err := s.Repo.FindByID(id)
	if err != nil {
		return book, err
	}

	// Update field satu-satu (biar ID nggak ketimpa 0 kalau field kosong)
	book.Code = newData.Code
	book.Title = newData.Title
	book.Author = newData.Author
	book.Publisher = newData.Publisher
	book.PublishYear = newData.PublishYear
	book.Stock = newData.Stock

	err = s.Repo.Update(&book)
	return book, err
}

func (s *BookService) DeleteBook(id int) error {
	return s.Repo.Delete(id)
}
