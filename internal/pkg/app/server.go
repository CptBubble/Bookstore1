package app

import (
	_ "Bookstore/docs"
	"Bookstore/internal/app/role"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Users struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Book string `json:"book"`
}

func (a *Application) StartServer() {

	r := gin.Default()

	r.Use(CORSMiddleware())

	// Запрос для свагера

	// Запросы для магазина:

	r.GET("/books", a.GetBooks)

	// Запросы для корзины
	r.GET("/books/:uuid/:quantity", a.GetDescription)
	r.GET("/book/:uuid", a.GetBook)
	r.GET("/cart/:books", a.GetCart1)

	r.GET("/cart/increase/:books", a.IncreaseQuantity)

	r.GET("/cart/decrease/:books", a.DecreaseQuantity)

	r.DELETE("/cart/delete/:books", a.DeleteCart)
	// Запросы для авторизации
	r.POST("/login", a.Login)

	r.POST("/sign_up", a.Register)

	r.GET("/logout", a.Logout)

	r.GET("/role", a.Role)

	// Запросы для всех авторизированных пользователей
	r.Use(a.WithAuthCheck(role.Buyer, role.Manager, role.Admin)).GET("/cart", a.GetCart)

	// Запросы для менеджеров
	r.Use(a.WithAuthCheck(role.Manager)).POST("/books", a.CreateStore)
	//
	r.Use(a.WithAuthCheck(role.Manager)).DELETE("/store/:uuid", a.DeleteStore)

	r.Use(a.WithAuthCheck(role.Manager)).PUT("/books/:uuid", a.ChangeStore)

	r.Use(a.WithAuthCheck(role.Manager)).GET("/orders", a.GetOrders)

	r.Use(a.WithAuthCheck(role.Manager)).PUT("/orders/:uuid", a.ChangeStatus)

	r.Use(a.WithAuthCheck(role.Manager)).GET("/user/:uuid", a.GetUser)
	_ = r.Run()

	log.Println("server down")

}

