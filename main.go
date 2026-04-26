package main

import (
	"coworkingApp/handlers"
	"coworkingApp/middleware"
	"coworkingApp/models"
	"coworkingApp/utils"
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var config models.CoworkingConfig

func init() {
	data, err := os.ReadFile("config/config.json")

	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal([]byte(data), &config); err != nil {
		panic(err)
	}
}

func main() {
	gin.SetMode(gin.DebugMode)
	db, err := gorm.Open(postgres.Open(config.Dsn))

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Room{})
	db.AutoMigrate(&models.Photo{})
	db.AutoMigrate(&models.Booking{})
	seedData(db)

	r := gin.Default()

	r.Use(middleware.EarlyExitOnPreflighRequests())
	r.Use(middleware.SetCorsPolicy(config.AllowedOrigins))
	r.Use(func(ctx *gin.Context) {
		ctx.Set("DbKey", db)
		ctx.Set("ConfigKey", config)
		ctx.Next()
	})

	r.GET("/rooms", handlers.GetAllRooms)
	r.GET("/rooms/:id", handlers.GetRoomByID)
	r.GET("/rooms/:id/photos", handlers.GetRoomPhotos)
	r.GET("/bookings", handlers.GetBookingsByUserID)
	r.GET("/bookings/:id", handlers.GetBookingsByID)
	r.POST("/bookings", handlers.AddBooking)
	r.DELETE("/bookings/:id", handlers.DeleteBooking)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}

func seedData(db *gorm.DB) {
	db.Delete(&models.User{}, "1 = 1")
	db.Delete(&models.Room{}, "1 = 1")
	db.Delete(&models.Photo{}, "1 = 1")
	db.Delete(&models.Booking{}, "1 = 1")

	userID := utils.GetUUID()
	db.Create(&models.User{
		ID:       userID,
		Email:    "wux@wux.com",
		Username: "wux",
		Password: "wux1xuw9",
	})
	db.Create([]*models.Room{
		{
			ID:            utils.GetUUID(),
			Name:          "Green",
			Cost:          99.99,
			NumberOfSeats: 6,
			Category:      "Spazio di piacere",
			MainPhoto:     "/green_0001.png",
			Photos: []models.Photo{
				{
					Url: "/green_0002.png",
				},
				{
					Url: "/green_0003.png",
				},
				{
					Url: "/green_0004.png",
				},
			},
		},
		{
			ID:            utils.GetUUID(),
			Name:          "Red",
			Cost:          599.99,
			NumberOfSeats: 150,
			Category:      "Sala conferenza",
			MainPhoto:     "/red_0001.png",
			Photos: []models.Photo{
				{
					Url: "/red_0002.png",
				},
				{
					Url: "/red_0003.png",
				},
				{
					Url: "/red_0004.png",
				},
			},
		},
		{
			ID:            utils.GetUUID(),
			Name:          "Yellow",
			Cost:          299.99,
			NumberOfSeats: 50,
			Category:      "Aula di apprendimento",
			MainPhoto:     "/yellow_0001.png",
			Photos: []models.Photo{
				{
					Url: "/yellow_0002.png",
				},
				{
					Url: "/yellow_0003.png",
				},
				{
					Url: "/yellow_0004.png",
				},
			},
		},
	})
}
