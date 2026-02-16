package database

type Book struct {
	ID          int     `json:id gorm:primary_Key,autoIncrement`
	Title       string  `json:title gorm:"unique" validate:"required"`
	Author      string  `json:author validate:"required"`
	Description string  `json:description`
	ISBN        string  `json:isbn unique validate:"required"`
	Price       float64 `json:price validate:"required"`
	Stock       int     `json:stock validate:"required"`
	CategoryID  int     `json:category_id`
	CreatedAt   string  `json:created_at`
	UpdatedAt   string  `json:updated_at`
}
