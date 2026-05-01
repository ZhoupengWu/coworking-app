package middleware

import (
	"coworkingapp/models"
	"coworkingapp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthorizeUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenHeader := ctx.GetHeader("Authorization")

		if tokenHeader == "" {
			ctx.JSON(http.StatusUnauthorized, models.CoworkingErr{
				Code:    models.MissingTokenErr,
				Message: "Please provide a JWT token along with the HTTP headers",
			})

			return
		}

		secretKey := ctx.MustGet("ConfigKey").(models.CoworkingConfig).SecretKey
		claims, err := utils.ValidateToken(tokenHeader, []byte(secretKey))

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, models.CoworkingErr{
				Code:    models.TokenNotValidErr,
				Message: err.Error(),
			})

			return
		}

		email := (*claims)["sub"].(string)
		db := ctx.MustGet("DbKey").(*gorm.DB)
		user, err := models.GetUserByEmail(db, email)

		if err != nil {
			coworkingErr := err.(models.CoworkingErr)
			ctx.JSON(coworkingErr.StatusCode, coworkingErr)

			return
		}

		ctx.Set("UserIDKey", user.ID)
		ctx.Next()
	}
}
