package user

import "indobotanical-api/models"

type Payload struct {
	Data struct {
		models.User
	} `json:"data"`
}
