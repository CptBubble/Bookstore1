package ds

import (
	"github.com/google/uuid"
)

type Book struct {
	UUID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key;"` //default:uuid_generate_v4()
	Name        string
	Saleprice   uint64
	Year        uint64
	Type        string
	Srokgodnost uint64
	Color       string
	Description string
	Image       string
}

type QuantityStores struct {
	Quantity uint64 `example:"10"`
}

type PriceStore struct {
	Price uint64 `example:"300"`
}

func (Book) TableName() string {
	return "book"
}
