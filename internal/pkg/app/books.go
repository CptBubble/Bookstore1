package app

import (
	"Bookstore/internal/app/ds"
	"Bookstore/swaggers/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

// GetBooks		godoc
// @Summary      	Get all books
// @Description  	Get a list of all books
// @Tags         	Info
// @Produce      	json
// @Success      	200 {object} ds.BookDocs
// @Failure 		500 {object} models.ModelError
// @Router       	/books [get]
func (a *Application) GetBooks(gCtx *gin.Context) {
	resp, err := a.repo.GetBooks()
	if err != nil {
		gCtx.JSON(
			http.StatusInternalServerError,
			&models.ModelError{
				Description: "can`t get a list",
				Error:       "db error",
				Type:        "internal",
			})
		return
	}

	gCtx.JSON(http.StatusOK, resp)
}

// GetBook		godoc
// @Summary      	Get book
// @Description  	Get book using its uuid
// @Tags         	Info
// @Produce      	json
// @Param 			UUID path string true "UUID магазина" format(uuid)
// @Success      	200 {object} ds.BookDocs
// @Failure 		500 {object} models.ModelError
// @Router       	/books/{UUID} [get]
func (a *Application) GetBook(gCtx *gin.Context) {
	UUID, err := uuid.Parse(gCtx.Param("uuid"))
	resp, err := a.repo.GetBook(UUID)
	if err != nil {
		gCtx.JSON(
			http.StatusInternalServerError,
			&models.ModelError{
				Description: "Invalid UUID format",
				Error:       "internal",
				Type:        "internal",
			})
		return
	}

	gCtx.JSON(http.StatusOK, resp)
}

// CreateBook		godoc
// @Summary     	Add a new book
// @Description		Adding a new book to database
// @Tags			Add
// @Produce      	json
// @Param 			Promo body ds.BookDocs true "Магазин"
// @Success 		201 {object} swagger.Create
// @Failure 		400 {object} models.ModelError
// @Failure 		500 {object} models.ModelError
// @Router  		/books [post]
func (a *Application) CreateBook(gCtx *gin.Context) {
	promo := ds.Book{}
	err := gCtx.BindJSON(&promo)
	if err != nil {
		gCtx.JSON(
			http.StatusBadRequest,
			&models.ModelError{
				Description: "adding failed",
				Error:       "db error",
				Type:        "internal",
			})
		return
	}

	err = a.repo.CreateBook(promo)
	if err != nil {
		gCtx.JSON(
			http.StatusInternalServerError,
			&models.ModelError{
				Description: "adding failed",
				Error:       "db error",
				Type:        "internal",
			})
		return
	}

	gCtx.JSON(
		http.StatusCreated,
		&models.ModelBookCreated{
			Success: true,
		})
}

// ChangePriceBook		godoc
// @Summary      	Change book price
// @Description  	Change the book price using its uuid
// @Tags         	Change
// @Produce      	json
// @Param 			UUID path string true "UUID книги" format(uuid)
// @Param 			Price body ds.PriceStore true "Новая цена"
// @Success      	200 {object} swagger.Change
// @Failure 		400 {object} models.ModelError
// @Failure 		404 {object} models.ModelError
// @Failure 	 	500 {object} models.ModelError
// @Router       	/books/{UUID} [put]
func (a *Application) ChangePriceBook(gCtx *gin.Context) {
	UUID, err := uuid.Parse(gCtx.Param("uuid"))
	if err != nil {
		gCtx.JSON(
			http.StatusBadRequest,
			&models.ModelError{
				Description: "Invalid UUID format",
				Error:       "db error",
				Type:        "internal",
			})
		return
	}

	promo := ds.Book{}
	err = gCtx.BindJSON(&promo)
	if err != nil {
		gCtx.JSON(
			http.StatusBadRequest,
			&models.ModelError{
				Description: "The price is negative or not int",
				Error:       "db error",
				Type:        "internal",
			})
		return
	}

	code, err := a.repo.ChangePriceBook(UUID, promo.Saleprice)
	if err != nil {
		if code == 404 {
			gCtx.JSON(
				http.StatusNotFound,
				&models.ModelError{
					Description: "UUID Not Found",
					Error:       "db error",
					Type:        "internal",
				})
			return
		} else {
			gCtx.JSON(
				http.StatusInternalServerError,
				&models.ModelError{
					Description: "Change failed",
					Error:       "db error",
					Type:        "internal",
				})
			return
		}
	}

	gCtx.JSON(
		http.StatusOK,
		&models.ModelPriceChanged{
			Success: true,
		})
}

// DeleteBook		godoc
// @Summary     	Delete a book
// @Description 	Delete a book using its uuid
// @Tags         	Delete
// @Produce      	json
// @Param 			UUID path string true "UUID книги" format(uuid)
// @Success      	200 {object} swagger.Delete
// @Failure 		400 {object} models.ModelError
// @Failure 		404 {object} models.ModelError
// @Failure 	 	500 {object} models.ModelError
// @Router       	/books/{UUID} [delete]
func (a *Application) DeleteBook(gCtx *gin.Context) {
	UUID, err := uuid.Parse(gCtx.Param("uuid"))
	if err != nil {
		gCtx.JSON(
			http.StatusBadRequest,
			&models.ModelError{
				Description: "Invalid UUID format",
				Error:       "db error",
				Type:        "internal",
			})
		return
	}

	code, err := a.repo.DeleteBook(UUID)
	if err != nil {
		if code == 404 {
			gCtx.JSON(
				http.StatusNotFound,
				&models.ModelError{
					Description: "UUID Not Found",
					Error:       "db error",
					Type:        "internal",
				})
			return
		} else {
			gCtx.JSON(
				http.StatusInternalServerError,
				&models.ModelError{
					Description: "Delete failed",
					Error:       "db error",
					Type:        "internal",
				})
			return
		}
	}

	gCtx.JSON(
		http.StatusOK,
		&models.ModelBookDeleted{
			Success: true,
		})
}
func (a *Application) CreateStore(gCtx *gin.Context) {
	store := ds.Book{}

	err := gCtx.BindJSON(&store)
	if err != nil {
		gCtx.JSON(
			http.StatusBadRequest,
			&models.ModelError{
				Description: "Invalid parameters",
				Error:       "Invalid parameters",
				Type:        "Invalid parameters",
			})
		return
	}

	err = a.repo.CreateStore(store)
	if err != nil {
		gCtx.JSON(
			http.StatusInternalServerError,
			&models.ModelError{
				Description: "Create failed",
				Error:       "Create failed",
				Type:        "Create failed",
			})
		return
	}

	gCtx.JSON(
		http.StatusCreated,
		&models.ModelBookCreated{
			Success: true,
		})
}
func (a *Application) ChangeStore(gCtx *gin.Context) {
	UUID, err := uuid.Parse(gCtx.Param("uuid"))
	if err != nil {
		gCtx.JSON(
			http.StatusBadRequest,
			&models.ModelError{
				Description: "Invalid UUID format",
				Error:       "Invalid UUID format",
				Type:        "Invalid UUID format",
			})
		return
	}

	store := ds.Book{}
	err = gCtx.BindJSON(&store)
	//if err != nil {
	//	gCtx.JSON(
	//		http.StatusBadRequest,
	//		&models.ModelError{
	//			Description: "The price is negative or not int",
	//			Error:       "The price is negative or not int",
	//			Type:        "The price is negative or not int",
	//		})
	//	return
	//}

	code, err := a.repo.ChangeStore(UUID, store)
	if err != nil {
		if code == 404 {
			gCtx.JSON(
				http.StatusNotFound,
				&models.ModelError{
					Description: "UUID Not Found",
					Error:       "UUID Not Found",
					Type:        "UUID Not Found",
				})
			return
		} else {
			gCtx.JSON(
				http.StatusInternalServerError,
				&models.ModelError{
					Description: "Change failed",
					Error:       "Change failed",
					Type:        "Change failed",
				})
			return
		}
	}

	gCtx.JSON(
		http.StatusOK,
		&models.ModelPriceChanged{
			Success: true,
		})
}
func (a *Application) DeleteStore(gCtx *gin.Context) {
	UUID, err := uuid.Parse(gCtx.Param("uuid"))
	if err != nil {
		gCtx.JSON(
			http.StatusBadRequest,
			&models.ModelError{
				Description: "Invalid UUID format",
				Error:       "Invalid UUID format",
				Type:        "Invalid UUID format",
			})
		return
	}

	code, err := a.repo.DeleteStore(UUID)
	if err != nil {
		if code == 404 {
			gCtx.JSON(
				http.StatusNotFound,
				&models.ModelError{
					Description: "UUID Not Found",
					Error:       "UUID Not Found",
					Type:        "UUID Not Found",
				})
			return
		} else {
			gCtx.JSON(
				http.StatusInternalServerError,
				&models.ModelError{
					Description: "Delete failed",
					Error:       "Delete failed",
					Type:        "Delete failed",
				})
			return
		}
	}

	gCtx.JSON(
		http.StatusOK,
		&models.ModelBookDeleted{
			Success: true,
		})
}

// GetVideo godoc
// @Summary Get a promo
// @Description Get a promo in store using its uuid
// @Tags Info
// @Produce json
// @Param UUID path string true "UUID магазина" format(uuid)
// @Param Quantity path string true "Кол-во"
// @Success 200 {object} swagger.StorePromo
// @Failure 400 {object} swagger.Error
// @Failure 404 {object} swagger.Error
// @Failure 500 {object} swagger.Error
// @Router /film/{UUID}/{Quantity} [get]
func (a *Application) GetDescription(gCtx *gin.Context) {
	jwtStr := gCtx.GetHeader("Authorization")
	userUUID := a.GetUserByToken(jwtStr)

	BookUUID, err := uuid.Parse(gCtx.Param("uuid"))
	quantity, _ := strconv.Atoi(gCtx.Param("quantity"))
	if err != nil {
		gCtx.JSON(
			http.StatusBadRequest,
			models.ModelError{
				Description: "Invalid UUID format",
				Error:       "Invalid UUID format",
				Type:        "Invalid UUID format",
			})
		return
	}

	code, Description, err := a.repo.GetDescription(uint64(quantity), BookUUID, userUUID)
	if err != nil {
		if code == 404 {
			gCtx.JSON(
				http.StatusNotFound,
				&models.ModelError{
					Description: "UUID Not Found",
					Error:       "Invalid UUID format",
					Type:        "Invalid UUID format",
				})
			return
		} else {
			gCtx.JSON(
				http.StatusInternalServerError,
				&models.ModelError{
					Description: "Delete failed",
					Error:       "Delete failed",
					Type:        "Delete failed",
				})
			return
		}
	}

	gCtx.JSON(http.StatusOK, Description)
}
