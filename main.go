package main

import (
	"indobotanical-api/user"
	"os"

	"indobotanical-api/config"
	"indobotanical-api/product"
	"indobotanical-api/transaction"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var err error

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/cdn", "cdn")

	e.GET("/", func(c echo.Context) error {
		return c.JSON(201, map[string]interface{}{
			"message": "API Response OK",
		})
	})

	// TODO:: add versioning API
	// Authentication
	e.POST("/auth/login", user.Login)
	e.POST("/auth/register", user.Register)

	// Authenticated Route
	r := e.Group("")
	r.Use(middleware.JWT([]byte(config.APPKEY)))

	r.GET("/transactions", transaction.Index)
	r.POST("/products", product.Post)
	r.POST("/transactions", transaction.Post)
	r.PATCH("/transactions/payment-proof/:id", transaction.PaymentProof)
	r.PATCH("/products/:id", product.Patch)

	e.GET("/products", product.Index)
	e.GET("/products/:id", product.Show)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
