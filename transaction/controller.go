package transaction

import (
	"fmt"
	"indobotanical-api/internal"
	"indobotanical-api/models"
	"indobotanical-api/product"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

func Index(c echo.Context) error {
	var transactions, err = FetchTransaction()
	if err != nil {
		return c.JSON(500, internal.ErrorResponse{500, err.Error()})
	}

	return c.JSON(200, internal.SuccessResponse{200, nil, transactions})
}

// Post store transaction
// Payload: ProductID, PromoCode, Amount
func Post(c echo.Context) error {
	var authenticatedId = internal.GetAuthenticatedUserID(c)

	var payload = new(Payload)
	if err := c.Bind(&payload); err != nil {
		return c.JSON(400, internal.ErrorResponse{400, err.Error()})
	}

	// TODO:: need refactoring this struct
	var transaction models.Transactions

	// Count Total Price
	var totalPrice = 0

	for _, indexProduct := range payload.Data.Products {
		if !bson.IsObjectIdHex(indexProduct.ProductID) {
			return c.JSON(404, internal.ErrorResponse{404, "ID is not valid"})
		}

		var product, err = product.FindProduct(indexProduct.ProductID)
		if err != nil {
			return c.JSON(400, internal.ErrorResponse{400, err.Error()})
		}

		// Check Product Avaibility
		if !product.Avaibility {
			return c.JSON(400, internal.ErrorResponse{400, "Product is not available"})
		}

		// Check Product Stock is bigger than amount
		if product.Stock < indexProduct.Amount {
			return c.JSON(400, internal.ErrorResponse{400, "Product stock is less than your amount"})
		}

		// Check Product Stock is bigger than amount
		if product.Stock < indexProduct.Amount {
			return c.JSON(400, internal.ErrorResponse{400, "Product stock is less than your amount"})
		}

		transaction.Products = append(transaction.Products, product)

		totalPrice += product.Price * indexProduct.Amount
	}

	// Count by Promo Code
	if payload.Data.PromoCode == "Kratom01" {
		totalPrice = totalPrice - (totalPrice * 10 / 100)
	}

	transaction.ID = bson.NewObjectId()
	transaction.TotalPrice = totalPrice
	transaction.PaymentProof = nil
	transaction.User = bson.ObjectIdHex(authenticatedId)
	transaction.CreatedAt = time.Now()

	transaction, err := StoreTransation(&transaction)
	if err != nil {
		return c.JSON(500, internal.ErrorResponse{500, err.Error()})
	}

	return c.JSON(201, transaction)
}

// TODO:: Move this to the helper or CDN
func PaymentProof(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	transactionID := c.Param("id")
	extension := filepath.Ext(file.Filename)
	randomNumber := strconv.Itoa(rand.Intn(1000))

	fileName := fmt.Sprintf("cdn/payment-proof/%s%s-%s", transactionID, randomNumber, extension)

	if extension == ".JPG" || extension == ".JPEG" || extension == ".png" {
		src, err := file.Open()
		if err != nil {
			return c.JSON(500, err.Error())
		}
		defer src.Close()

		dst, err := os.Create(fileName)
		if err != nil {
			return c.JSON(500, err.Error())
		}
		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			return c.JSON(500, err.Error())
		}

		return c.JSON(400, internal.SuccessResponse{400, "File uploaded", nil})
	}

	return c.JSON(400, internal.ErrorResponse{400, "File not allowed"})
}
