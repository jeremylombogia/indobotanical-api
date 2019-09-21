package transaction

import (
	"time"

	"indobotanical-api/internal"

	"indobotanical-api/models"
	"indobotanical-api/product"

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

	// Check Valid Product ID
	var productID = payload.Data.ProductID

	if !bson.IsObjectIdHex(productID) {
		return c.JSON(404, internal.ErrorResponse{404, "ID is not valid"})
	}

	// Check Product Record
	var product, err = product.FindProduct(productID)
	if err != nil {
		return c.JSON(400, internal.ErrorResponse{400, err.Error()})
	}

	// Check Product Avaibility
	if !product.Avaibility {
		return c.JSON(400, internal.ErrorResponse{400, "Product is not available"})
	}

	// Check Product Stock is bigger than amount
	if product.Stock < payload.Data.Amount {
		return c.JSON(400, internal.ErrorResponse{400, "Product stock is less than your amount"})
	}

	// Count Total Price
	var totalPrice = product.Price * payload.Data.Amount

	// Count by Promo Code
	if payload.Data.PromoCode == "Kratom01" {
		totalPrice = totalPrice - (totalPrice * 10 / 100)
	}

	var transaction = models.Transactions{
		ID:           bson.NewObjectId(),
		Products:     product.ID,
		TotalPrice:   totalPrice,
		PaymentProof: nil,
		User:         bson.ObjectIdHex(authenticatedId),
		CreatedAt:    time.Now(),
	}

	if transaction, err = StoreTransation(&transaction); err != nil {
		return c.JSON(500, internal.ErrorResponse{500, err.Error()})
	}

	return c.JSON(201, internal.SuccessResponse{201, "Transaction created", transaction})
}

func PaymentProof(c echo.Context) error {
	UploadFile()

	return c.JSON(200, "Okay")
	// file, err := c.FormFile("file")
	// if err != nil {
	// 	return err
	// }

	// src, err := file.Open()
	// if err != nil {
	// 	return err
	// }
	// defer src.Close()

	// return err
}
