package app

import (
	"Bookstore/internal/app/ds"
	"Bookstore/swaggers/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

//func (a *Application) AddOrder(gCtx *gin.Context) {
//	jwtStr := gCtx.GetHeader("Authorization")
//	userUUID := a.GetUserByToken(jwtStr)
//	order := ds.Order{}
//	order.UserUUID = userUUID
//	err := gCtx.BindJSON(&order)
//	if err != nil {
//		gCtx.JSON(
//			http.StatusBadRequest,
//			&swagger.Error{
//				Description: "Invalid parameters",
//				Error:       swagger.Err400,
//				Type:        swagger.TypeClientReq,
//			})
//		return
//	}
//	err = a.repo.AddOrder(order)
//	if err != nil {
//		gCtx.JSON(
//			http.StatusInternalServerError,
//			&swagger.Error{
//				Description: "Create failed",
//				Error:       swagger.Err500,
//				Type:        swagger.TypeInternalReq,
//			})
//		return
//	}
//	gCtx.JSON(
//		http.StatusOK,
//		&swagger.Create{
//			Success: true,
//		})
//
//}

func (a *Application) GetOrders(gCtx *gin.Context) {
	stDate := gCtx.Query("start_date")
	endDate := gCtx.Query("end_date")
	status := gCtx.Query("status")
	resp, err := a.repo.GetOrders(stDate, endDate, status)
	if err != nil {
		gCtx.JSON(
			http.StatusInternalServerError,
			&models.ModelError{
				Description: "can`t get a list",
				Error:       "can`t get a list",
				Type:        "can`t get a list",
			})
		return
	}
	gCtx.JSON(http.StatusOK, resp)
}

func (a *Application) ChangeStatus(gCtx *gin.Context) {
	UUID, err := uuid.Parse(gCtx.Param("uuid"))
	if err != nil {
		gCtx.JSON(
			http.StatusBadRequest,
			&models.ModelError{
				Description: "Invalid UUID format",
				Error:       "can`t get a list",
				Type:        "can`t get a list",
			})
		return
	}
	order := ds.Order{}
	err = gCtx.BindJSON(&order)
	if err != nil {
		gCtx.JSON(
			http.StatusBadRequest,
			&models.ModelError{
				Description: "The price is negative or not int",
				Error:       "can`t get a list",
				Type:        "can`t get a list",
			})
		return
	}
	resp, err := a.repo.ChangeStatus(UUID, order.Status)
	if err != nil {
		if resp == 404 {
			gCtx.JSON(
				http.StatusNotFound,
				&models.ModelError{
					Description: "UUID Not Found",
					Error:       "can`t get a list",
					Type:        "can`t get a list",
				})
			return
		} else {
			gCtx.JSON(
				http.StatusInternalServerError,
				&models.ModelError{
					Description: "Change failed",
					Error:       "can`t get a list",
					Type:        "can`t get a list",
				})
			return
		}
	}

	gCtx.JSON(
		http.StatusOK,
		&models.ModelBookCreated{
			Success: true,
		})

}
