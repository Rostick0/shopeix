package product

import "time"

type Product struct {
	ID         int64     `gorm:"primaryKey"`
	Title      string    `gorm:"size:255;not null"`
	OldPrice   float32   `gorm:"not null"`
	Price      float32   `gorm:"not null"`
	CategoryId *int64    `gorm:"null"`
	Rating     *float32  `gorm:"null"`
	IsShow     bool      `gorm:"default:false"`
	IsChecked  bool      `gorm:"default:false"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

type CreateProductRequest struct {
	Title      string  `validate:"required,min=2,max=255"`
	OldPrice   float32 `validate:"number,min=0"`
	Price      float32 `validate:"required,number,min=0"`
	CategoryId *int64
	IsShow     bool
}

type UpdateProductRequest struct {
	Title      *string  `validate:"required,min=2,max=255"`
	OldPrice   *float32 `validate:"number,min=0"`
	Price      *float32 `validate:"required,number,min=0"`
	CategoryId *int64
	IsShow     *bool
}

type GetListProductRequest struct {
	Title      string  `validate:"min=2,max=255"`
	OldPrice   float32 `validate:"number,min=0"`
	Price      float32 `validate:"required,number,min=0"`
	CategoryId *int64
	IsShow     bool
}
