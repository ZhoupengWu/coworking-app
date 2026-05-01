package handlers

import (
	"coworkingapp/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteBooking(ctx *gin.Context) {
	userID := ctx.MustGet("UserIDKey").(string)
	db := ctx.MustGet("DbKey").(*gorm.DB)

	if err := models.DeleteBookingByID(db, ctx.Param("id"), userID); err != nil {
		coworkingErr := err.(models.CoworkingErr)
		ctx.JSON(coworkingErr.StatusCode, coworkingErr)

		return
	}

	ctx.Status(http.StatusNoContent)
}
