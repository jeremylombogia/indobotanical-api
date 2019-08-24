package transaction

import (
	"time"

	"github.com/jeremylombogia/indobotanical-api/internal"
	"github.com/jeremylombogia/indobotanical-api/models"
	"github.com/jeremylombogia/indobotanical-api/product"
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

func Post(c echo.Context) error {
	var authenticatedId = internal.GetAuthenticatedUserID(c)

	var payload = new(Payload)
	if err := c.Bind(&payload); err != nil {
		return c.JSON(400, internal.ErrorResponse{400, err.Error()})
	}

	// Check Stock
	var productId = payload.Data.ProductID

	if !bson.IsObjectIdHex(productId) {
		return c.JSON(404, internal.ErrorResponse{404, "ID is not valid"})
	}

	var product, err = product.FindProduct(productId)
	if err != nil {
		return c.JSON(400, internal.ErrorResponse{400, err.Error()})
	}

	if !product.Avaibility {
		return c.JSON(400, internal.ErrorResponse{400, "Product is not available"})
	}

	if product.Stock < payload.Data.Amount {
		return c.JSON(400, internal.ErrorResponse{400, "Product stock is less than your amount"})
	}

	var totalPrice = product.Price * payload.Data.Amount

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

// UploadPaymentProof give either succes response or error response
func UploadPaymentProof() {
	// Check all of "0" transaction status

}
