package product

import "time"

type Product struct {
	ID          int
	Title       string
	Price       int
	Description string
	Rating      int
	Discount    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
