package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteBooking(ctx *gin.Context) {
	ctx.String(http.StatusOK, "DeleteBooking()")
}
