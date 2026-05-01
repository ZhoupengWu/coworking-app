package handlers

import (
	"coworkingapp/models"
	"coworkingapp/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookingDto struct {
	RoomID   string `json:"room_id" binding:"required"`
	BookedOn string `json:"booked_on" binding:"required"`
}

func mapBookingDtoToModel(dto BookingDto, userID string) (model *models.Booking, err error) {
	bookedOn, err := time.Parse("2006-01-02", dto.BookedOn)

	if err != nil {
		return nil, models.CoworkingErr{
			StatusCode: http.StatusBadRequest,
			Code:       models.DateWrongFormatErr,
			Message:    err.Error(),
		}
	}

	model = &models.Booking{
		ID:        utils.GetUUID(),
		BookedOn:  bookedOn,
		CreatedAt: time.Now(),
		RoomID:    dto.RoomID,
		UserID:    userID,
	}

	return
}

func AddBooking(ctx *gin.Context) {
	var bookingDto BookingDto

	if err := ctx.ShouldBind(&bookingDto); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())

		return
	}

	userID := ctx.MustGet("UserIDKey").(string)
	model, err := mapBookingDtoToModel(bookingDto, userID)

	if err != nil {
		coworkingErr := err.(models.CoworkingErr)
		ctx.JSON(coworkingErr.StatusCode, coworkingErr)

		return
	}

	db := ctx.MustGet("DbKey").(*gorm.DB)
	id, err := models.CreateBooking(db, *model)

	if err != nil {
		coworkingErr := err.(models.CoworkingErr)
		ctx.JSON(coworkingErr.StatusCode, coworkingErr)

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id": *id,
	})
}
