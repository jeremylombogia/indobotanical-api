package main

import (
	"github.com/jeremylombogia/indobotanical-api/product"
	"github.com/jeremylombogia/indobotanical-api/user"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var err error

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/products", product.Index)
	e.GET("/product/:id", product.Show)
	e.POST("/product", product.Post)
	e.PATCH("/product/:id", product.Patch)

	e.POST("/auth/login", user.Login)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))

}
