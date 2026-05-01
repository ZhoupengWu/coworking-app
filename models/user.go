package models

import (
	"coworkingapp/utils"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

type User struct {
	ID       string
	Email    string
	Username string
	Password string
}

func SignUpUser(db *gorm.DB, user User) (id string, err error) {
	if err = db.Model(&User{}).First(&User{}, "email = ?", user.Email).Error; err == nil {
		return "", CoworkingErr{
			StatusCode: http.StatusBadRequest,
			Code:       EmailAlreadyInUseErr,
			Message:    "Please change the email and retry",
		}
	}

	user.ID = utils.GetUUID()

	if err = db.Model(&User{}).Create(&user).Error; err != nil {
		return "", CoworkingErr{
			StatusCode: http.StatusInternalServerError,
			Code:       DbErr,
			Message:    err.Error(),
		}
	}

	return user.ID, nil
}

func LoginUser(db *gorm.DB, username, password string) (res *User, err error) {
	if err = db.Model(&User{}).Where("username = ? and password = ?", username, password).First(&res).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, CoworkingErr{
				StatusCode: http.StatusNotFound,
				Code:       InvalidCredentialsErr,
				Message:    err.Error(),
			}
		}

		return nil, CoworkingErr{
			StatusCode: http.StatusInternalServerError,
			Code:       DbErr,
			Message:    err.Error(),
		}
	}

	return
}

func GetUserByEmail(db *gorm.DB, email string) (res *User, err error) {
	if err = db.Model(&User{}).Where("email = ?", email).First(&res).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, CoworkingErr{
				StatusCode: http.StatusNotFound,
				Code:       ObjectNotFoundErr,
				Message:    err.Error(),
			}
		}

		return nil, CoworkingErr{
			StatusCode: http.StatusInternalServerError,
			Code:       DbErr,
			Message:    err.Error(),
		}
	}

	return
}
