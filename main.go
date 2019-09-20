package main

import (
	"indobotanical-api/user"
	"os"

	"github.com/jeremylombogia/indobotanical-api/config"
	"github.com/jeremylombogia/indobotanical-api/product"
	"github.com/jeremylombogia/indobotanical-api/transaction"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var err error

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// TODO:: add versioning API

	// Authentication
	e.POST("/auth/login", user.Login)
	e.POST("/auth/register", user.Register)

	// Authenticated Route
	r := e.Group("")
	r.Use(middleware.JWT([]byte(config.APPKEY)))

	r.POST("/products", product.Post)
	r.POST("/transactions", transaction.Post)
	r.PATCH("/products/:id", product.Patch)

	e.GET("/products", product.Index)
	e.GET("/products/:id", product.Show)
	e.GET("/transactions", transaction.Index)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "1323"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
