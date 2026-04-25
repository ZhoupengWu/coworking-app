package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBookingsByUserID (ctx *gin.Context) {
	ctx.String(http.StatusOK, "GetBookingsByUserID()")
}

func GetBookingsByID (ctx *gin.Context) {
	ctx.String(http.StatusOK, "GetBookingsByID()")
}