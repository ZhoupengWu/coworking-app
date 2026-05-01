package handlers

import (
	"coworkingapp/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetBookingsByUserID(ctx *gin.Context) {
	userID := ctx.MustGet("UserIDKey").(string)
	db := ctx.MustGet("DbKey").(*gorm.DB)
	bookings, err := models.GetBookingsByUserID(db, userID)

	if err != nil {
		coworkingErr := err.(models.CoworkingErr)
		ctx.JSON(coworkingErr.StatusCode, coworkingErr)

		return
	}

	ctx.JSON(http.StatusOK, bookings)
}

func GetBookingsByID(ctx *gin.Context) {
	userID := ctx.MustGet("UserIDKey").(string)
	db := ctx.MustGet("DbKey").(*gorm.DB)
	booking, err := models.GetBookingByID(db, ctx.Param("id"), userID)

	if err != nil {
		coworkingErr := err.(models.CoworkingErr)
		ctx.JSON(coworkingErr.StatusCode, coworkingErr)

		return
	}

	ctx.JSON(http.StatusOK, booking)
}
