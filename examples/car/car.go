package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type car struct {
	ID int64
	Brand string
	MaxSpeed int
	FuelType string
}

func main () {
	dsn := "host=localhost port=5430 user=postgres password=postgres dbname=postgres sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic(err)
	}

	sqlDb, err := db.DB()

	if err != nil {
		panic(err)
	}

	if err := sqlDb.Ping(); err != nil {
		panic(err)
	}

	db.AutoMigrate(&car{})
	db.Create([]*car{
		{
			ID: 56824872385798278,
			Brand: "Lamborghini",
			MaxSpeed: 330,
			FuelType: "Benzina",
		},
		{
			ID: 50884873382798278,
			Brand: "BMW",
			MaxSpeed: 216,
			FuelType: "Diesel",
		},
		{
			ID: 14824278395758208,
			Brand: "FIAT",
			MaxSpeed: 160,
			FuelType: "Benzina",
		},
	})

	var cars []car

	if err := db.Model(&car{}).Find(&cars).Error; err != nil {
		panic(err)
	}

	for _, v := range cars {
		fmt.Println(v)
	}
}