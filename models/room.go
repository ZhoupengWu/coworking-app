package models

import (
	"errors"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type Photo struct {
	ID     int64
	Url    string
	RoomID string
	Room   Room
}

type Room struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Cost          float64   `json:"cost"`
	NumberOfSeats int       `json:"number_of_seats"`
	Category      string    `json:"category"`
	MainPhoto     string    `json:"main_photo"`
	IsAvailable   bool      `gorm:"-:all"`
	Photos        []Photo   `json:"-"`
	Bookings      []Booking `json:"-"`
}

func GetRoomByID(db *gorm.DB, id string) (res *Room, err error) {
	if err = db.Model(&Room{}).First(&res, "id = ?", id).Error; err != nil {
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

func GetRooms(db *gorm.DB, dayToBook time.Time) (res []Room, err error) {
	if err = db.Model(&Room{}).Preload("Bookings").Find(&res).Error; err != nil {
		return nil, CoworkingErr{
			StatusCode: http.StatusInternalServerError,
			Code:       DbErr,
			Message:    err.Error(),
		}
	}

	for k, room := range res {
		res[k].IsAvailable = true

		for _, booking := range room.Bookings {
			if booking.BookedOn.Equal(dayToBook) {
				res[k].IsAvailable = false

				break
			}
		}
	}

	return
}

func GetRoomPhotos(db *gorm.DB, roomID string) (res []string, err error) {
	_, err = GetRoomByID(db, roomID)

	if err != nil {
		return
	}

	if err = db.Model(&Photo{}).Where("room_id = ?", roomID).Select("url").Find(&res).Error; err != nil {
		return nil, CoworkingErr{
			StatusCode: http.StatusInternalServerError,
			Code:       DbErr,
			Message:    err.Error(),
		}
	}

	return
}
