package product

import "github.com/jeremylombogia/indobotanical-api/models"

type Payload struct {
	Data struct {
		models.Products
	} `json:"data"`
}
