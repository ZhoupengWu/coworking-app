package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllRooms (ctx *gin.Context) {
	ctx.String(http.StatusOK, "GetAllRooms()")
}

func GetRoomByID (ctx *gin.Context) {
	ctx.String(http.StatusOK, "GetRoomByID()")
}

func GetRoomPhotos (ctx *gin.Context) {
	ctx.String(http.StatusOK, "GetRoomPhotos()")
}