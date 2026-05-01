package handlers

import (
	"coworkingapp/models"
	"coworkingapp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserInfo struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignUpReq struct {
	UserInfo
	Email string `json:"email" binding:"required"`
}

func SignUp(ctx *gin.Context) {
	var signUpReq SignUpReq

	if err := ctx.ShouldBind(&signUpReq); err != nil {
		ctx.JSON(http.StatusBadRequest, models.CoworkingErr{
			Code:    models.ValidationErr,
			Message: err.Error(),
		})

		return
	}

	db := ctx.MustGet("DbKey").(*gorm.DB)
	user := models.User{
		Email:    signUpReq.Email,
		Username: signUpReq.Username,
		Password: signUpReq.Password,
	}
	id, err := models.SignUpUser(db, user)

	if err != nil {
		coworkingErr := err.(models.CoworkingErr)
		ctx.JSON(coworkingErr.StatusCode, coworkingErr)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func Login(ctx *gin.Context) {
	var userInfo UserInfo

	if err := ctx.ShouldBind(&userInfo); err != nil {
		ctx.JSON(http.StatusBadRequest, models.CoworkingErr{
			Code:    models.ValidationErr,
			Message: err.Error(),
		})

		return
	}

	db := ctx.MustGet("DbKey").(*gorm.DB)
	signeduser, err := models.LoginUser(db, userInfo.Username, userInfo.Password)

	if err != nil {
		coworkingErr := err.(models.CoworkingErr)
		ctx.JSON(coworkingErr.StatusCode, coworkingErr)

		return
	}

	secretKey := ctx.MustGet("ConfigKey").(models.CoworkingConfig).SecretKey
	token, err := utils.GenerateToken(signeduser.Email, []byte(secretKey))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.CoworkingErr{
			Code:    models.TokenGenerationErr,
			Message: err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
