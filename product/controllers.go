package product

import (
	"time"

	"github.com/jeremylombogia/indobotanical-api/internal"
	"github.com/jeremylombogia/indobotanical-api/models"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

// Index present all products data
func Index(c echo.Context) error {
	var products, err = FetchProduct()
	if err != nil {
		return c.JSON(500, err.Error())
	}

	return c.JSON(200, internal.SuccessResponse{200, nil, products})
}

// Show present all single request data
func Show(c echo.Context) error {
	var id = c.Param("id")

	if !bson.IsObjectIdHex(id) {
		return c.JSON(404, internal.ErrorResponse{400, "You entered invalid id"})
	}

	var product, err = FindProduct(id)
	if err != nil {
		return c.JSON(404, internal.ErrorResponse{404, err.Error()})
	}

	return c.JSON(200, internal.SuccessResponse{200, nil, product})
}

// Post store the request from payload to repository
func Post(c echo.Context) (err error) {
	var payload = new(Payload)
	if err = c.Bind(&payload); err != nil {
		return c.JSON(400, internal.ErrorResponse{400, err.Error()})
	}

	// Fill models.Products struct from Payload.Data.Products
	var product = models.Products{}
	product = payload.Data.Products
	product.ID = bson.NewObjectId()
	product.CreatedAt = time.Now()

	if product, err = StoreProduct(&product); err != nil {
		return c.JSON(500, internal.ErrorResponse{500, err.Error()})
	}

	return c.JSON(201, internal.SuccessResponse{201, "Record created", product})
}

// Patch store the request from payload to repository
func Patch(c echo.Context) (err error) {
	var id = c.Param("id")

	if !bson.IsObjectIdHex(id) {
		return c.JSON(404, internal.ErrorResponse{404, "ID is not valid"})
	}

	var payload = new(Payload)
	if err = c.Bind(payload); err != nil {
		return c.JSON(400, internal.ErrorResponse{400, "Error request"})
	}

	var newProduct = models.Products{}
	newProduct.ID = bson.ObjectIdHex(id)
	newProduct = payload.Data.Products

	product, err := UpdateProduct(&newProduct, id)
	if err != nil {
		return c.JSON(500, internal.ErrorResponse{500, "Error server"})
	}

	return c.JSON(201, internal.SuccessResponse{201, "Record updated", product})
}
