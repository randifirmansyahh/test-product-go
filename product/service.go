package product

type Service interface {
	FindAll() ([]Product, error)
	FindByID(ID int) (Product, error)
	Create(productRequest ProductRequest) (Product, error)
	Update(ID int, productRequest ProductRequest) (Product, error)
	Delete(ID int) (Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Product, error) {
	return s.repository.FindAll()
}

func (s *service) FindByID(id int) (Product, error) {
	return s.repository.FindByID(id)
}

func (s *service) Create(productRequest ProductRequest) (Product, error) {
	price, _ := productRequest.Price.Int64()
	rating, _ := productRequest.Rating.Int64()
	discount, _ := productRequest.Price.Int64()
	product := Product{
		Title:       productRequest.Title,
		Price:       int(price),
		Description: productRequest.Description,
		Rating:      int(rating),
		Discount:    int(discount),
	}
	return s.repository.Create(product)
}

func (s *service) Update(ID int, productRequest ProductRequest) (Product, error) {
	find, _ := s.repository.FindByID(ID)

	price, _ := productRequest.Price.Int64()
	rating, _ := productRequest.Rating.Int64()
	discount, _ := productRequest.Price.Int64()

	find.Title = productRequest.Title
	find.Price = int(price)
	find.Description = productRequest.Description
	find.Rating = int(rating)
	find.Discount = int(discount)

	return s.repository.Update(find)
}

func (s *service) Delete(ID int) (Product, error) {
	find, _ := s.repository.FindByID(ID)
	return s.repository.Delete(find)
}
