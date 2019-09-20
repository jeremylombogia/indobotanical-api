package product

import "indobotanical-api/models"

type Payload struct {
	Data struct {
		models.Products
	} `json:"data"`
}
