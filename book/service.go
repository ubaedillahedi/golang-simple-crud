package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(id int) (Book, error)
	Create(book BookRequest) (Book, error)
	Update(id int, book BookRequest) (Book, error)
	Delete(id int) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	books, err := s.repository.FindAll()
	return books, err
	// return s.repository.FindAll()
}

func (s *service) FindByID(id int) (Book, error) {
	book, err := s.repository.FindByID(id)
	return book, err
}

func (s *service) Create(bookRequest BookRequest) (Book, error) {
	price, _ := bookRequest.Price.Int64()
	discount, _ := bookRequest.Discount.Int64()
	rating, _ := bookRequest.Rating.Int64()
	book := Book{
		Title:       bookRequest.Title,
		Description: bookRequest.Description,
		Price:       int(price),
		Discount:    int(discount),
		Rating:      int(rating),
	}
	newBook, err := s.repository.Create(book)
	return newBook, err
}

func (s *service) Update(id int, bookRequest BookRequest) (Book, error) {
	b, err := s.repository.FindByID(id)

	price, _ := bookRequest.Price.Int64()
	discount, _ := bookRequest.Discount.Int64()
	rating, _ := bookRequest.Rating.Int64()

	b.Title = bookRequest.Title
	b.Description = bookRequest.Description
	b.Price = int(price)
	b.Discount = int(discount)
	b.Rating = int(rating)

	newBook, err := s.repository.Update(b)

	return newBook, err
}

func (s *service) Delete(id int) (Book, error) {
	b, _ := s.repository.FindByID(id)

	_, err := s.repository.Delete(b)

	return b, err

}