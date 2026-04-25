package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	ID string `json:"id"`
	BookedOn time.Time `json:"booked_on"`
	CreatedAt time.Time `json:"created_at"`
	RoomID string `json:"room_id"`
	UserID string `json:"-"`
	Room Room `json:"-"`
	User User `json:"-"`
}

func CreateBooking (db *gorm.DB, booking Booking) (id *string, err error) {
	if err = db.Model(&Booking{}).Create(&booking).Error; err != nil {
		return nil, err
	}

	return
}

func GetBookingsByUserID (db *gorm.DB, userID string) (res []Booking, err error) {
	if err = db.Model(&Booking{}).Where("user_id = ?", userID).Find(&res).Error; err != nil {
		return nil, err
	}

	return
}

func GetBookingByID (db *gorm.DB, id, userID string) (res *Booking, err error) {
	if err = db.Model(&Booking{}).Where("id = ? and user_id = ?", id, userID).First(&res).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		return nil, err
	}

	return
}

func DeleteBookingByID (db *gorm.DB, id, userID string) (err error) {
	booking, err := GetBookingByID(db, id, userID)

	if err != nil {
		return err
	}

	if err = db.Model(&Booking{}).Delete(&booking).Error; err != nil {
		return err
	}

	return nil
}