//type inter struct {
//	Status string `json:"status"`
//}
//
//// GetBookList godoc
//// @Summary      Get all records
//// @Description  Get a list of all books
//// @Tags         Info
//// @Produce      json
//// @Success      200  {object}  ds.Book
//// @Failure 500 {object} models.ModelError
//// @Router       /books [get]
//func (a *Application) GetBookList(gCtx *gin.Context) {
//	resp, err := a.repo.GetBookList()
//	if err != nil {
//		gCtx.JSON(
//			http.StatusInternalServerError,
//			&models.ModelError{
//				Description: "can`t get a list",
//				Error:       "db error",
//				Type:        "internal",
//			})
//		return
//	}
//	gCtx.JSON(http.StatusOK, resp)
//
//}
//
//// GetBookPrice godoc
//// @Summary      Get price for a book
//// @Description  Get a price via uuid of a book
//// @Tags         Info
//// @Produce      json
//// @Param UUID query string true "UUID книги"
//// @Success      200  {object}  models.ModelBookPrice
//// @Failure 	 500 {object} models.ModelError
//// @Router       /books/:uuid [get]
//func (a *Application) GetBookPrice(gCtx *gin.Context) {
//	uuid := gCtx.Param("uuid")
//	//uuid := c.Query("UUID")
//	log.Println(uuid)
//	respName, respDesc, respPrice, err := a.repo.GetBookPrice(uuid)
//	if err != nil {
//		if respName == "no product found with this uuid" {
//			gCtx.JSON(
//				http.StatusBadRequest,
//				&models.ModelError{
//					Description: "No product found with this uuid",
//					Error:       "uuid error",
//					Type:        "client",
//				})
//			return
//		}
//		gCtx.JSON(
//			http.StatusInternalServerError,
//			&models.ModelError{
//				Description: "can`t get a price",
//				Error:       "db error",
//				Type:        "internal",
//			})
//		return
//	}
//	gCtx.JSON(
//		http.StatusOK,
//		&models.ModelBookPrice{
//			Name:        respName,
//			Description: respDesc,
//			Saleprice:   strconv.Itoa(respPrice),
//		})
//
//}
//
//// ChangePrice   godoc
//// @Summary      Change book price
//// @Description  Change a price for a book via its uuid
//// @Tags         Change
//// @Produce      json
//// @Param UUID query string true "UUID книги"
//// @Param Price query int true "Новая цена"
//// @Success      200  {object}  models.ModelPriceChanged
//// @Failure 	 500 {object} models.ModelError
//// @Router       /books/:uuid [put]
//func (a Application) ChangePrice(gCtx *gin.Context) {
//	uuidR := gCtx.Param("uuid")
//	log.Println(uuidR)
//	newPrice, _ := strconv.Atoi(gCtx.Query("Price"))
//	log.Println(newPrice)
//	if newPrice <= 0 {
//		gCtx.JSON(
//			http.StatusBadRequest,
//			&models.ModelError{
//				Description: "The price cannot be non -negative",
//				Error:       "Price error",
//				Type:        "client",
//			})
//		return
//	}
//	inputUuid, _ := uuid.Parse(uuidR)
//	err, messageError := a.repo.ChangePrice(inputUuid, newPrice)
//	if err != nil {
//		if messageError == "record not found" {
//			gCtx.JSON(
//				http.StatusNotFound,
//				&models.ModelError{
//					Description: "record failed",
//					Error:       "db error",
//					Type:        "client",
//				})
//			return
//		}
//		gCtx.JSON(
//			http.StatusInternalServerError,
//			&models.ModelError{
//				Description: "Update failed",
//				Error:       "db error",
//				Type:        "internal",
//			})
//		return
//	}
//	gCtx.JSON(
//		http.StatusOK,
//		&models.ModelPriceChanged{
//			Success: true,
//		})
//}
//
//// DeleteBook   godoc
//// @Summary      Delete a book
//// @Description  Delete a book via its uuid
//// @Tags         Change
//// @Produce      json
//// @Param UUID query string true "UUID книги"
//// @Success      200  {object}  models.ModelBookDeleted
//// @Failure 	 500 {object} models.ModelError
//// @Router       /books [delete]
//func (a *Application) DeleteBook(gCtx *gin.Context) {
//	uuid := gCtx.Param("uuid")
//	err := a.repo.DeleteBook(uuid)
//	if err != nil {
//		gCtx.JSON(
//			http.StatusNotFound,
//			&models.ModelError{
//				Description: "delete failed",
//				Error:       "db error",
//				Type:        "internal",
//			})
//		return
//	}
//	gCtx.JSON(
//		http.StatusOK,
//		&models.ModelBookDeleted{
//			Success: true,
//		})
//
//}
//
//// AddBook godoc
//// @Summary      Add a new book
//// @Description  Adding a new book to database
//// @Tags         Add
//// @Produce      json
//// @Param Name query string true "Название книги"
//// @Param Saleprice query int true "Цена книги"
//// @Param Year query int true "Год производства"
//// @Param Type query string true "Тип книги"
//// @Param Srokgodnost query int true "Срок годности"
//// @Param Color query string true "Цвет книги"
//// @Param Pokupatel query string true "Покупатель"
//// @Param Description query string true "Описание"
//// @Success      201  {object}  models.ModelBookCreated
//// @Failure 500 {object} models.ModelError
//// @Router       /books [Post]
//func (a *Application) AddBook(gCtx *gin.Context) {
//	book := ds.Book{}
//	if err := gCtx.BindJSON(&book); err != nil {
//		gCtx.JSON(
//			http.StatusInternalServerError,
//			&models.ModelError{
//				Description: "adding failed",
//				Error:       "db error",
//				Type:        "internal",
//			})
//		return
//	}
//	if book.Saleprice <= 0 {
//		gCtx.JSON(
//			http.StatusBadRequest,
//			&models.ModelError{
//				Description: "The price cannot be non -negative",
//				Error:       "Price error",
//				Type:        "client",
//			})
//		return
//	}
//	book.UUID = uuid.New()
//	err := a.repo.AddBook(book)
//	if err != nil {
//		gCtx.JSON(
//			http.StatusInternalServerError,
//			&models.ModelError{
//				Description: "adding failed",
//				Error:       "db error",
//				Type:        "internal",
//			})
//		return
//	}
//	gCtx.JSON(
//		http.StatusOK,
//		&models.ModelBookCreated{
//			Success: true,
//		})
//}
