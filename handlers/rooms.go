package handlers

import (
	"coworkingapp/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllRooms(ctx *gin.Context) {
	rawDayToBook := ctx.Query("day_to_book")

	if rawDayToBook == "" {
		ctx.JSON(http.StatusBadRequest, models.CoworkingErr{
			Code:    models.MissingParams,
			Message: "day_to_book is required",
		})

		return
	}

	dayToBook, err := time.Parse("2006-01-02", rawDayToBook)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.CoworkingErr{
			Code:    models.DateWrongFormatErr,
			Message: err.Error(),
		})

		return
	}

	db := ctx.MustGet("DbKey").(*gorm.DB)
	rooms, err := models.GetRooms(db, dayToBook)

	if err != nil {
		coworkingErr := err.(models.CoworkingErr)
		ctx.JSON(coworkingErr.StatusCode, coworkingErr)

		return
	}

	ctx.JSON(http.StatusOK, rooms)
}

func GetRoomByID(ctx *gin.Context) {
	db := ctx.MustGet("DbKey").(*gorm.DB)
	room, err := models.GetRoomByID(db, ctx.Param("id"))

	if err != nil {
		coworkingErr := err.(models.CoworkingErr)
		ctx.JSON(coworkingErr.StatusCode, coworkingErr)

		return
	}

	ctx.JSON(http.StatusOK, room)
}

func GetRoomPhotos(ctx *gin.Context) {
	db := ctx.MustGet("DbKey").(*gorm.DB)
	photos, err := models.GetRoomPhotos(db, ctx.Param("id"))

	if err != nil {
		coworkingErr := err.(models.CoworkingErr)
		ctx.JSON(coworkingErr.StatusCode, coworkingErr)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":     ctx.Param("id"),
		"photos": photos,
	})
}
