package main

import (
	"aspera-delivery/handlers"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Welcome to the Aspera Delivery System!")

	e := echo.New()

	e.POST("/upload", handlers.Uploader)

	e.Logger.Fatal(e.Start(":8080"))

}
