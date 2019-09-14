package main

import (
	"fmt"
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

	// Authentication
	e.POST("/auth/login", user.Login)
	e.POST("/auth/register", user.Register)

	// Authenticated
	r := e.Group("")
	r.Use(middleware.JWT([]byte(config.APPKEY)))
	// Authenticated Product
	r.POST("/product", product.Post)
	r.PATCH("/product/:id", product.Patch)

	// Authenticated Transaction
	r.POST("/transaction", transaction.Post)

	// Product
	e.GET("/products", product.Index)
	e.GET("/product/:id", product.Show)

	// Transaction
	e.GET("/transactions", transaction.Index)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Errorf("$PORT not set")
	}

	e.Logger.Fatal(e.Start(":" + port))
}
