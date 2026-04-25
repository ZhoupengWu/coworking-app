package main

import (
	"coworkingApp/handlers"

	"github.com/gin-gonic/gin"
)

func main () {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

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



















/*
func main () {
	router := http.NewServeMux()
	router.HandleFunc("/example", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("/example endpoint has been invoked")
		w.Write([]byte(fmt.Sprintf("This is a demo endpoint...")))
	})
	router.HandleFunc("/health", HandlerHealthCheck)

	err := http.ListenAndServe(":8089", router)

	if err != nil {
		panic(err)
	}
}

func HandlerHealthCheck (w http.ResponseWriter, r *http.Request) {
	fmt.Println("HandlerHealthCheck has been invoked")

	if r.Method != http.MethodGet {
		w.Write([]byte(fmt.Sprintf("The request must be a GET")))

		return
	}

	w.Write([]byte(fmt.Sprintf("The system is health")))
} */