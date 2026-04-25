package models

type Photo struct {
	Id int64
	Url string
	RoomId string
	Room Room
}

type Room struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Cost float64 `json:"cost"`
	NumberOfSeats int `json:"number_of_seats"`
	Category string `json:"category"`
	MainPhoto string `json:"main_photo"`
	IsAvailable bool `gorm:"-:all"`
	Photos []Photo `json:"-"`
	Bookings []Booking `json:"-"`
}