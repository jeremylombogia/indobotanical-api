package user

import "github.com/jeremylombogia/indobotanical-api/models"

type Payload struct {
	Data struct {
		models.User
	} `json:"data"`
}
