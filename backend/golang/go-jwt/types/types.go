package types

import "time"

type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName"  validate:"required"`
	Email     string `json:"email"  validate:"required,email"`
	Password  string `json:"password"  validate:"required,min=3,max=32"`
}
type LoginUserPayload struct {
	Email    string `json:"email"  validate:"required,email"`
	Password string `json:"password"  validate:"required"`
}

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

// type Product struct {
// 	ID          int     `json:"id"`
// 	Name        string  `json:"name"`
// 	Description string  `json:"description"`
// 	Image       string  `json:"image"`
// 	Price       float64 `json:"price"`
// 	// note that this isn't the best way to handle quantity
// 	// because it's not atomic (in ACID), but it's good enough for this example
// 	Quantity  int       `json:"quantity"`
// 	CreatedAt time.Time `json:"createdAt"`
// }

type Product struct {
	ID          int       `json:"id" validate:"required"`
	Name        string    `json:"name" validate:"required,min=2,max=100"`
	Description string    `json:"description" validate:"required,min=10"`
	Image       string    `json:"image" validate:"required,url"`
	Price       float64   `json:"price" validate:"required,gt=0"`
	Quantity    int       `json:"quantity" validate:"gte=0"`
	CreatedAt   time.Time `json:"createdAt" validate:"required"`
}

// Repository pattern for implementing Data access layer
// With Service Layer
type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(user User) error
}

type ProductStore interface {
	GetAllProducts() ([]Product, error)
	CreateProduct(product Product) error
}
