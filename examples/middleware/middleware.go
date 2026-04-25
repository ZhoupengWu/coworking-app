package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	r.Use(func(ctx *gin.Context) {
		fmt.Println("BEFORE - HandleExample")
		ctx.Next()
		fmt.Println("AFTER - HandleExample")
	})
	r.GET("/example", HandleExample)

	if err := r.Run(":8089"); err != nil {
		panic(err)
	}
}

func HandleExample(ctx *gin.Context) {
	fmt.Println("HandleExample()")
	ctx.String(http.StatusOK, fmt.Sprintln("HandleExample..."))
}
