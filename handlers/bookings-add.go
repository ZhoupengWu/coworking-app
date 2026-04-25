package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddBooking(ctx *gin.Context) {
	ctx.String(http.StatusOK, "AddBooking()")
}
