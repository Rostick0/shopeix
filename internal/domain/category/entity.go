package category

type Category struct {
	ID         int64  `gorm:"primaryKey"`
	Name       string `gorm:"size:255;not null"`
	CategoryId *int64 `gorm:"null"`
}

type CategoryRepository interface {
	FindByID(id int64) (*Category, error)
	FindAll() ([]*Category, error)
	// Create(u *Category) error
	// Update(u *Category) error
	// Delete(id int64) error
}
