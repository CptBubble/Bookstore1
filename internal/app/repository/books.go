package repository

import (
	"Bookstore/internal/app/ds"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *Repository) GetBooks() ([]ds.Book, error) {
	var stores []ds.Book
	err := r.db.Order("uuid").Find(&stores).Error
	return stores, err
}

func (r *Repository) GetBook(UUID uuid.UUID) (ds.Book, error) {
	var store ds.Book
	err := r.db.First(&store, UUID).Error
	return store, err
}
func (r *Repository) GetBookName(uuid uuid.UUID) (string, error) {
	var store ds.Book
	err := r.db.Select("name").First(&store, "uuid = ?", uuid).Error
	return store.Name, err
}

func (r *Repository) CreateBook(store ds.Book) error {
	err := r.db.Create(&store).Error
	return err
}

var image = map[string]string{
	"Пятёрочка":   "https://res.cloudinary.com/dh4qv3hob/image/upload/v1667665906/Promos/Five_gioiio.png",
	"Магнит":      "https://res.cloudinary.com/dh4qv3hob/image/upload/v1667665905/Promos/Magnit_mtz50g.png",
	"Лента":       "https://res.cloudinary.com/dh4qv3hob/image/upload/v1667665905/Promos/Lenta_ocur5v.png",
	"ВиТ":         "https://res.cloudinary.com/dh4qv3hob/image/upload/v1667665906/Promos/ViT_f2xubg.png",
	"ДОДО":        "https://res.cloudinary.com/dh4qv3hob/image/upload/v1667665905/Promos/DODO_kkmdol.webp",
	"Яндекс Плюс": "https://res.cloudinary.com/dh4qv3hob/image/upload/v1667665906/Promos/YandexPlus_cch1ec.jpg",
	"Lamoda":      "https://res.cloudinary.com/dh4qv3hob/image/upload/v1667665905/Promos/Lamoda_xor4fl.jpg",
	"OZON":        "https://res.cloudinary.com/dh4qv3hob/image/upload/v1667665906/Promos/OZON_w2no08.png",
	"Wildberries": "https://res.cloudinary.com/dh4qv3hob/image/upload/v1667665906/Promos/WB_ri56ng.png",
}

func (r *Repository) ChangePriceBook(uuid uuid.UUID, price uint64) (int, error) {
	var store ds.Book
	err := r.db.First(&store, uuid).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 404, err
		}
		return 500, err
	}
	err = r.db.Model(&store).Update("Saleprice", price).Error
	if err != nil {
		return 500, err
	}
	return 0, nil
}

func (r *Repository) DeleteBook(uuid uuid.UUID) (int, error) {
	var store ds.Book
	err := r.db.First(&store, uuid).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 404, err
		}
		return 500, err
	}
	err = r.db.Delete(&store, uuid).Error
	if err != nil {
		return 500, err
	}
	return 0, nil
}

func (r *Repository) CreateStore(store ds.Book) error {
	err := r.db.Create(&store).Error
	return err
}

func (r *Repository) ChangeStore(uuid uuid.UUID, store ds.Book) (int, error) {
	store.UUID = uuid

	err := r.db.Model(&store).Updates(ds.Book{Name: store.Name, Saleprice: store.Saleprice, Year: store.Year, Type: store.Type, Srokgodnost: store.Srokgodnost, Color: store.Color, Description: store.Description, Image: store.Image}).Error
	if err != nil {
		return 500, err
	}

	return 0, nil
}

func (r *Repository) DeleteStore(uuid uuid.UUID) (int, error) {
	var store ds.Book
	err := r.db.First(&store, uuid).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 404, err
		}
		return 500, err
	}
	err = r.db.Delete(&store, uuid).Error
	if err != nil {
		return 500, err
	}
	return 0, nil
}
func (r *Repository) GetDescription(quantity uint64, bookUUID uuid.UUID, userUUID uuid.UUID) (int, string, error) {
	var film ds.Book
	err := r.db.First(&film, bookUUID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 404, "", err
		}
		return 500, "", err
	}

	var cart ds.Cart
	err = r.db.First(&cart, bookUUID, userUUID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 404, "", err
		}
		return 500, "", err
	}
	err = r.db.Delete(&cart, bookUUID, userUUID).Error
	if err != nil {
		return 500, "", err
	}

	err = r.AddOrder(userUUID, bookUUID, quantity)
	if err != nil {
		return 500, "", err
	}

	PromoString := film.Description
	return 0, PromoString, nil
}
