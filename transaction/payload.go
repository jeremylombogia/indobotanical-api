package transaction

type Payload struct {
	Data struct {
		ProductID  string      `json:"productId"`
		Amount     int         `json:"amount"`
		PromoCode  string      `json:"promoCode"`
		UploadFile interface{} `json:"uploadFile"`
	} `json:"data"`
}
